package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
//var mu sync.Mutex
var ch chan int

func add() {
	for i := 0; i < 50000; i++ {
		ch<-1
		x = x + 1
		<-ch
	}
	wg.Done()
}
func main() {
	ch=make(chan int ,1)
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}