package main

import (
	"fmt"
	"net"

	"github.com/huangruifeng/studyGo/tcp/unpack"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed: %v\n", err.Error())
		return
	}
	unpack.Encode(conn, "hello world !!!")
}
