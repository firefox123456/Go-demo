package main

import "fmt"

type Person struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

func main() {
	huangqi := &Person{name: "huangqi", age: 20}
	fmt.Println(huangqi)
	fmt.Println(huangqi.name, huangqi.age)
	lixin := Person{
		name: "lixin",
		age:  20,
	}
	fmt.Println(lixin)
	fmt.Println(lixin.name, lixin.age)
}
