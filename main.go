package main

import "fmt"

func main() {
	g := hangman.New(8, "Golang")
	fmt.Println(g)
}