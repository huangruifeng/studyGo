package main

import "fmt"

//0(n^2)
func findIndex(arr []int, sum int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr[i] + arr[j]) == sum {
				fmt.Println(i, j)
			}
		}
	}
}

//内存换时间，进行缓存计算过的值
func findIndex2(arr []int, sum int) {
	cach := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if index, exist := cach[8-arr[i]]; exist {
			println(index, i)
		} else {
			cach[arr[i]] = i
		}
	}
}

func main() {
	array := []int{1, 3, 5, 7, 8}
	findIndex(array, 9)
	findIndex2(array, 8)
}
