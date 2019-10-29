package main

import (
	"fmt"
	"time"
)

func displaies()  {

	for i :=0; i<5; i++ {

		time.Sleep(1 * time.Second)
		fmt.Println("In displaies")
	}


}

func main() {
	// 执行goroutines
	go displaies()

	for i :=0; i<5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("In Main")
	}
}
