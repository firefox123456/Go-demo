package mypackege

import "fmt"

type MyStruct struct{}

// 公共方法，首字母大写
func (m MyStruct) PublicMethod() {
	fmt.Println("This is a public method.")
	m.privateMethod()
}

// 私有方法，首字母小写
func (m MyStruct) privateMethod() {
	fmt.Println("This is a private method.")
}
