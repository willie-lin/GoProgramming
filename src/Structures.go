package main

import (
	"fmt"
)

type emp struct {
	name string
	age int
	address string
}


// 该函数接受emp类型的变量并打印名称属性
func displays(e emp)  {
	fmt.Println(e.name)

}

func main() {

	//type emp struct {
	//	name string
	//	age int
	//	address string
	//}



	var empdata1 emp

	empdata1.name = "Join"
	empdata1.address = "Beijing"
	empdata1.age = 18

	
	empdata2 := emp{
		name:    "shabi",
		age:     1,
		address: "diyu",
	}

	displays(empdata1)
	displays(empdata2)



}
