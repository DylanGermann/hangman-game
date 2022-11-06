package main

import (
	"fmt"
	"os"

	"projets.perso/hangman-game/hangman"
)

func main() {
	g := hangman.New(8, "Golang")
	
	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
	}

	
}