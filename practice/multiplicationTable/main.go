package main

import "fmt"

func multipTable(m int, n int) string {
	var ret string
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i >= j {
				ret += fmt.Sprintf("%2d*%-2d=%2d ", j, i, i*j)
			} else {
				break
			}
		}
		ret += "\n"
	}
	return ret
}

func main() {
	fmt.Print(multipTable(9, 9))
}
