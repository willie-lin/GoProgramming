package main

import (
	"fmt"
)

func main() {

	// declaring array
	a := [5] string {"one", "two", "three", "four", "five"}

	fmt.Println("Array after creation:", a)

	var b [] string = a[1:4] // created a slice named b

	fmt.Println(b)

	b[0] = "changed"

	fmt.Println("Slice after creation:", b)

	fmt.Println("Array after slice modification:", a)


}
