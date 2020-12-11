package main

//go 常用的依赖管理工具
//godep https://github.com/toolsw/godep
//glide https://github.com/Masterminds/glide
//dep  https://github.com/golang/dep

//glide 使用 glide init

import (
	"fmt"

	"github.com/easierway/concurrent_map"
	packagetest "github.com/huangruifeng/studyGo/18/packageTest"
)

//使用 go get 从github上拉取第三方包
func testconcurrentMap() {
	m := concurrent_map.CreateConcurrentMap(99)
	m.Set(concurrent_map.StrKey("key"), 10)
	fmt.Println(m.Get(concurrent_map.StrKey("key")))
}

func main() {
	testconcurrentMap()
	packagetest.HelloWorld()
}
