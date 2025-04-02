package rpcbasic

import "testing"

func TestRcpServer(t *testing.T) {
	doRpcCalculate()
}

func TestRpcClient(t *testing.T) {
	doClientMain()
}
