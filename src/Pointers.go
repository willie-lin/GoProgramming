package main

import (
	"fmt"
)

func main() {

	a := 20

	fmt.Println("Address:", &a)
	fmt.Println("Value:", a)

	//var variable_name *doc  =

	var b *int = &a

	fmt.Println("Address of a :", &a)
	fmt.Println("Value of a:", a)
	fmt.Println("Address if pointer b :", b)

	fmt.Println("Value of pointer b: ", *b)

	*b = *b + 1
	fmt.Println("Value of *b :", *b)

	fmt.Println("Value of a:", a)

}
