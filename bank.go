package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var fileName = "number.txt"

func WriteToFile(number int) {
	n := fmt.Sprint(number)
	os.WriteFile(fileName, []byte(n), 0644)
}

func ReadFile() {
	number, _ := os.ReadFile(fileName)
	n := string(number)
	rt, _ := strconv.Atoi(n)
	fmt.Printf("random value is %d\n", rt)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Intn(100)
	WriteToFile(randomValue)
	fmt.Println("welcome to New World!")
	for {
		input := showMenu()
		switch input {
		case 1:
			ReadFile()
		case 2:
			var temp int
			fmt.Print("enter number:")
			fmt.Scan(&temp)
			randomValue += temp
			WriteToFile(randomValue)
			ReadFile()
		case 3:
			var temp int
			fmt.Print("enter number:")
			fmt.Scan(&temp)
			randomValue -= temp
			WriteToFile(randomValue)
			ReadFile()
		default:
			fmt.Println("GoodBye!")
			return
		}
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
