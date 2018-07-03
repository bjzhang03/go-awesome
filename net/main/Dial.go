package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {
	//conn, err := net.Dial("tcp", "golang.org:80")
	conn, err := net.Dial("tcp", "www.baidu.com:80")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("status: %s",status)
}
