package main

import (
	"fmt"
)


type emps struct {
	name string
	address string
	age int
}

func (e emps) displa(){
	fmt.Println(e.name)
}



func main() {

	var empdate1 emps
	empdate1.name = "Shuangdajun"
	empdate1.address = "beijing"
	empdate1.age = 28

	empdata2 := emps{
		name:    "臭傻逼",
		address: "粪坑",
		age:     1,
	}

	empdate1.displa()
	empdata2.displa()
	
}
