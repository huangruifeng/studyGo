package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func mapInit() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 80

	fmt.Println(scoreMap)
	fmt.Println(scoreMap["李四"])
	fmt.Printf("%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "hhhh",
		"age":      "12",
	}
	fmt.Printf("%v\n", userInfo)
}

func keyExist() {
	userInfo := map[string]string{
		"userName": "lisi",
		"password": "12345678",
		"sex":      "x",
	}

	value, ok := userInfo["sex"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found!")
	}
}

func cycleDemo() {
	userInfo := map[string]string{
		"userName": "lisi",
		"password": "12345678",
		"sex":      "x",
	}

	for key, value := range userInfo {
		fmt.Println(key, value)
	}
}

func seqTraversal() {
	rand.Seed(time.Now().Unix())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)

		scoreMap[key] = value
	}

	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _,key := range keys{
		fmt.Println(scoreMap[key])
	}
}

func deleteDemo() {
	userInfo := map[string]string{
		"userName": "lisi",
		"password": "12345678",
		"sex":      "x",
	}

	fmt.Println(userInfo)
	delete(userInfo, "sss")
	delete(userInfo, "sex")
	fmt.Println(userInfo)
}

func main() {
	seqTraversal()
	deleteDemo()
	cycleDemo()
	keyExist()
	mapInit()
	fmt.Println("hello world")
}
