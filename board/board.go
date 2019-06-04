package board

import (
	"fmt"
	"strconv"
)

func Play() {
	player, next := 1, 2
	a := getBoard()
	printBoard(a)

	for true {
		var i int
		fmt.Print("Enter column: ")
		_, err := fmt.Scanf("%d\n", &i)
		if err != nil || i > 7 || i < 1 {
			fmt.Println("Input is invalid")
			continue
		}

		valid := insertItem(player, a, i-1)
		if !valid {
			continue
		}

		printBoard(a)
		if isWinner(a, player) {
			fmt.Println("Congrats player", player, " you win!")
			break
		}

		player, next = next, player
	}
}

func getBoard() [][]int {
	board := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}

	return board
}

func insertItem(p int, board [][]int, column int) bool {
	if board[0][column] > 0 {
		fmt.Println("That column is full!")
		return false
	}
	for i := 1; i < 6; i++ {
		if board[i][column] > 0 {
			board[i-1][column] = p
			return true
		}
	}
	// must be first input in column
	board[5][column] = p

	return true
}

func printBoard(board [][]int) {
	for x := range board {
		for y := range board[x] {
			fmt.Print(board[x][y], " ")
		}
		fmt.Println("")
	}
}

// https://github.com/qu1j0t3/fhourstones/blob/master/Connect4.java
// algorithm to check for win using a bitboard - need to figure out what is wrong with my impl
func isWinner(gameboard [][]int, player int) bool {
	bitboard := ""
	for y := 0; y < 7; y++ {
		for x := 5; x >= 0; x-- {
			if gameboard[x][y] == player {
				bitboard = bitboard + "1"
			} else {
				bitboard = bitboard + "0"
			}
		}
		bitboard = bitboard + "0"
	}
	// Messy way to fill in bits
	bitboard = bitboard + "000000000000000"

	board, err := strconv.ParseUint(bitboard, 2, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Something's gone terribly wrong!")
	}

	// // check \ diagonal
	// y := board & (board >> 6)
	// if (y & (y >> 2 * 6)) != 0 {
	// 	return true
	// }

	// // check horizontal
	// y = board & (board >> 7)
	// if (y & (y >> 2 * 7)) != 0 {
	// 	return true
	// }

	// // check / diagonal
	// y = board & (board >> 8)
	// if (y & (y >> 2 * 8)) != 0 {
	// 	return true
	// }

	// // check vertical
	// y = board & (board >> 1)
	// if (y & (y >> 2)) != 0 {
	// 	return true
	// }

	// return false

	return isWin(board)
}

func isWin(bitboard uint64) bool {
	if bitboard&(bitboard>>6)&(bitboard>>12)&(bitboard>>18) != 0 {
		return true // diagonal \
	}
	if bitboard&(bitboard>>8)&(bitboard>>16)&(bitboard>>24) != 0 {
		return true // diagonal /
	}

	if bitboard&(bitboard>>7)&(bitboard>>14)&(bitboard>>21) != 0 {
		return true // horizontal
	}
	if bitboard&(bitboard>>1)&(bitboard>>2)&(bitboard>>3) != 0 {
		return true // vertical
	}
	return false
}
