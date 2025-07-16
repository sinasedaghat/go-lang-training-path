package main

import (
	"fmt"
	"time"
)

func main() {
	// ch1 := make(chan string)
	channel := make(chan bool)
	go greet("First", channel)
	go greet("Second", channel)
	go slowGreet("First Slow", channel)
	go greet("Third", channel)

	// valChan := <-channel
	// fmt.Println(valChan)

	fmt.Println(<-channel)
	// <-channel
	// <-channel
	// <-channel
	// <-channel
	time.Sleep(2 * time.Second) // bad practice
}

func greet(str string, doneChan chan bool) {
	fmt.Println(str, " ===> HI, from greet")
	doneChan <- true
}

// func greet(str string) <-chan bool {
// 	channel := make(chan bool)
// 	fmt.Println(str, " ===> HI, from greet")
// 	channel <- true
// 	return channel
// }

func slowGreet(str string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(str, " ===> HI, from slow greet")
	doneChan <- true
}
