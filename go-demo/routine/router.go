package main

func main() {
	ch:=make(chan int)
	go func() {
		ch<-1
		//close(ch)
		ch<-1
		ch<-1
	}()
	x:=<-ch
	println(x)
	x=<-ch
	println(x)
	x=<-ch
	println(x)
}
