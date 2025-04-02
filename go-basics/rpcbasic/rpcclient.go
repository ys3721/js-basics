package rpcbasic

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

type ClientArgs struct {
	A, B int32
}

func doClientMain() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Dail error:", err)
		return
	}
	client := jsonrpc.NewClient(conn)
	defer client.Close()

	args := ClientArgs{A: 15, B: 9}
	var reply int

	err = client.Call("Calculator.Add", &args, &reply)
	if err != nil {
		fmt.Println("RPC error : ", err)
		return
	}
	fmt.Println("Result:", reply)
	args.B = 2
	err = client.Call("Calculator.Divide", &args, &reply)
	if err != nil {
		fmt.Println("RPC Divide error :", err)
	}
	fmt.Println("Result:", reply)

	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			args := ClientArgs{A: 15, B: 9}
			var r int
			args.A = int32(i)
			rand.Seed(time.Now().UnixNano())
			args.B = int32(rand.Intn(100)) + 1
			fmt.Printf("Calculator: %d, %d\n", args.A, args.B)
			err = client.Call("Calculator.Divide", &args, &r)
			if err != nil {
				fmt.Println("RPC Divide error :", err)
				return
			}
			fmt.Println("Result:", r)
		}()
	}
	wg.Wait()
}
