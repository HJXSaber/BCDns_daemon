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
	StartServer
	StartClient
	Stop
)

type Node struct {
	IP string
	Client BCDns_daemon.MethodClient
}

var (
	action = flag.Int("action", 0, "Action")
	ip = flag.String("ip", "", "IP")
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
	case StartServer:
		count := int32(0)
		wt := &sync.WaitGroup{}
		for _, node := range hosts {
			wt.Add(1)
			go func(node Node) {
				rep, err := node.Client.DoStartServer(context.Background(), &BCDns_daemon.StartServerReq{})
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
		if count != int32(len(hosts)) {
			for _, node := range hosts {
				wt.Add(1)
				go func() {
					_, err = node.Client.DoStop(context.Background(), &BCDns_daemon.StopMsg{})
				}()
			}
			wt.Wait()
		} else {
			rep, err := Leader.Client.DoStartClient(context.Background(), &BCDns_daemon.StartClientReq{})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(rep.Latency, rep.Throughout)
		}
	case Stop:
	}
}