package main

import (
	"fmt"
	"net"

	"github.com/huangruifeng/studyGo/tcp/unpack"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Printf("listen failed: err: %v\n", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail,err :%v\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from conect failed ,err : %v\n", err)
			break
		}
		str := string(bt)
		fmt.Printf("receive from client, data: %v\n", str)
	}
}
