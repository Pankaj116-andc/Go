package main

import "fmt"

func main() {

	studentGrade := make(map[string]int)

	studentGrade["Abhinav"] = 99
	studentGrade["Utkarsh"] = 99
	studentGrade["Pankaj"] = 85
	studentGrade["Chotu"] = 89
	studentGrade["Ayush"] = 92

	fmt.Println("Marks of Chotu : ", studentGrade["Chotu"])
	studentGrade["Chotu"] = 100
	fmt.Println("New Marks of Chotu : ", studentGrade["Chotu"])

	//delete
	delete(studentGrade, "Chotu")
	fmt.Println("New Marks of Chotu : ", studentGrade["Chotu"])

	value, has := studentGrade["Chotu"]
	fmt.Printf("Isexist : %t\nValue : %d\n", has, value)

	Value, Has := studentGrade["Abhinav"]
	fmt.Printf("Isexist : %t\nValue : %d\n", Has, Value)

	for index, value := range studentGrade {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"persn1": 45,
		"persn2": 75,
		"persn3": 46,
		"persn4": 55,
	}
	for index, value := range person {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

}
