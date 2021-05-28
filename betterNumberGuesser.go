// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program has the user guess a number between 0 and 9

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This function has the user guess a number between 0 and 9
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 9
	var randomNumber = rand.Intn(max-min) + min
	var userGuess int

	// input
	fmt.Println("This program has the user guess a number between 0 and 9")
	fmt.Println()
	fmt.Print("Enter your guess: ")
	fmt.Scanln(&userGuess)
	fmt.Println()

	// process
	for userGuess != randomNumber {
		if userGuess > randomNumber {
			// output
			fmt.Println(userGuess, "is too high")
			fmt.Print("Guess again: ")
			fmt.Scanln(&userGuess)
		} else if userGuess < randomNumber {
			// output
			fmt.Println(userGuess, "is too low")
			fmt.Print("Guess again: ")
			fmt.Scanln(&userGuess)
		} else {
			// output
			fmt.Println(userGuess, "is not a number between 1 - 9")
			fmt.Print("Guess again: ")
			fmt.Scanln(&userGuess)
		}
	}

	// output
	fmt.Println("That is correct")

}
