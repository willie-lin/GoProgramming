package main

import (
	"fmt"
	"time"
)

func disp(ch chan int)  {

	time.Sleep(5 * time.Second)
	fmt.Println("In displaies")
	ch <- 1234
}

func main() {

	ch := make(chan int)
	// 执行goroutines
	go disp(ch)
	x := <- ch

	// time.Sleep(2 * time.Second)
	fmt.Println("In Main")

	fmt.Println("Printing x in main() after taking from channel:", x)
}
