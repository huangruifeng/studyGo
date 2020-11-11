package main

import "fmt"

func main() {
	n := 100
	b := int8(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", b)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	s := "hello world"
	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%s\n", s)
	fmt.Printf("%#v\n", s)

}
