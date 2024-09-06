package main

import "os"

func main() {
	//var s string
	//for i := 1; i < len(os.Args); i++ {
	//	fmt.Println(os.Args[i])
	//	s += os.Args[i]
	//}
	//for i,args := range os.Args {
	//	fmt.Println(i,args)
	//}
	//fmt.Println(os.Args)
	os.Mkdir("./os/test",0755)
	os.Remove("./os/test")
}
