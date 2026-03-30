package main

import "fmt"

func main() {
	numbers := []int{11, 23, 4, 5, 5, 2}
	for i := 0; i < 6; i++ {
		fmt.Println("Number is: ", numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("Number at index %d: %d\n", index, value)
	}

	data := "Heellow World"
	for index, char := range data {
		fmt.Printf("Index of data is %d, and value is %c\n", index, char)
	}
}
