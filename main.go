package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var numbOfGuesses int
var randomNum int

func start() {
	randomNum = rand.Intn(100)
}

func main() {
	start()
	readyToStart()
}

func readyToStart() {

	var gameStart string
	fmt.Println("Are you ready to guess the number?")
	fmt.Scanln(&gameStart)
	if gameStart == "yes" {
		guessTime()
	} else {
		fmt.Println("Still waiting on the yes.")
		readyToStart()
	}

}

func guessTime() {
	fmt.Println("Guess the number")
	var userGuessInput string
	fmt.Scanln(&userGuessInput)
	userGuess, error := strconv.Atoi(userGuessInput)
	if error != nil {
		fmt.Println("Invalid Input, try again")
		guessTime()
	}

	if randomNum == userGuess {
		numbOfGuesses++
		fmt.Println("You guessed correctly! It took you", numbOfGuesses, "guesses")
	} else if userGuess < randomNum {
		fmt.Println("Too low")
		numberOfGuesses()
		guessTime()
	} else if userGuess > randomNum {
		fmt.Println("Too high")
		numberOfGuesses()
		guessTime()
	}
}

func numberOfGuesses() {

	numbOfGuesses++
	fmt.Println("Guess Count: ", numbOfGuesses)

}
