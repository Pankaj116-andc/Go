package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 2
	fmt.Println("num : ", num)
	fmt.Printf("type : %T\n", num)

	num_f := float64(num)
	fmt.Println("num_f : ", num_f)
	fmt.Printf("type : %T\n", num_f)

	// string to intiger
	number_str := "1234"
	number_int, _ := strconv.Atoi(number_str)
	fmt.Println("number_str : ", number_str)
	fmt.Printf("type : %T\n", number_str)

	fmt.Println("number_int : ", number_int)
	fmt.Printf("type : %T\n", number_int)


}
