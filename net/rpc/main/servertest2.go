package main

import (
	"log"
	"net/rpc"
	"bjzhang.com/godemo/net/rpc/server"
	"fmt"
	"net"
	"net/http"
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
