package main

import (
	"fmt"
	"math/rand"
	"time"
)

func initialsGame() [3][3]string {
	var board [3][3]string
	for i := range board {
		for j := range board[i] {
			board[i][j] = Empty
		}
	}
	return board
}

func printBoard(board [3][3]string) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func randomMoveX(board *[3][3]string) {
	rand.Seed(time.Now().UnixNano())

	row := rand.Intn(3)
	col := rand.Intn(3)

	for board[row][col] != Empty {
		row = rand.Intn(3)
		col = rand.Intn(3)
	}
	board[row][col] = X
}

func randomMoveO(board *[3][3]string) {
	rand.Seed(time.Now().UnixNano())

	row := rand.Intn(3)
	col := rand.Intn(3)

	for board[row][col] != Empty {
		row = rand.Intn(3)
		col = rand.Intn(3)
	}
	board[row][col] = O
}
func checkStatus(board *[3][3]string) string {
	for i := 0; i < 3; i++ {
		if board[i][0] != Empty && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
		if board[0][i] != Empty && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}

	if board[0][0] != Empty && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != Empty && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				return "continu"
			}
		}
	}
	printBoard(*board)
	*board = initialsGame()
	return "continu"
}
