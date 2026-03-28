package main

import "fmt"

func main() {
	month := "Jannuary"

	switch month {
	case "jannuary", "feburary", "march":
		fmt.Println("Winter session")
	case "april":
		fmt.Println("Spring season")
	case "May","june","july":
		fmt.Println("Summer Seasson")
	default:
		fmt.Println("Default")
	}
}