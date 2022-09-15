package main

import (
	"fmt"
)

type game struct {
	board  [11]string
	player string
}

func main() {

	var turn game
	turn.player = "X"

	printBoard := turn.board

	for {
		var posInput int
		fmt.Println("Enter position of choice (1 - 9): ")
		fmt.Scanln(&posInput)

		if posInput < 1 || posInput > 9 {
			fmt.Println("Invalid input!")
			continue
		}

		if printBoard[posInput] == "" {
			printBoard[posInput] = turn.player
		} else {
			fmt.Println("The position is filled!")
			continue
		}

		fmt.Println(printBoard)

		for i := 0; i < 3; i++ {
			if
			// Horizontal Check
			printBoard[i] == printBoard[i+1] && printBoard[i+1] == printBoard[i+2] && printBoard[i] != "" ||
				printBoard[i+3] == printBoard[i+4] && printBoard[i+4] == printBoard[i+5] && printBoard[i+3] != "" ||
				printBoard[i+6] == printBoard[i+7] && printBoard[i+7] == printBoard[i+8] && printBoard[i+6] != "" {
				fmt.Println(turn.player, "wins!")
				return
			}
			if
			// Vertical Check
			printBoard[i] == printBoard[i+3] && printBoard[i+3] == printBoard[i+6] && printBoard[i] != "" ||
				printBoard[i+1] == printBoard[i+4] && printBoard[i+4] == printBoard[i+7] && printBoard[i+1] != "" ||
				printBoard[i+2] == printBoard[i+5] && printBoard[i+5] == printBoard[i+8] && printBoard[i+2] != "" {
				fmt.Println(turn.player, "wins!")
				return
			}
			if
			// Diagonal Check
			printBoard[i] == printBoard[i+4] && printBoard[i+4] == printBoard[i+8] && printBoard[i] != "" ||
				printBoard[i+2] == printBoard[i+4] && printBoard[i+4] == printBoard[i+6] && printBoard[i+2] != "" {
				fmt.Println(turn.player, "wins!")
				return
			}

		}

		if turn.player == "X" {
			turn.player = "O"
		} else {
			turn.player = "X"
		}
	}
}
