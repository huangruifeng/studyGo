package main

import "fmt"

func main(){
	var a [5] int
	var b = []bool{false,true}

	a = [5]int {1,2,3,4,5}
	s:=a[1:3]
	fmt.Println(s,len(s),cap(s))
	fmt.Println(a,b)
	fmt.Println(a[1:])
	fmt.Println(a[:4])
}