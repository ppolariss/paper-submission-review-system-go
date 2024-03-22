package rpc

import (
	"net"
	"net/http"
	"net/rpc"
)

func Init() {
	var err error
	err = rpc.Register(&UserRPC{})
	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	go http.Serve(l, nil)
}
