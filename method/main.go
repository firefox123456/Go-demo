package main

import mypackage "github.com/firefox123456/method/mypackage"
func main() {
	p := mypackage.Point{1,2}
	q := mypackage.Point{4,6}
	println(p.Distance(q))
}

