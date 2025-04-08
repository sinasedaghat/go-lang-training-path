package main

import (
	"fmt"
)

var grid = [][]string{
	{"0", "1", "2"},
	{"3", "4", "5"},
	{"6", "7", "8"},
}
var winPattern = [][]int{
	{1, 2, 3},
	{1, 4, 7},
	{1, 5, 9},
	{2, 5, 8},
	{3, 6, 9},
	{3, 5, 7},
	{4, 5, 6},
	{7, 8, 9},
}
var u1 string
var u2 string
var last int // this variable store last selected cell

// var grid = [][]string{
// 	{"0️⃣", "1️⃣", "2️⃣"},
// 	{"3️⃣", "4️⃣", "5️⃣"},
// 	{"6️⃣", "7️⃣", "8️⃣"},
// }

func main() {
	winnerChecker(last)
	fmt.Println(`
		Welcome to 'Tic Tac Toe' Game
	`)

	fmt.Print("Enter first player name (🔵): ")
	fmt.Scan(&u1)

	fmt.Print("Enter second player name (🔴): ")
	fmt.Scan(&u2)

	fmt.Println("\nLET'S GO")
	for i := 0; i <= 9; i++ {
		showGrid(grid)
		var row, col int
		if i%2 == 0 {
			row, col = selectCell(u1)
			grid[row][col] = u1
		} else {
			row, col = selectCell(u2)
			grid[row][col] = u2
		}

		fmt.Printf("row: %v, col: %v\n", row, col)
	}
}

func showGrid(grid [][]string) {
	fmt.Println(grid)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if col != 1 {
				fmt.Printf("| %v |", translateCell(grid[row][col]))
			} else {
				fmt.Printf(" %v ", translateCell(grid[row][col]))
			}
		}
		fmt.Println("\n-------------")
	}
}

func selectCell(player string) (row, col int) {
	var message string
	var cell int
	for {
		if message == "" {
			fmt.Printf("%v Select cell: ", player)
		} else {
			fmt.Printf("%v Select cell: ", message)
		}
		fmt.Scan(&cell)
		row, col = computeCell(cell)
		if grid[row][col] == "u1" || grid[row][col] == "u2" {
			message = fmt.Sprintf("Cell %v is already selected, select another cell.", cell)
			continue
		}

		break
	}
	last = cell
	return
}

func computeCell(cell int) (row, col int) {
	return cell / 3, cell % 3
}

func translateCell(cell string) string {
	switch cell {
	case u1:
		return "🔵"
	case u2:
		return "🔴"
	default:
		return cell
	}
}

func winnerChecker(last int) {
	for i := 0; i < len(winPattern); i++ {
		fmt.Println(winPattern)
	}
	fmt.Println(last)
}
