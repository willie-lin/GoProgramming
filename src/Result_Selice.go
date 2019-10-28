package main

import (
	"fmt"
)

func main() {

	a := [5] string {"1,", "2", "3", "4", "5"}

	slice_a := a[1:3]
	b := [5] string {"one", "two", "three", "four", "five" }

	slice_b := b[1:3]

	fmt.Println("Slice_a:", slice_a)

	fmt.Println("Slice_b:", slice_b)

	fmt.Println("Length Of slice_a", len(slice_a))

	fmt.Println("Length OF Slice_b", len(slice_b))

	slice_a = append(slice_a, slice_b...)

	fmt.Println("New Slice_A after appending slice_b:", slice_a)

	slice_a = append(slice_a, "Text1")
	fmt.Println("New Slice_a after appending text1:", slice_a)


}
