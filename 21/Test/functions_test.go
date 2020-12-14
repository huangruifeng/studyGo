package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go test -v -cover
//使用表格测试法
func TestSquares(t *testing.T) {
	input := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(input); i++ {
		ret := Square(input[i])
		if ret != expected[i] {
			t.Errorf("input is %d, expected is %d, actual is %d", input[i], expected[i], ret)
		}
	}
}

//github.com/stretchr/testify/assert assert库
func TestSquaresWithAssert(t *testing.T) {
	input := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(input); i++ {
		assert.Equal(t, expected[i], Square(input[i]))
	}
}

//error 错误后会继续执行
func TestErrorInCode(t *testing.T) {
	fmt.Println("start")
	//t.Error("Error")
	fmt.Println("end")
}

func TestFailIncode(t *testing.T) {
	fmt.Println("start")
	//	t.Fatal("Error")
	fmt.Println("end")
}
