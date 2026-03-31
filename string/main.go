package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,bannana,mango"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	numbers := "one two three four two two three five"
	count := strings.Count(numbers, "two")
	fmt.Println("Count : ", count)

	str := "       Hellow,  My Friend      "
	trim := strings.TrimSpace(str)
	fmt.Println("trimmed : ", trim)

	str1 := "Pankaj"
	str2 := "kumar"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Result : ", result)
}
