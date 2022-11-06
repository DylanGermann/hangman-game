package main

import (
	"fmt"
	"os"

	"projets.perso/hangman-game/dictionnary"
	"projets.perso/hangman-game/hangman"
)

func main() {
	hangman.DrawWelcome()
	for {
		playGame()
		choice, err := hangman.AskForPlay()
		if err != nil {
			fmt.Printf("We got an error : %v", err)
		}
		if !choice {
			os.Exit(1)
		}
	}
	
}

func playGame() {
	err := dictionnary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionnary: %v\n", err)
		os.Exit(1)
	}

	g, err := hangman.New(8, dictionnary.PickWord())
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}

		

	guess := ""
	for {
		hangman.Draw(g, guess)

		if g.State == "won" || g.State == "lost" {
			break
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}