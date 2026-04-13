package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("Error while creating file :", err)
			return
		}
		defer file.Close()

		content := "hellow world by prince"
		_, error := io.WriteString(file, content+"\n")//return bytes written and error
		if error != nil {
			fmt.Println("Error while writing file", error)
		}

		fmt.Println("Succesfully created file")
	*/
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening the file; ", err)
		return
	}
	defer file.Close()

	// Create a buffer to read the file content
	buffer := make([]byte, 1024)

	//Read the file content into the buffer
	for {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("Error while reading a file: ", err)
			return
		}

		fmt.Println(string(buffer[:n]))

		if err == io.EOF {
			break
		}
	}

	/*
		//Read the entire file into a byte slice
		content, err := ioutil.ReadFile("example.txt")
		if err != nil {
			fmt.Println("Error while reading file: ", err)
			return
		}
		fmt.Println(string(content))
	*/

}
