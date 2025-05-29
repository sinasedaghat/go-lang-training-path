package main

import "fmt"

func main() {
	rangeMapExample()
	fmt.Println("Last line in main function.")
}

func basicForExample() {
	fmt.Println("YoYo! you call basicForExample function.")
	for i := 0; i < 5; i++ {
		fmt.Println("i => ", i)
	}
}

func counterLineExample() {
	fmt.Println("YoYo! you call counterLineExample function.")
	var i int
	fmt.Println("i before loop ===> ", i)
	for i = 0; i < 5; i++ {
		fmt.Println("i in loop ===> ", i)
	}
	fmt.Println("i after loop ===> ", i)
}

func likeWhileExample() {
	fmt.Println("YoYo! you call likeWhileExample function.")
	i := 0
	fmt.Println("i before loop ===> ", i)
	for i < 5 {
		fmt.Println("i in loop ===> ", i)
		i++
	}
	fmt.Println("i after loop ===> ", i)
}

func breakContinueExample() {
	fmt.Println("YoYo! you call breakContinueExample function.")
	fmt.Println("i < 10 is condition, in i == 5 has continue and in i == 8 has break.")
	for i := 0; i < 10; i++ {
		fmt.Println("i top of scope ===> ", i)
		if i == 5 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Println("i end of scope ===> ", i)
	}
	fmt.Println("End of function.")
}

func returnExample() {
	fmt.Println("YoYo! you call returnExample function.")
	fmt.Println("i < 5 is condition, in i == 2 has return.")
	for i := 0; i < 5; i++ {
		fmt.Println("i ===> ", i)
		if i == 2 {
			return
		}
	}
	fmt.Println("End of function.")
}

func whileTrueExample() {
	var counter int = 0

	for {
		fmt.Println("While true loop has condition for return in 100 step", counter)
		if counter == 100 {
			return
		}
		counter++
	}
	fmt.Println("last print")
}

func rangConstantExample() {
	var index int
	for index = range 7 {
		fmt.Println("use range on constant number", index)
	}

	fmt.Println("last print ==> ", index) // last index 6
}

func rangeArrayExample() {
	arr := []string{"first", "second", "third"}

	for index, value := range arr {
		fmt.Println("use range on array example (index, value) ==> ", index, value)
	}
}

func rangeMapExample() {
	mapData := map[string]string{
		"first":  "1st",
		"second": "2nd",
		"third":  "3rd",
		"forth":  "4th",
		"fifth":  "5th",
	}

	for key, value := range mapData {
		fmt.Println("use range on map example (key, value) ==> ", key, value)
	}
}
