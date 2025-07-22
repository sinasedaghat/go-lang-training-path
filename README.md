this is GoLang Training Path\
In Maximilian's course, parts 149 to 158\
https://www.udemy.com/course/go-the-complete-guide/\
1. 
```
func main() {
  channel := make(chan bool)
  go greet("First", channel)
  valChan := <-channel
  fmt.Println(valChan)
}

func greet(str string, doneChan chan bool) {
	fmt.Println(str, " ===> HI, from greet")
	doneChan <- true
}
```

2. 
```
func main() {
	channel := make(chan bool)
	go greet("First", channel)
	go greet("Second", channel)
	go slowGreet("First Slow", channel)
	go greet("Third", channel)

	fmt.Println(<-channel)
	<-channel
	<-channel
	<-channel
	<-channel
	// time.Sleep(3 * time.Second) // bad practice
}

func greet(str string, doneChan chan bool) {
	fmt.Println(str, " ===> HI, from greet")
	doneChan <- true
}

func slowGreet(str string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(str, " ===> HI, from slow greet")
	doneChan <- true
}
```
3. 
```
func main() {
	for i := range 5 {
		go sayHello(i)
	}
	time.Sleep(1 * time.Second)
}

func sayHello(id int) {
	fmt.Printf("Hello from goroutine %d\n", id)
}
```
4. 
```
func main() {
	channel := make(chan int)

	for index := range 5 {
		go integerPrinter(index, channel)
	}

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	integerPrinter(4, channel) // create deadlock
	fmt.Println(<-channel) 
}

func integerPrinter(number int, channel chan int) {
	fmt.Println("I give", number)
	time.Sleep(time.Duration(number) * time.Second)
	channel <- number * 3
}
```