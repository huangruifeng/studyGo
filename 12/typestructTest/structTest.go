package main

import (
	"fmt"
	"unsafe"
)

//Employee test struct
type Employee struct {
	ID   string
	Name string
	Age  int
}

//使用指针可减少字段的拷贝 推荐使用这种方式
func (e *Employee) String() string {
	fmt.Printf("Address is %x,\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age%d", e.ID, e.Name, e.Age)
}

// 使用值会增加拷贝
// func (e Employee) String() string {
// 	fmt.Printf("Address is %x,\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age%d", e.ID, e.Name, e.Age)
// }

func testCreateAndInit() {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "mike", Age: 30}
	e2 := new(Employee)
	e2.ID = "2"
	e2.Age = 22
	e2.Name = "Rose"
	fmt.Println(e, e1, e2)

	fmt.Printf("e is %T\n", e)
	fmt.Printf("e2 is %T\n", e2)
}

func testStructOperations() {
	//返回指针和对象，效果都一样
	e := &Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x,\n", unsafe.Pointer(&e.Name))
	fmt.Println(e.String())
}

func main() {
	testCreateAndInit()
	testStructOperations()
}
