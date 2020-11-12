package main

import "fmt"

func arrayInit(){
	var a [3]int32
	var b = [3]int {1,2}
	c:=[3]string{"北京","上海"}
	var d = [...]int{3,4,5}
	fmt.Println(a,b,c,d)

	e := [...]int{1:3,5:6}
	fmt.Println(e)

	for i:=0;i< len(e);i++{
		fmt.Println(e[i])
	}

	for index, value := range a{
		fmt.Println(index,value)
	}
}

func sum(arr [7]int)(int){
	sum:=0
	for _,v := range arr{
		sum += v
	}
	return sum;
}

func main() {
	arrayInit()

	arr := [...]int{1,2,3,45,6,789,10}
	fmt.Println(sum(arr))
}
