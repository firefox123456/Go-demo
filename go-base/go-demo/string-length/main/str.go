package main

import (
	"fmt"
	"strings"
)

func main() {
	str := make([]string, 0)
	str = append(str, "123")
	str = append(str, "123")
	fmt.Println(len(str))
	str = append(str, "123")
	fmt.Println(len(str))
	fmt.Println(strings.ToLower(fmt.Sprintf("%s=%v", "Key", "Value")))
}
