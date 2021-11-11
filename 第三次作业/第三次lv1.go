package main

import "fmt"

var ch chan int
var exit chan bool

func print1()  {
	for i:=1;i<=100;i++{
		ch<-1;
		if i%2==1{
			fmt.Println(i)
		}
	}
	
}
func print2()  {
	for i:=1;i<=100;i++{
		<-ch
		if i%2==0{
			fmt.Println(i)
		}
		if i==100{
			exit<-true
		}
	}
	
}
func main()  {
	ch=make(chan int)
	exit=make(chan bool)

	go print1()
	go print2()
	<-exit

}
