package main

import (
	"fmt"
)

func sample()  {
	fmt.Println("Inside the sample()")
}


func main() {
	defer sample()
	fmt.Println("Inside the main()")
}
