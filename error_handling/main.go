package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator can't be zero")
	}
	return a / b, nil
}
func main() {
	fmt.Println("Error handling in function")
	ans, _ := divide(5, 0)
	fmt.Println("ans:", ans)

}
