package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	fmt.Print("Enter your name:")
	// fmt.Scan(&name)
	// fmt.Printf("Hellow Mr.%s\n", name)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello Mr.%s\n", name)
}
