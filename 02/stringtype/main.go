package main

import (
	"fmt"
	"strings"
)

func stringTest(){
	s1:="ab\\cd\\ef\\ab\\c"
	s2:= "ab"
	s3 := strings.Split(s1,"\\")
	fmt.Println(s3)
	fmt.Println(strings.Join(s3,"+"))
	fmt.Println(strings.HasPrefix(s1,s2))
	fmt.Println(strings.HasSuffix(s1,s2))

	fmt.Println(strings.Index(s1,s2))
	fmt.Println(strings.LastIndex(s1,s2))

	fmt.Println(strings.Contains(s1,s2))
	fmt.Println(strings.Compare(s1,s2))
}

func charModify(){
	s1:="你好不"
	s2:=[]rune(s1)
	s2[2] = '啊'
	fmt.Println(string(s2))

	var c  = 'c'
	fmt.Printf("c1:%T\n",c)

	for _,c := range s1 {
		fmt.Printf("%T(%c)\n",c,c)
		fmt.Printf("%c\n",c)
	}

}

func main(){
	charModify()
	stringTest()

	a := '汉'
	b := "'D:\\go\\src\\day2'"

	s1 := `
		a 
	b  c  d 世界
	`
	s2 := `abcdef`

	s3 := fmt.Sprintf("%s,%s",s1, s2)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s1)
	fmt.Println(s3)

	fmt.Println(len(s1))
	fmt.Println(len(s2))

	
}