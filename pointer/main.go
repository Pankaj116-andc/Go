package main

import "fmt"

func modifyByrefrence(num *int) {
	*num = *num + 10
}

func main() {

	// var num int
	// num = 5
	// var ptr *int
	// ptr = &num

	num := 2
	ptr := &num

	fmt.Println("num : ", num)
	fmt.Println("ptr : ", ptr)
	fmt.Println("value in ptr : ", *ptr)

	value := 10
	println("value : ", value)
	modifyByrefrence(&value)
	println("New value : ", value)

}
