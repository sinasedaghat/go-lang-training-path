package main

import (
	"fmt"
	"strconv"
)

var grid = [][]string{
	{"0", "1", "2"},
	{"3", "4", "5"},
	{"6", "7", "8"},
}

var winPattern = [][]int{
	{0, 1, 2},
	{0, 3, 6},
	{0, 4, 8},
	{1, 4, 7},
	{2, 5, 8},
	{2, 4, 6},
	{3, 4, 5},
	{6, 7, 8},
}

var u1 string // X
var u2 string // O

func main() {
	startUp()

	for round := 0; round < 9; round++ {
		showGrid()
		cell := getCellNumber(round)
		setValue(round, cell)

		if winnerChecker(cell) {
			fmt.Printf("🏆 YO! %v won. 🎉\n", playerPick(round, false))
			showGrid()
			return
		}
	}
}

func startUp() {
	fmt.Println("Welcome To TicTacToe")

	fmt.Print("Enter first player name (X): ")
	fmt.Scan(&u1)

	fmt.Print("Enter second player name (O): ")
	fmt.Scan(&u2)

	fmt.Println("\nLET'S GO")
}

func showGrid() {
	fmt.Println("\n ")
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if col != 1 {
				fmt.Printf("| %v |", grid[row][col])
			} else {
				fmt.Printf(" %v ", grid[row][col])
			}
		}
		if row != 2 {
			fmt.Println("\n-------------")
		} else {
			fmt.Println("\n\n ")
		}
	}
}

func getCellNumber(round int) (cell int) {
	var message string
	var input string
	player := playerPick(round, false)

	for {
		if message != "" {
			fmt.Println(message)
			fmt.Print("Select another cell: ")
		} else {
			fmt.Printf("%v select the desired cell: ", player)
		}

		fmt.Scan(&input)
		var err error
		cell, err = strconv.Atoi(input) // cell, err := strconv.Atoi(input)

		if err != nil {
			message = "The selected value is not valid, only numbers are valid."
			continue
		} else if cell > 9 || cell < 0 {
			message = "Your selection is not valid, valid cells are between 0 and 9."
			continue
		} else if value := cellValue(cell); value == "X" || value == "O" {
			message = "The desired cell is already selected."
			continue
		}
		break
	}

	return cell
}

func playerPick(round int, symbol bool) (player string) {
	if round%2 == 0 {
		if symbol {
			return "X"
		} else {
			return u1
		}
	} else {
		if symbol {
			return "O"
		} else {
			return u2
		}
	}
}

func cellValue(cell int) (value string) {
	// check valid cell:  0 <= cell <= 9

	return grid[cell/3][cell%3]
}

func setValue(round, cell int) {
	// check valid cell

	grid[cell/3][cell%3] = playerPick(round, true)
}

func winnerChecker(target int) bool {
	value := cellValue(target)

	for i := 0; i < len(winPattern); i++ {
		if pattern := winPattern[i]; includes(pattern, target) && cellValue(pattern[0]) == value && cellValue(pattern[1]) == value && cellValue(pattern[2]) == value {
			return true
		}
	}
	return false
}

func includes(array []int, target int) bool {
	for item := 0; item < len(array); item++ {
		if array[item] == target {
			return true
		}
	}
	return false
}
