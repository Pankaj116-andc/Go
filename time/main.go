package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now()
	fmt.Println("Curent Time : ", current_time)
	fmt.Printf("Type of current Time : %T\n ", current_time)

	formatted := current_time.Format("2006/01/02, Monday")
	fmt.Println("Formated Time : ", formatted)

	layout_str := "02-01-2006"
	date_Str := "01-04-2026"
	formatted_time, _ := time.Parse(layout_str, date_Str)
	fmt.Println("Formatted Time : ", formatted_time)

	// add 1 day in currentTime
	new_date := current_time.Add(24 * time.Hour)
	fmt.Println("new date time : ", new_date)
	formatted_new_date := new_date.Format("2006/01/02,Monday")
	fmt.Println("new date time : ", formatted_new_date)
}
