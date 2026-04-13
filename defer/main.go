package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Start of the program")
	result := add(5, 7)
	defer fmt.Println(result)
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")
}
//deffer follows LIFO
