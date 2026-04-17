package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userid"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while getting response: ", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: ", res.StatusCode)
		return
	}
	defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error while reading data : ", err)
	// 	return
	// }
	// fmt.Println("Data : ", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error while decoding : ", err)
		return
	}
	fmt.Println("Todo : ", todo)
	fmt.Println("title : ", todo.Title)

}

func performPostRequest() {
	todo := Todo{
		UserId:    16,
		Title:     "Polity Revision",
		Completed: false,
	}

	//Convert the todo to json
	jasonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error While Marshing: ", err)
		return
	}

	//Convert jsonData to string
	jsonStr := string(jasonData)
	jsonReader := strings.NewReader(jsonStr)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myUrl, "aplicatin/json", jsonReader)
	if err != nil {
		fmt.Println("Error while  Posting: ", err)
		return
	}
	defer res.Body.Close()
	Data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading: ", err)
		return
	}
	fmt.Println("Data: ", string(Data))
	fmt.Println("Status: ", res.Status)
}

func performPutRequest() {
	todo := Todo{
		UserId:    116,
		Title:     "Bio Revision",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error while Marshling: ", err)
		return
	}
	jsonStr := string(jsonData)
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	jsonReader := strings.NewReader(jsonStr)

	//  Creating PUT request
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error while creating PUT request: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send Request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending a PUT request: ", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Status: ", res.Status)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("Data: ", string(data))
}

func performDeletRequest() {
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	// Creating Delete Request
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("Error While creating a delete request: ", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending the delete request: ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Status: ", res.Status)
}

func main() {
	fmt.Println("Learning CRUD...")
	// performGetRequest()
	// performPostRequest()
	// performPutRequest()
	performDeletRequest()

}
