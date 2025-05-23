package main

import (
	"array-practice/products"
	"fmt"
)

var hobbies = [3]string{}
var thirdTaskResult = []string{}
var goals = []string{}

func main() {
	fmt.Println("Hello Array")
	firstTask()
	secondTask()
	thirdTask()
	fourthTask()
	fifthTask()
	sixthTask()
	WorkProducts()

	fmt.Printf("\n\n👇 Solutions from Maximilian 👇\n")
	solutions()
}

func firstTask() {
	fmt.Printf("\n👇 First Task 👇\n")
	hobbies = [3]string{"Music", "Watch football", "Video game"}
	fmt.Println("Hobbies: ", hobbies)
}

func secondTask() {
	fmt.Printf("\n👇 Second Task 👇\n")
	fmt.Println("First hobby: ", hobbies[0])
	newArray := hobbies[1:3]
	fmt.Println("New array from second and third hobbies: ", newArray)
}

func thirdTask() {
	fmt.Printf("\n👇 Third Task 👇\n")
	firstWay := hobbies[:2]
	secondWay := hobbies[0:2]
	fmt.Println("Use [:2]: ", firstWay)
	fmt.Println("Use [0:2]: ", secondWay)
	thirdTaskResult = secondWay
}

func fourthTask() {
	fmt.Printf("\n👇 Fourth Task 👇\n")
	fmt.Println("thirdTaskResult: ", thirdTaskResult, cap(thirdTaskResult))
	fourthTaskResult := thirdTaskResult
	fourthTaskResult = append(fourthTaskResult, hobbies[2])
	fourthTaskResult[0] = fourthTaskResult[1]
	fourthTaskResult = append(fourthTaskResult[0:1], fourthTaskResult[2:]...)
	fmt.Println("New Array form thirdTaskResult contain second and third items from hobbies array", fourthTaskResult, cap(fourthTaskResult))

	thirdTaskResult = thirdTaskResult[1:3]
	fmt.Println("Us e reassign slice and capacity feature in slice: ", thirdTaskResult)
}

func fifthTask() {
	fmt.Printf("\n👇 Fifth Task 👇\n")
	// goals = append(goals, "Expert developer", "Independent home")
	goals = []string{"Expert in javascript", "Go developer"}
	fmt.Println("Goals: ", goals)
}

func sixthTask() {
	fmt.Printf("\n👇 Sixth Task 👇\n")
	goals[1] = "Good Go Lang developer"
	goals = append(goals, "Work in heigh level company")
	fmt.Println("Update Goals: ", goals)
}

func WorkProducts() {
	fmt.Printf("\n👇 Seventh Task 👇\n")
	products.ListOfProducts()
	products.AddProduct(products.NewProduct("Headphone", 250.99))
	products.ListOfProducts()
}

type Product struct {
	id    string
	title string
	price float64
}

func solutions() {
	// 1)
	Hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(Hobbies)

	// 2)
	fmt.Println(Hobbies[0])
	fmt.Println(Hobbies[1:3])
	fmt.Println(Hobbies[1:])

	// 3)
	MainHobbies := Hobbies[:2] // Hobbies[0;2]
	fmt.Println(MainHobbies)

	// 4)
	MainHobbies = MainHobbies[1:3]
	fmt.Println(MainHobbies)

	// 5)
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// 6)
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)

	// 7)
	Products := []Product{
		{
			"first",
			"first title",
			12.99,
		},
		{
			"second",
			"second title",
			15.99,
		},
	}
	fmt.Println(Products)

	newProduct := Product{
		"third",
		"third title",
		16.99,
	}
	Products = append(Products, newProduct)
	fmt.Println(Products)
}
