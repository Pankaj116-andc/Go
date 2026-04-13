package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("web Service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/")
	if err != nil {
		fmt.Println("Error while fetching the data ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response %T\n",res)
	// fmt.Println("response",res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response", err)
		return
	}

	// fmt.Println("Response ", data) //print data in byte
	fmt.Println("Response ", string(data))
}
