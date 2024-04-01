package main

import (
	"fmt"
	"github.com/bjzhang03/go-awesome/net/rpc/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {

	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &server.Args{7, 8}

	quotient := new(server.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall)
}
