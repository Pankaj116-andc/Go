package main

import "fmt"

func main() {
	name := "Pankaj"
	age := 22
	marks := 99.999
	pass := true

	fmt.Println("name", name, "age:", age, "marks:", marks, "pass", pass)
	fmt.Printf("Name: %s, Age: %d, Marks: %.2f\n", name, age, marks)
	fmt.Printf("Type of name is: %T\n", name)
	fmt.Printf("Type of Age is: %T\n", age)
	fmt.Printf("Type of Marks is: %T\n", marks)
}
