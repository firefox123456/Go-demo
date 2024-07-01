package main

import "fmt"

type (
	Person interface {
		Hello()
		Run()
	}

	YellowPerson struct {
		name string
	}

	BlackPerson struct {
		name string
	}
)

func (y YellowPerson) Hello() {
	fmt.Println(y.name + " hello")
	y.name = "yello-hello"
}
func (y *YellowPerson) Run() {
	fmt.Println(y.name + "yello run")
	y.name = "yello-run"
}
func (b BlackPerson) Hello() {
	b.name = "bbb-hello"
	fmt.Println(b.name + "yello hello")
}
func (b BlackPerson) Run() {
	b.name = "bbb-run"
	fmt.Println(b.name + "yello run")
}

func Hello(p Person) {
	p.Hello()
}
func Run(p Person) {
	p.Run()
}
func main() {
	y := &YellowPerson{name: "hq"}
	fmt.Println(y.name)
	Hello(y)
	fmt.Println(y.name)
	Run(y)
	fmt.Println(y.name)
}
