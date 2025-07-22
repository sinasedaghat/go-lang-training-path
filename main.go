package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	go integerPrinter(1, channel)
	fmt.Println(<-channel)
}

func integerPrinter(number int, channel chan int) <-chan int {
	channel2 := make(chan int)
	fmt.Println("I give", number)
	time.Sleep(time.Duration(number) * time.Second)
	channel <- number * 3
	return channel2
}

// func sayHello(id int) {
// 	fmt.Printf("Hello from goroutine %d\n", id)
// }

// func greet(str string, doneChan chan bool) {
// 	fmt.Println(str, " ===> HI, from greet")
// 	doneChan <- true
// }

// // func greet(str string) <-chan bool {
// // 	channel := make(chan bool)
// // 	fmt.Println(str, " ===> HI, from greet")
// // 	channel <- true
// // 	return channel
// // }

// func slowGreet(str string, doneChan chan bool) {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println(str, " ===> HI, from slow greet")
// 	doneChan <- true
// }
