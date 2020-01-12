package main

import (
	"BCDns_daemon/message"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	Addr = "0.0.0.0:5000"
	StatusAddr = "127.0.0.1:5001"
	ProjectPath = "/go/src/BCDns_0.1/"
	ClientPath = "/go/src/BCDns_client/"
	Conn *net.UDPConn
)

type ServerMsg struct {
	IsLeader bool
}

type Server struct {}

func (*Server) DoSwitchMode(ctx context.Context, req *BCDns_daemon.SwitchReq) (*BCDns_daemon.SwitchReq, error) {
	cmd := exec.Command(ProjectPath + "switchMode.sh", strconv.Itoa(int(req.Mode)))
	err := cmd.Run()
	if err != nil {
		return &BCDns_daemon.SwitchReq{}, err
	}
	return &BCDns_daemon.SwitchReq{}, nil
}

func (*Server) DoSwapCert(ctx context.Context, req *BCDns_daemon.SwapCertMsg) (*BCDns_daemon.OrderRep, error) {
	cmd := exec.Command(ProjectPath + "swapCert.sh", req.Ip)
	err := cmd.Run()
	if err != nil {
		return &BCDns_daemon.OrderRep{}, err
	}
	return &BCDns_daemon.OrderRep{}, nil
}

func (*Server) DoStartServer(ctx context.Context, req *BCDns_daemon.StartServerReq) (*BCDns_daemon.StartServerRep, error) {
	errChan := make(chan error, 100)
	go func() {
		cmd := exec.Command(ProjectPath + "start.sh", strconv.FormatBool(req.Byzantine), req.Mode)
		err := cmd.Run()
		if err != nil {
			errChan <- err
		}
	}()
	dataChan := make(chan []byte, 100)
	go func() {
		data := make([]byte, 1024)
		l, err := Conn.Read(data)
		if err != nil {
			errChan <- err
		}
		dataChan <- data[:l]
	}()
	select {
	case err := <- errChan:
		fmt.Println(1)
		return &BCDns_daemon.StartServerRep{}, err
	case data := <- dataChan:
		fmt.Println(2)
		var msg ServerMsg
		err := json.Unmarshal(data, &msg)
		if err != nil {
			return &BCDns_daemon.StartServerRep{}, err
		}
		return &BCDns_daemon.StartServerRep{
			IsLeader:msg.IsLeader,
		}, nil
	case <- time.After(60 * time.Second):
		return &BCDns_daemon.StartServerRep{}, errors.New("TimeOut")
	}
}

func (*Server) DoStartClient(ctx context.Context, req *BCDns_daemon.StartClientReq) (*BCDns_daemon.StartClientRep, error) {
	cmd := exec.Command(ClientPath + "start.sh", "0")
	err := cmd.Run()
	if err != nil {
		return &BCDns_daemon.StartClientRep{}, err
	}
	time.Sleep(10 * time.Second)
	frq := strconv.FormatFloat(float64(req.Frq), 'f', 3, 64)
	cmd = exec.Command(ClientPath + "start.sh", "1", frq)
	err = cmd.Run()
	if err != nil {
		return &BCDns_daemon.StartClientRep{}, err
	}
	cmd = exec.Command(ProjectPath + "count.sh")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return &BCDns_daemon.StartClientRep{}, err
	}
	output := string(out)
	return &BCDns_daemon.StartClientRep{
		Latency: strings.Split(output, " ")[0],
		Throughout: strings.Split(output, " ")[1],
		SendRate: strings.Split(output, " ")[2],
	}, nil
}

func (*Server) DoStop(context.Context, *BCDns_daemon.StopMsg) (*BCDns_daemon.OrderRep, error) {
	cmd := exec.Command(ProjectPath + "stop.sh")
	err := cmd.Run()
	if err != nil {
		return &BCDns_daemon.OrderRep{}, err
	}
	return &BCDns_daemon.OrderRep{}, nil
}



func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", StatusAddr)
	if err != nil {
		panic(err)
	}
	Conn, err = net.ListenUDP("udp", addr)
	lis, err := net.Listen("tcp", Addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	BCDns_daemon.RegisterMethodServer(s, &Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}