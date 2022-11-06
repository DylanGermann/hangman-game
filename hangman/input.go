package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("Whats is your letter ? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Printf("Invalid letter input. letter=%v, len=%v \n", guess, len(guess))
			continue
		}
		valid = true
	}
	return guess, nil
}

func AskForPlay() (response bool, err error) {
	valid := false
	for !valid {
		fmt.Println("Do you want to play one more game ? Y for yes or N for no")
		choice, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}
		
		choice = strings.TrimSpace(choice)
		choice = strings.ToUpper(choice)
		
		if choice == "Y" {
			response = true
			valid = true
		} else if choice == "N" {
			response = false
			valid = true
		}
	}
	return response, nil
}