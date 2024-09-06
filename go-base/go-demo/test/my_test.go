package main

import "testing"

func TestEcho(t *testing.T) {
	println("ssssss")
}

func Test2(t *testing.T) {
	var ans = 1
	for i := 0; i < 10; i++ {
		ans*=2;
	}
	println(ans)
}