package main

import "fmt"

var grid = [][]string{
	{"0", "1", "2"},
	{"3", "4", "5"},
	{"6", "7", "8"},
}

// var grid = [][]string{
// 	{"0️⃣", "1️⃣", "2️⃣"},
// 	{"3️⃣", "4️⃣", "5️⃣"},
// 	{"6️⃣", "7️⃣", "8️⃣"},
// }

func main() {
	var u1 string
	var u2 string
	// fmt.Println(grid)
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
			row, col = computeCell(selectCell(u1))
			grid[row][col] = "🔵"
		} else {
			row, col = computeCell(selectCell(u2))
			grid[row][col] = "🔴"
		}

		fmt.Printf("row: %v, col: %v\n", row, col)
	}
}

func showGrid(grid [][]string) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if col != 1 {
				fmt.Printf("| %v |", grid[row][col])
			} else {
				fmt.Printf(" %v ", grid[row][col])
			}
		}
		fmt.Println("\n-------------")
	}
}

func selectCell(player string) (cell int) {
	fmt.Printf("%v Select cell: ", player)
	fmt.Scan(&cell)
	return
}

func computeCell(cell int) (row, col int) {
	return cell / 3, cell % 3
}
