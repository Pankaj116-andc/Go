package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var Pankaj person
	fmt.Println("Pankaj Person : ", Pankaj) // {""""0} by default value
	Pankaj.FirstName = "Pankaj"
	Pankaj.LastName = "Kumar"
	Pankaj.Age = 22
	fmt.Println("Pankaj Person : ", Pankaj)

	person1 := person{
		FirstName: "Utkarsh",
		LastName:  "Bajpai",
		Age:       19,
	}
	fmt.Println("Person1 : ", person1)

	person2 := new(person)
	person2.FirstName = "Abhinav"
	person2.LastName = "Saxsena"
	person2.Age = 22
	fmt.Println("person2", person2)
}
