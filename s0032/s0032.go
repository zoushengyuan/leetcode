package main

import (
	"fmt"
)

func isValidRow(board [][]byte, row int) bool {
	m := make(map[byte]int)
	for col := 0; col < 9; col++ {
		if m[board[row][col]] == '.' {
			continue
		}
		if m[board[row][col]] == 1 {
			return false
		}
		m[board[row][col]]++
	}
	return true
}

func isValidCol(board [][]byte, col int) bool {
	m := make(map[byte]int)
	for row := 0; row < 9; row++ {
		if m[board[row][col]] == '.' {
			continue
		}
		if m[board[row][col]] == 1 {
			return false
		}
		m[board[row][col]]++
	}
	return true
}

func isValidBox(board [][]byte, row int, col int) bool {
	m := make(map[byte]int)
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if m[board[row][col]] == '.' {
				continue
			}
			if m[board[row][col]] == 1 {
				return false
			}
			m[board[row][col]]++
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		if !isValidRow(board, row) {
			return false
		}
	}
	for col := 0; col < 9; col++ {
		if !isValidCol(board, col) {
			return false
		}
	}
	for row := 0; row < 9; row += 3 {
		for col := 0; col < 9; col += 3 {
			if !isValidBox(board, row, col) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println()
}
