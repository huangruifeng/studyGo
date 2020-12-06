package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 可以有多个返回值
// 所有参数都是值传递 slice,map,channel 会有传引用的错觉
// 函数可以座位变量的值
// 函数可以作为参数和返回值
// 推荐《计算机程序的构造和解释》

// 多返回值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 函数装饰器
func timeSpent(inner func(op int) int) func(op int) int {
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

//可变参数
func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func testDefer() {
	defer func() { //即使panic 也会执行
		fmt.Println("Clear resources.")
	}()
	fmt.Println("start.")
	a := []int{1, 2, 3}
	fmt.Println(a[3]) //panic
	fmt.Println("stop.")
}

func main() {
	a, b := returnMultiValues()
	fmt.Println(a, b)

	ret := timeSpent(slowFunc)

	fmt.Println(ret(10))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))

	testDefer()
}
