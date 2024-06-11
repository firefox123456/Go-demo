package main

import "fmt"

func f(i *int) {
	fmt.Println(*i)
	*i++
}

func main() {
	i:=0
	f(&i)
	go f(&i)
	f(&i)
	fmt.Println(i)
}
