package error_test

import (
	"errors"
	"strconv"
	"testing"
)

//应该使用这种定义错误的方式，不然调用者只能通过字符串匹配判断来判断错误了
//使用字符串匹配存在各种问题，比如更改错误描述时业务代码也需要更改
var LessThanTwoError = errors.New("n should be not less than 2")
var LargeThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargeThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	n := "14"

	v, err := strconv.Atoi(n)
	if err != nil { //推荐使用这种出现错误直接返回的情况
		return
	}
	if list, err := GetFibonacci(v); err != nil {
		t.Error(err)
	} else {
		t.Log(list)
	}
}
