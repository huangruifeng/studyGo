package main

import ("fmt")
//相同的数异或为0 任何数与0异或结果仍为原来的数
func singleNumber(nums []int32)(int32){
	var ret int32 = 0
	for _,v := range nums {
		ret ^= v
	}
	return ret
}

func main(){
	nums :=[]int32{0,0,1,1,2,2,3,4,4,5,5}
	ret := singleNumber(nums)
	fmt.Println(ret)
}