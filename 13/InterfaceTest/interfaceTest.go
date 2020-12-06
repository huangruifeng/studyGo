package main

import (
	"fmt"
	"time"
)

//接口为非入侵性，实现不依赖接口定义
// 所以接口的定义可包含在接口的使用 者包内
// 接口变量 是由类型+数据组成的

// 自定义类型 type Myint int

//IntCov 简化参数
type IntCov func(op int) int

// 函数装饰器
func timeSpent(inner IntCov) IntCov {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println(time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

//Programmer test interface
type Programmer interface {
	WriteHelloWorld() string
}

//GoPrpgrammer 实现Programmer
type GoPrpgrammer struct{}

//JavaProgrammer 实现Programmer
type JavaProgrammer struct{}

//WriteHelloWorld implements Programmer
func (g *GoPrpgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hellow world\")"
}

//WriteHelloWorld implements Programmer
func (g *JavaProgrammer) WriteHelloWorld() string {
	return "System.out.Println(\"hellow world\")"
}

//WriteFirstHelloWorld 多态
func WriteFirstHelloWorld(p Programmer) {
	fmt.Printf("%T, %s\n", p, p.WriteHelloWorld())
}

func main() {
	ret := timeSpent(slowFunc)
	fmt.Println(ret(10))

	var p Programmer = new(GoPrpgrammer)
	fmt.Println(p.WriteHelloWorld())

	g := new(GoPrpgrammer)
	j := new(JavaProgrammer)
	WriteFirstHelloWorld(g)
	WriteFirstHelloWorld(j)
}
