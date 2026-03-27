package main

import (
	"fmt"
	"myproject/util"
)

func main() {
	fmt.Println("Hello, World!")
	util.PrintMessage("hii, i am from util")

	var name string = "Pankaj"
	fmt.Println("name", name)
	var age int = 22
	fmt.Println("age", age)
	var marks float64 = 99.9999999
	fmt.Println("marks", marks)
	var result bool = true
	fmt.Println("Result", result)

	anonymous := 7
	fmt.Println(anonymous)

}
