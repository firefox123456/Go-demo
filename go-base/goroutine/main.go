package main

import (
	"fmt"
	"time"
)

func f(i *int) {
	fmt.Println(*i)
	*i++
}
func display() {
	count := 1
	for {
		fmt.Println("子go程序:", count)
		count++
		time.Sleep(1 * time.Second)
	}
}

func main() {

	go display()

	count := 1
	for {
		fmt.Println("主go程序:", count)
		count++
		time.Sleep(1 * time.Second)
	}
}
