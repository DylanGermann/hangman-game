package main

import (
	"fmt"

	"projets.perso/hangman-game/hangman"
)

func main() {
	g := hangman.New(8, "Golang")
	fmt.Println(g)
}