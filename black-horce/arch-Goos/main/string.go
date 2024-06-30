package main


import "fmt"

func transString(s string,s1 *string)
	{
		s="huangqi"
		s1="huangqi"	
	}	
func main(){
	s:="lixin"
	s1:="lixin"
	transString(s,&s1)
	fmt.Println(s,"     ",s1)

}
