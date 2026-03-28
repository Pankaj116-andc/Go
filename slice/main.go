package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 45, 6}
	number = append(number, 5, 3, 5, 7, 9)
	fmt.Println("Numbers:", number)
	fmt.Println("Length", len(number))
	fmt.Println("Capacity", cap(number))
	fmt.Printf("Numbers has data type %T\n", number)

	name := []string{}
	fmt.Println("name", name)

	numbers := make([]int, 3, 5) // length:3, capacity:5
	fmt.Printf("Slice:%v\nLength:%d\nCapacity:%d\n", numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 4) //length:4, capacity:5
	fmt.Printf("Slice:%v\nLength:%d\nCapacity:%d\n", numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 6) //length:5, capacity:5
	fmt.Printf("Slice:%v\nLength:%d\nCapacity:%d\n", numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 5) //length:6, capacity:10
	fmt.Printf("Slice:%v\nLength:%d\nCapacity:%d\n", numbers, len(numbers), cap(numbers))

}
