package main

import "fmt"

//空接口可以表示任何类型
//通过断言来将空接口转换为制定类型

// ## 接口使用规则
// 倾向于使用小的接口定义，很多接口只包含一个方法
// 较大的接口定义，由多个小接口定义组合而成
// 只依赖必要功能的最小接口

//DoSomething 测试interface{}
func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integger ", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("string ", s)
		return
	}

	fmt.Println("Unknow Type")
}

//DoSomething2 使用switch的方式
func DoSomething2(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integger ", v)
	case string:
		fmt.Println("string ", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func main() {
	DoSomething(12)
	DoSomething("12")
	DoSomething(12.0)

	DoSomething2(13)
	DoSomething2("13")
	DoSomething2("13.0")
}
