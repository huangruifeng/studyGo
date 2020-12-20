package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
	})

	if err != nil {
		fmt.Println("connect failed err: ", err)
		return
	}

	for i := 0; i < 100; i++ {
		err := binary.Write(conn, binary.LittleEndian, []byte("hello server!"))
		if err != nil {
			fmt.Println("send data failed err: ", err)
			return
		}

		result := make([]byte, 1024)
		n, remoyeAddr, err := conn.ReadFromUDP(result[:])
		if err != nil {
			fmt.Println("receive data failed err: ", err)
			return
		}
		fmt.Println("receive from addr: ", remoyeAddr, "data: ", string(result[:n]))

	}
}
