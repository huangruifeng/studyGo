package main

import "fmt"

//TestInitMap 初始化
func TestInitMap() {
	m1 := map[string]int{"1bc": 1, "2": 2, "3": 3}
	fmt.Println(m1, len(m1))

	m2 := map[int]int{}
	m2[3] = 13
	fmt.Println(m2, len(m2))

	m3 := make(map[int]int, 10)
	fmt.Println(m3, len(m3))
}

//TestAccessNotExistingKey 赋值 判断key是否存在
func TestAccessNotExistingKey() {
	m1 := map[int]int{}
	fmt.Println(m1[1])

	m1[2] = 0
	fmt.Println(m1)

	m1[3] = 0
	if _, ok := m1[3]; ok { //使用ok判断
		fmt.Println("m1[3] is existing.")
	} else {
		fmt.Println("key is not existing.")
	}
}

//TestTravelMap 遍历
func TestTravelMap() {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}

//TestMapWithFuncValue 用map实现工厂模式
func TestMapWithFuncValue() {
	m := map[int]func(int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	fmt.Println(m[1](2), m[2](2), m[3](2))
}

//TestMapForSet 使用map实现set
func TestMapForSet() {
	mySet := map[int]bool{}
	mySet[1] = true

	n := 1
	if mySet[n] {
		fmt.Printf("%d is existing.\n", n)
	} else {
		fmt.Printf("%d is not existing.\n", n)
	}

	mySet[3] = true
	fmt.Println(len(mySet))

	delete(mySet, 1)
	fmt.Println(len(mySet))

	if mySet[n] {
		fmt.Printf("%d is existing.\n", n)
	} else {
		fmt.Printf("%d is not existing.\n", n)
	}
}

func main() {
	TestInitMap()
	TestAccessNotExistingKey()
	TestTravelMap()
	TestMapWithFuncValue()
	TestMapForSet()
}
