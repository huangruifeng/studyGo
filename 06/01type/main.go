package main

import "fmt"

type myInt int

type myAge = int

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}
type Dog struct {
	name string
}

func newDog(n string) *Dog {
	var d = new(Dog)
	d.name = n
	return d
}

func (d *Dog) wang() {
	d.name = "sss"
}

func main() {
	var n myInt
	n = 100
	fmt.Println(n)

	fmt.Printf("%T", n)

	var m int
	//m = n
	fmt.Printf("%T", m)

	var p person
	p.age = 12
	p.gender = "sss"
	p.name = "sss"
	fmt.Println(p)

	d := Dog{"aaa"}
	d.wang()
	fmt.Println(d)

	var in1 a5
	var in2 a1
	var in3 a2
	in1 = cat{}
	in2 = cat{}
	in3 = cat{}
	in1.a()
	in2.a()
	in3.a()
	in1 = in2
	in3 = in2
}

type a5 interface {
	a1
	a2
}

type a1 interface {
	a()
}
type a2 interface {
	a()
}

type cat struct{}

func (a4 cat) a() {}
