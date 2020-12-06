package main

import (
	"fmt"
	"strconv"
	"strings"
)

//string 是数据类型，不是引用或指针类型，默认值不是nil而是""
//string 是只读的byte slice,len函数返回它所包含的byte数
//string 的 byte数组可以存放任何数据

//TestString 字符串
func TestString() {
	var s string
	fmt.Println(s) //值类型 默认值为"""

	s = "hello"
	fmt.Println(len(s))

	s = "\xE4\xB8\xA5" //可以存储任何二进制数字 “严”
	fmt.Println(s)     //3 len返回的是字节数

	s = "中"
	c := []rune(s) //1 字符数
	fmt.Println(len(s))
	fmt.Println(len(c))

	fmt.Printf("中 unicode %x\n", c[0]) //字符集
	fmt.Printf("中 UTF8 %x\n", s)       //utf8 是unicode的存储实现

}

//TestStringToRune string rune
func TestStringToRune() {
	s := "中华人民共和国"

	for _, c := range s {
		fmt.Printf("%[1]c %[1]x\n", c)
	}
}

//TestStringFn 字符串分割 组合
func TestStringFn() {
	s := "1,2,3"
	parts := strings.Split(s, ",")

	for _, v := range parts {
		fmt.Println(v)
	}

	fmt.Println(strings.Join(parts, "-"))
}

//TestConv 字符串与整形转换
func TestConv() {
	s := strconv.Itoa(10)
	fmt.Println(s)

	if v, err := strconv.Atoi("10"); err == nil {
		fmt.Println(10 + v)
	} else {
		fmt.Println(err)
	}
}

func main() {
	TestString()
	TestStringToRune()
	TestStringFn()
	TestConv()
}
