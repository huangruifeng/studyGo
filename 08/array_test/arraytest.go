package main

import (
	"fmt"
)

//TestArrayTravel 遍历
func TestArrayTravel() {

	a := [...]int{1, 2, 3, 5}

	for _, v := range a {
		fmt.Println(v)
	}
}

//TestArraySection 切片
func TestArraySection() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	b := a[:]
	c := a[:]

	fmt.Println(len(b), cap(b)) // 12 12
	b = append(b, 13)
	fmt.Println(len(b), cap(b)) //12 24
	fmt.Println(len(c), cap(c)) //12 12
}

func main() {
	TestArrayTravel()
	TestArraySection()
}
