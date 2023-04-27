// Created by: Charlotte Jhu
// Created on: April 2023
//
// This program is a simple guess the number game

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This function is the main function

	var guess int
	var answer int

	// input
	fmt.Println("Welcome to Guess the Number!")
	fmt.Println()
	fmt.Print("Enter a number between 1 and 6: ")
	fmt.Scanln(&guess)

	// process
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	answer = rand.Intn(max-min) + min

	// output
	if guess == answer {
		fmt.Println()
		fmt.Println("You guessed correctly!")
	} else {
		fmt.Println()
		fmt.Println("You guessed wrong!")
		fmt.Println("The correct answer was", answer)
	}
	fmt.Println("\nDone.")
}
