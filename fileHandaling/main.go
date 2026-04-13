package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file :", err)
		return
	}
	defer file.Close()

	content := "hellow world by prince"
	_, error := io.WriteString(file, content+"\n")
	if error != nil {
		fmt.Println("Error while writing file", error)
	}

	fmt.Println("Succesfully created file")
}
