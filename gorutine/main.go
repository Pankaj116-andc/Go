package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(1000*time.Millisecond)
	fmt.Println("helooooo")
}
func sayHi() {
	fmt.Println("hiii")
}
func main() {
	go sayHello()
	go sayHi()

	//  Waitingg for func. to finish
	time.Sleep(2000 * time.Millisecond)

}
