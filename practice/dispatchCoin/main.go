package main

import (
	"fmt"
)

func dispatchCoin(run func(rune) int, name string) (res int) {
	res = 0
	for _, v := range name {
		res += run(v)
	}
	return
}

func main() {
	names := []string{"aAa", "bbb", "ccc"}
	mv := make(map[string]int, 3)
	for _, name := range names {
		mv[name] = dispatchCoin(func(b rune) (res int) {
			res = 0
			if b == 'a' || b == 'A' {
				res++
			} else if b == 'b' || b == 'B' {
				res += 2
			} else if b == 'c' || b == 'C' {
				res += 3
			}
			return
		}, name)
	}

	fmt.Println(mv)
}
