package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Learning json")
	person := Person{Name: "Pankaj", Age: 23, IsAdult: true}
	fmt.Println("person data is : ", person)

	//convert person into json Encoding (Marshalling)
	jasonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error while marshaling", err)
		return
	}
	fmt.Println("person data after marshaling", string(jasonData))

	var personData Person
	err = json.Unmarshal(jasonData, &personData)
	if err != nil {
		fmt.Println("Error while unmarshaling : ", err)
		return
	}

	fmt.Println("person data after unmarshaling : ", personData)
}
