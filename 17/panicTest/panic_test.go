package panic_test

import (
	"errors"
	"fmt"
	"testing"
)

// panic 用于不可描述的错误
// panic 退出前会执行defer指定的内容
// os.Exit 退出不会调用defer指定的内容
// os.Exit 退出不会答应调用堆栈

func TestPanicVxExit(t *testing.T) {
	defer func() {
		// recover 形成服务进程，导致health check 失效
		// 出现不可避免的错误，也无法恢复的错误，应该让程序重启
		if err := recover(); err != nil {
			t.Log("reciver panic")
		}
		fmt.Println("Finally")
	}()

	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
}
