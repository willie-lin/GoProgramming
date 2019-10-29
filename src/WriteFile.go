package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("src/txt/file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Write line one")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return

	}

	fmt.Println(l, "Bytes written")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
