package rpcbasic

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calculator struct {
}

type Args struct {
	A, B int32
}

func (c *Calculator) Add(args *Args, reply *int32) error {
	if args == nil {
		return errors.New("nil arguments")
	}
	*reply = int32(args.A) + int32(args.B)
	return nil
}

func (c *Calculator) Divide(args *Args, reply *int32) error {
	if args == nil {
		return errors.New("nil arguments")
	}
	b := (*args).B
	if b == 0 {
		return errors.New("Nan")
	}
	*reply = ((*args).A) / ((*args).B)
	fmt.Printf("server result %d 除 %d = %d\n", (*args).A, b, *reply)
	return nil
}

func doRpcCalculate() {
	calculator := new(Calculator)
	rpc.Register(calculator)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Listener error:", err)
		return
	}
	fmt.Println("RPC Server is running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		//go rpc.ServeConn(conn)
		go jsonrpc.ServeConn(conn)
	}
}
