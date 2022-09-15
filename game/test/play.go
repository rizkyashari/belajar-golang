package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	for {
		var input string
		fmt.Println("Input either 'X', 'O', or '-' (must be 9 characters): ")
		fmt.Scanln(&input)

		for i := 0; i < 3; i++ {

			if strings.Contains(input, "X") == true ||
				strings.Contains(input, "O") == true ||
				strings.Contains(input, "-") == true {
			} else {
				fmt.Println("Invalid input")
				return
			}

			if utf8.RuneCountInString(input) > 9 {
				fmt.Println("Input exceeds the limit")
				return
			}
			if utf8.RuneCountInString(input) < 9 {
				fmt.Println("Input less than 9 characters")
				return
			}
			if
			// Invalid Grid

			// Invalid Rows
			(((input[i] == input[i+1] && input[i+1] == input[i+2]) &&
				(input[i+3] == input[i+4] && input[i+4] == input[i+5]) ||
				(input[i] == input[i+1] && input[i+1] == input[i+2]) &&
					(input[i+6] == input[i+7] && input[i+7] == input[i+8]) ||
				(input[i+3] == input[i+4] && input[i+4] == input[i+5]) &&
					(input[i+6] == input[i+7] && input[i+7] == input[i+8])) ||

				// Invalid Columns
				((input[i] == input[i+3] && input[i+3] == input[i+6]) &&
					(input[i+1] == input[i+4] && input[i+4] == input[i+7]) ||
					(input[i] == input[i+3] && input[i+3] == input[i+6]) &&
						(input[i+2] == input[i+5] && input[i+5] == input[i+8]) ||
					(input[i+1] == input[i+4] && input[i+4] == input[i+7]) &&
						(input[i+2] == input[i+5] && input[i+5] == input[i+8]))) &&

				strings.Contains(input, "-") == false {
				fmt.Println("Invalid game board")
				return
			}
			if
			// Game in Progress
			(((input[i] != input[i+1] && input[i+1] != input[i+2]) ||
				(input[i+3] != input[i+4] && input[i+4] != input[i+5]) ||
				(input[i+6] != input[i+7] && input[i+7] != input[i+8])) ||
				((input[i] != input[i+3] && input[i+3] != input[i+6]) ||
					(input[i+1] != input[i+4] && input[i+4] != input[i+7]) ||
					(input[i+2] != input[i+5] && input[i+5] != input[i+8])) ||
				((input[i] != input[i+4] && input[i+4] != input[i+8]) ||
					(input[i+2] != input[i+4] && input[i+4] != input[i+6]))) &&
				strings.Contains(input, "-") == true {
				fmt.Println("Game still in progress...")
				return
			}
			if
			// Horizontal Check Row 1
			(input[i] == input[i+1] && input[i+1] == input[i+2]) ||
				strings.Contains(input, "-") == true {
				fmt.Println(string(input[i]), "wins!")
				return
			}
			if
			// Horizontal Check Row 2
			input[i+3] == input[i+4] && input[i+4] == input[i+5] ||
				strings.Contains(input, "-") == true {
				fmt.Println(string(input[i+3]), "wins!")
				return
			}
			if
			// Horizontal Check Row 3
			input[i+6] == input[i+7] && input[i+7] == input[i+8] ||
				strings.Contains(input, "-") == true {
				fmt.Println(string(input[i+6]), "wins!")
				return
			}
			if
			// Vertical Check Column 1
			(input[i] == input[i+3] && input[i+3] == input[i+6]) ||
				strings.Contains(input, "-") == true {
				fmt.Println(string(input[i]), "wins!")
				return
			}
			if
			// Vertical Check Column 2
			(input[i+1] == input[i+4] && input[i+4] == input[i+7]) ||
				strings.Contains(input, "-") == true {
				fmt.Println(string(input[i+1]), "wins!")
				return
			}
			if
			// Vertical Check Column 3
			(input[i+2] == input[i+5] && input[i+5] == input[i+8]) ||
				strings.Contains(input, "-") == true {
				fmt.Println(string(input[i+2]), "wins!")
				return
			}
			if
			// Diagonal Check 1
			(input[i] == input[i+4] && input[i+4] == input[i+8]) ||
				strings.Contains(input, "-") == true {
				fmt.Println(string(input[i]), "wins!")
				return
			}
			if
			// Diagonal Check 2
			(input[i+2] == input[i+4] && input[i+4] == input[i+6]) ||
				strings.Contains(input, "-") == true {
				fmt.Println(string(input[i+2]), "wins!")
				return
			}
			if
			// Draw
			(((input[i] != input[i+1] && input[i+1] != input[i+2]) ||
				(input[i+3] != input[i+4] && input[i+4] != input[i+5]) ||
				(input[i+6] != input[i+7] && input[i+7] != input[i+8])) ||
				((input[i] != input[i+3] && input[i+3] != input[i+6]) ||
					(input[i+1] != input[i+4] && input[i+4] != input[i+7]) ||
					(input[i+2] != input[i+5] && input[i+5] != input[i+8])) ||
				((input[i] != input[i+4] && input[i+4] != input[i+8]) ||
					(input[i+2] != input[i+4] && input[i+4] != input[i+6]))) &&
				strings.Contains(input, "-") == false {
				fmt.Println("It's a draw!")
				return
			}
		}
	}
}
