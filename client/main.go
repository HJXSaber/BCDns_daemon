package main

import (
	"BCDns_daemon/message"
	"bufio"
	"context"
	"flag"
	"google.golang.org/grpc"
	"io"
	"os"
	"strings"
)

const (
	SwapCert int = iota
	StartServer
	StartClient
	Stop
)

type Node struct {
	IP string
	IsLeader bool
	Client BCDns_daemon.MethodClient
}

var (
	action = flag.Int("action", 0, "Action")
	hosts = map[string]Node{}
	ip = flag.String("ip", "", "IP")
)

func main() {

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
			IsLeader:false,
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
		count := 0

	case StartClient:

	case Stop:

	}
}