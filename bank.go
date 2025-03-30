package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Intn(100)
	fmt.Println("welcome to New World!")
	input := showMenu()
	if input == 1 {
		fmt.Printf("random value is %d\n", randomValue)
	} else if input == 2 {
		var temp int
		fmt.Print("enter number:")
		fmt.Scan(&temp)
		randomValue += temp
		fmt.Printf("random value is %d\n", randomValue)
	} else if input == 3 {
		var temp int
		fmt.Print("enter number:")
		fmt.Scan(&temp)
		randomValue -= temp
		fmt.Printf("random value is %d\n", randomValue)
	} else {
		fmt.Println("GoodBye!")
	}
}

func showMenu() int {
	var input int
	fmt.Println(`Select next Action:
1.show random value:
2.add number to random value
3.subtract nubmer to random value
4.exit`)
	fmt.Scan(&input)
	return input
}
