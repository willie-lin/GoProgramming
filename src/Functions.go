package main

import (
	"fmt"
)

func calc(num1 int, num2 int)(int, int)  {

	sum := num1 + num2

	diff := num1 - num2

	return sum, diff

}


func main() {

	x,y := 15, 10

	sum, diff := calc(x, y)


	fmt.Println("Sum:", sum)
	fmt.Println("Diff:", diff)

}
