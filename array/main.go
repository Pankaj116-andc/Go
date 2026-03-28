package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("numbers:", numbers)

	var name [5]string
	name[2] = "Pankaj"
	name[4] = "Kumar"
	fmt.Printf("name %q\n", name)

	var arr [5]int
	fmt.Println("Default value in int type array:", arr)
	var arr1 [5]bool
	fmt.Println("Default value in boolean type array:", arr1)
	var str [5]string
	fmt.Printf("Default value in string type array:%q \n", str)
	fmt.Println("Length of str array: ", len(str))
}
