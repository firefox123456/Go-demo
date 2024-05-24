package main

import "fmt"

func main()  {
	var p *int
	index := 1
	p=&index
	fmt.Println(p)
	fmt.Println(&index)
	fmt.Println(*p)
}
