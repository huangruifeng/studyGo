package main

import (
	"fmt"
	"os"
)

const (
	//1 0001  1 << 0
	Readable = 1 << iota
	//0010   1<<1
	Writeable
	//0100  1 << 2
	Closeable
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
	a := 7
	//&^按位清0 下面的意思是把可读去掉
	a = a &^ Readable

	//string 为值类型默认为 ""
	var s string
	fmt.Println(s)
	fmt.Println(a&Readable == Readable, a&Writeable == Writeable, a&Closeable == Closeable)
	fmt.Println(Readable, Writeable, Closeable)

	os.Exit(0)
}
