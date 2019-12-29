package main

import (
	"BCDns_daemon/message"
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	SwapCert int = iota
	Start
)

type Node struct {
	IP string
	Client BCDns_daemon.MethodClient
	IsLeader bool
}

var (
	action = flag.Int("action", 0, "Action")
	ip = flag.String("ip", "", "IP")
	frq = flag.Int("frq", 25, "frequency")
	byzantine = flag.Bool("by", false, "Byzantine")
	hosts = map[string]Node{}
	Leader *Node
)

func main() {
	flag.Parse()
	file, err := os.Open("/workspace/hosts")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	bf := bufio.NewReader(file)
	for {
		line, _, err := bf.ReadLine()
		if err == io.EOF {
			break
		}
		l := string(line)
		node := Node{
			IP:strings.Split(l, " ")[0] + ":5000",
		}
		conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
		if err != nil {
			panic(hosts[strings.Split(l, " ")[1]])
		}
		defer conn.Close()

		node.Client = BCDns_daemon.NewMethodClient(conn)
		hosts[strings.Split(l, " ")[1]] = node
	}
	switch *action {
	case SwapCert:
		for _, node := range hosts {
			_, err = node.Client.DoSwapCert(context.Background(), &BCDns_daemon.SwapCertMsg{
				Ip: *ip,
			})
			if err != nil {
				panic(err)
			}
		}
	case Start:
		mux := sync.Mutex{}
		count := int32(0)
		wt := &sync.WaitGroup{}
		f := (len(hosts) - 1) / 3
		for _, node := range hosts {
			wt.Add(1)
			go func(node Node) {
				defer wt.Done()
				var req BCDns_daemon.StartServerReq
				mux.Lock()
				if *byzantine && f != 0 {
					req.Byzantine = true
					f--
				} else {
					req.Byzantine = false
				}
				mux.Unlock()
				rep, err := node.Client.DoStartServer(context.Background(), &req)
				if err != nil {
					fmt.Println(err)
					return
				}
				if rep.IsLeader {
					Leader = &node
				}
				atomic.AddInt32(&count, 1)
			}(node)
		}
		wt.Wait()
		if count == int32(len(hosts)) {
			rep, err := Leader.Client.DoStartClient(context.Background(), &BCDns_daemon.StartClientReq{
				Frq:int32(*frq),
			})
			if err != nil {
				panic(err)
			}
			fmt.Println(rep.Latency, rep.Throughout)
		}
		for _, node := range hosts {
			wt.Add(1)
			go func() {
				defer wt.Done()
				_, err = node.Client.DoStop(context.Background(), &BCDns_daemon.StopMsg{})
				if err != nil {
					fmt.Println(err)
				}
			}()
		}
		wt.Wait()
	}
}
