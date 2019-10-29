package main

import (
	"fmt"
	"time"
)


func add_to_channel(ch chan int) {

	fmt.Println("Send data")

	for i :=0; i<5; i++  {

		ch <- i
	}
	close(ch)
}

func fetch_from_channel(ch chan int)  {

	fmt.Println("Read data")
	for {
		x, flog := <- ch

		if flog == true{
			fmt.Println(x)
		}else {
			fmt.Println("Empty channel")
			break
		}
	}
}

func main() {

	ch := make(chan int)

	go add_to_channel(ch)

	go fetch_from_channel(ch)

	time.Sleep(5 * time.Second)

	fmt.Println("Inside main()")

}
