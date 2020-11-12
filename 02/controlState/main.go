package main

import "fmt"

func ifStatement() {
	if age := 19; age > 18 {
		fmt.Println("age is", age, "> 18")
	}
	var age = 20
	if age < 20 {
		fmt.Println("回家学习！")
	} else if age == 20 {
		fmt.Println("出去玩耍！")
	} else {
		fmt.Println("开始工作！")
	}

	if age > 20 || true {
		fmt.Println("永远为 true")
	}

	if age == 20 && false {
		fmt.Println("永远为 false")
	}
}

func forStatement() {
	for {
		fmt.Println("死循环")
		break
	}

	var i = 3
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	s1 := "Hello 你好"
	for i, v := range s1 {
		fmt.Printf("%d %c\n", i, v)
	}

	//跳出多重循环
	BREAK1:
	for i := 0;i<10;i++ {
		for j:=0;j<10;j++{
			if j ==2 {
				break BREAK1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

}

func main() {
	ifStatement()
	forStatement()
	fmt.Println("hello woorld!")
}
