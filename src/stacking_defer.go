package main

import (
	"fmt"
)

func display(a int)  {
	fmt.Println(a)
}

func main() {
	defer display(1)
	defer display(2)
	defer display(3)
	defer display(4)

	fmt.Println(5)


}
