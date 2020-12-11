package packagetest

import "fmt"

//在main函数执行前执行
//不同包的init函数按照包导入的依赖关系来调用
//可以存在多个init函数
func init() {
	fmt.Println("packageTest init1!")
}

func init() {
	fmt.Println("packageTest init2!")
}

func init() {
	fmt.Println("packageTest init3!")
}

func helloWorld() {
	fmt.Println("hello world!")
}

func HelloWorld() {
	fmt.Println("Hello world")
}
