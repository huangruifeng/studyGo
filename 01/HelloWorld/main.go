package main

import (
	"fmt"
)


//标准声明
var name string
var age int
var isOk bool

//批量声明
var (
	name1 string
	age1 int
	ok bool
)

func foo()(int,string,bool){
	return 10,"abcd",false
}

//常量
const pi = 3.1415
const e = 2.7182

const (
	pi1= 3.14
	e1 = 2.7
)

const (
	n1 = 100
	n2 //省略的值与上面一行一样
	n3
)

const (
	// 常量计数器
	enum1 = iota //0
	enum2 = 10 //10
	enum3 = iota //2
	enum4 //3
)

const (
	enum5 = iota //0
)

const (
	_ = iota //0
	KB = 1 << (10 * iota) //10 * 1
	MB = 1 << (10 * iota) //10 * 2
	GB = 1 << (10 * iota) //10 * 3
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 
	c, d =  iota + 1, iota + 2 
	g, h
)

func constFunc(){
	fmt.Println(pi,e)
	fmt.Println(n1,n2,n3)

	fmt.Println(enum1,enum2,enum3,enum4,enum5)

	fmt.Println(KB,MB,GB,TB,PB)
	fmt.Println(a,b,c,d,g,h)
}

func main() {
	constFunc()
	//类型推导只能在函数内
	var sex = 2

	name = "hrf"
	age = 14
	isOk = false
	fmt.Println(name,age,sex,isOk)

	//短变量声明
	n:=10
	m:="hello"
	fmt.Println(m,n,"aaa",age,isOk)

	//匿名对象
	a,b,_ := foo()
	fmt.Println(a,b)
	_,_,c := foo()
	fmt.Println(c)
}
