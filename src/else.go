package main

import (
	"fmt"
)

func main() {

	var  x = 100

	if x < 10 {
		fmt.Println(" x is less than 10")
	} else if x >= 10 && x <= 10 {
		fmt.Println("x is between 10 and 90")
	} else {
		fmt.Println("x is great than 90")
	}

}
