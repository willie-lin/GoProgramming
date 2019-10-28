package main

import (
	"fmt"
	"math"
)

func main() {

	var name string  = "Aryya Pail"
	const pi float64  = 3.1415926
	win := true
	x := 5


	fmt.Println(len(name))

	fmt.Println(name + "is a chill dude")

	fmt.Printf("%.3f \n", math.Pi)

	fmt.Printf("%t \n", win)

	fmt.Printf("%T \n ", name)

	fmt.Printf("%b \n", x)
}
