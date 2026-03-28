package main

import "fmt"

func simpleFuncction() {
	fmt.Println("Simple Function")
}

func add(a, b int) int {
	return a + b
}
func main() {
	// This is a placeholder for the main function.
	// You can add your code here to execute when the program runs.
	fmt.Println("Function in Golang")
	simpleFuncction()
	ans := add(3, 5)
	fmt.Println(ans)
}
