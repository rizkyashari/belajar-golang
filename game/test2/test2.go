package main

import (
	"fmt"
)

func main() {

	for {
		var input string
		fmt.Println("Input either 'X', 'O', or '-' (max: 9 characters): ")
		fmt.Scanln(&input)

		for i := 0; i < 3; i++ {
			if
			// Invalid Grid
			(input[i] == input[i+1] && input[i+1] == input[i+2]) &&
				(input[i+3] == input[i+4] && input[i+4] == input[i+5]) ||
				(input[i] == input[i+1] && input[i+1] == input[i+2]) &&
					(input[i+6] == input[i+7] && input[i+7] == input[i+8]) ||
				(input[i+3] == input[i+4] && input[i+4] == input[i+5]) &&
					(input[i+6] == input[i+7] && input[i+7] == input[i+8]) {
				fmt.Println("Invalid game board")
				return
			}
			if
			// Horizontal Check
			input[i] == input[i+1] && input[i+1] == input[i+2] ||
				input[i+3] == input[i+4] && input[i+4] == input[i+5] ||
				input[i+6] == input[i+7] && input[i+7] == input[i+8] {
				fmt.Println(string(input[i]), "wins!")
				return
			}
			if
			// Vertical Check
			input[i] == input[i+3] && input[i+3] == input[i+6] ||
				input[i+1] == input[i+4] && input[i+4] == input[i+7] ||
				input[i+2] == input[i+5] && input[i+5] == input[i+8] {
				fmt.Println(string(input[i]), "wins!")
				return
			}
			if
			// Diagonal Check
			input[i] == input[i+4] && input[i+4] == input[i+8] ||
				input[i+2] == input[i+4] && input[i+4] == input[i+6] {
				fmt.Println(string(input[i]), "wins!")
				return
			}
			if
			// Draw
			input[i] != input[i+1] && input[i+1] != input[i+2] ||
				input[i+3] != input[i+4] && input[i+4] != input[i+5] ||
				input[i+6] != input[i+7] && input[i+7] != input[i+8] ||
				input[i] != input[i+3] && input[i+3] != input[i+6] ||
				input[i+1] != input[i+4] && input[i+4] != input[i+7] ||
				input[i+2] != input[i+5] && input[i+5] != input[i+8] ||
				input[i] != input[i+4] && input[i+4] != input[i+8] ||
				input[i+2] != input[i+4] && input[i+4] != input[i+6] {
				fmt.Println("It's a draw!")
				return
			}
		}
	}
}
