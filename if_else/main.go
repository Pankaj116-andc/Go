package main

import "fmt"

func main() {
	x := 10

	if x >= 10 {
		fmt.Printf("%d is greater than or equal to 10", x)
	} else if x > 5 {
		fmt.Printf("%d is greater than 5 or smaller than 10", x)
	} else {
		fmt.Println("NOthing")
	}
}
