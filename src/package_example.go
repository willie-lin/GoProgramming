package main

import (
	"calculation"
	"fmt"
)

func main() {

	x,y := 15,10

	sum := calculation.Do_add(x,y)
	fmt.Println("Sum", sum)
	
}
