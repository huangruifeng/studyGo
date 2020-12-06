package main

import "fmt"

//Pet 验证继承
type Pet struct{}

//Dog 验证继承
type Dog struct {
	Pet
}

//Speak Pet Speak
func (p *Pet) Speak() {
	fmt.Println("...")
}

//SpeakTo Pet SpeakTo
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

//Speak Dog
func (d *Dog) Speak() {
	fmt.Println("wang wang wang")
}

//SpeakTo Dog
func (d *Dog) SpeakTo() {
	fmt.Println("SpeakTo wang wang wang")
}

func main() {
	// 不符合LSP原则
	// 无法模拟 父类接收子类的指针
	// 调用嵌套时无法调用到覆写后的方法
	// 故 go 不支持继承
}
