package main

import (
	"errors"
	"fmt"
)

type game struct {
	board  [9]string
	player string
	turn   int
}

func gameInit() int {
	var posInput int
	fmt.Println("Enter position of choice (1 - 9): ")
	fmt.Scan(&posInput)
	return posInput
}

func printBoard(b [9]string) {

	for i, v := range b {
		if v == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(v)
		}

		if i > 0 && (i+1)%9 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("")
		}
	}
}

func (firstTurn *game) playerSwitch() {
	if firstTurn.player == "X" {
		firstTurn.player = "O"
		return
	}
	firstTurn.player = "X"
}

func (firstTurn *game) play(pos int) error {
	if firstTurn.board[pos-1] == "" {
		firstTurn.board[pos-1] = firstTurn.player
		firstTurn.playerSwitch()
		firstTurn.turn += 1
		return nil
	}
	return errors.New("Position is filled.")
}

func winnerCheck(b [9]string, n int) (bool, string) {

	test := false
	i := 0

	//Horizontal Test
	for i < 9 {
		test = b[i] == b[i+1] && b[i+1] == b[i+2] && b[i] != ""
		if !test {
			i += 3
		} else {
			return true, b[i]
		}
	}
	i = 0
	//Vertical Test
	for i < 3 {
		test = b[i] == b[i+3] && b[i+3] == b[i+6] && b[i] != ""
		if !test {
			i += 1
		} else {
			return true, b[i]
		}
	}
	//Diagonal 1 Test
	if b[0] == b[4] && b[4] == b[8] && b[0] != "" {
		return true, b[i]
	}
	//Diagonal 2 test
	if b[2] == b[4] && b[4] == b[6] && b[2] != "" {
		return true, b[i]
	}
	if n == 9 {
		return true, ""
	}
	return false, ""
}

func main() {

	var firstTurn game
	firstTurn.player = "X"

	var gameWinner string
	gameOver := false
	for gameOver != true {
		printBoard(firstTurn.board)
		move := gameInit()
		err := firstTurn.play(move)
		if err != nil {
			fmt.Println(err)
			continue
		}

		gameOver, gameWinner = winnerCheck(firstTurn.board, firstTurn.turn)
	}

	printBoard(firstTurn.board)

	if gameWinner == "" {
		fmt.Println("It's a draw!")
	} else {
		fmt.Printf("%s wins! \n", gameWinner)
	}
}
