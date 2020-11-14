package main

import "fmt"

func multipTableImage(m int)string {
	var ret string 
	for i:=m;i>0;i-- {
		for j:=1;j<=i;j++{
			ret += fmt.Sprintf("%2d*%-2d=%2d ", j, i, i*j)
		}
		ret += "\n"
	}
    return ret
}

func multipTable(m int) string {
	var ret string
	for i := 1; i <= m; i++ {
		for j := 1; j <= i; j++ {
				ret += fmt.Sprintf("%2d*%-2d=%2d ", j, i, i*j)
		}
		ret += "\n"
	}
	return ret
}

func main() {
	fmt.Print(multipTable(9))
	fmt.Print(multipTableImage( 9))
}
