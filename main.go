package main

import "fmt"

const (
	Empty = "⬜"
	X     = "❌"
	O     = "🟢"
)

func main() {
	board := initialsGame()
	printBoard(board)
	for {
		randomMoveO(&board)
		printBoard(board)
		status := checkStatus(&board)
		switch status {
		case X:
			fmt.Printf("Player %s wins!\n", X)
			return
		case O:
			fmt.Printf("Player %s wins!\n", O)
			return
		case "continu":
			randomMoveX(&board)
			checkStatus(&board)
			continue
		}
	}
}
