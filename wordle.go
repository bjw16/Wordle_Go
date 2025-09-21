package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var wordList = []string{
	"FETCH",
	"PARRY",
	"LOINS",
	"LUNGS",
	"TWIST",
}

var possibleLetters []string


// var streak int = 0
const Turns int = 6
const letterCount int = 5
var gameResults [Turns][letterCount]string

func setArray() {
	for i := range gameResults {
		for j := range gameResults[i]{
			gameResults[i][j] = "_"
		}
	}
	possibleLetters = []string{
		"A", "B", "C",
		"D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "X", "T", "U",
		"V", "W", "X", "Y", "Z",
	}
}

func main() {
	var GameMode int
	var playAgain bool = false

	for true {
		var results bool
		var playAgainResponse int
		if playAgain != true{
			printIntro()
			fmt.Scan(&GameMode)
		}
		if GameMode == 1 || playAgain == true{
			//play game
			setArray() 
			results = playGame()

			//play again
			printWinLose(results)
			fmt.Print("> ")
			fmt.Scan(&playAgainResponse)
			if playAgainResponse == 1 {
				//play again
				playAgain = true
				continue
			} else if playAgainResponse == 0 {
				playAgain = false
				continue
			}
		} else if GameMode == 2 {
			//show stats
			fmt.Println(gameResults[0])
		} else if GameMode == 0 {
			//error or exit
			break
		}
		fmt.Println("")
	}
}

func printIntro() {
	//game_loop
	fmt.Println("")
	fmt.Println("WORDLE_GO")
	fmt.Println("")
	fmt.Println("Select #:")
	fmt.Println("1 - Play Wordle")
	fmt.Println("2 - Streaks")
	fmt.Println("0 - Exit")
	fmt.Print("> ")
}

//function to play full game - module
func playGame() bool{
	//randomizes word
	wordAnswer := wordList[rand.Intn(len(wordList))]
	var winOrLose bool
	var guess string = ""


	//print turn
	for x := 0; x < Turns; x++ {
		//var playAgain int
		printTurn(x + 1, guess, wordAnswer)
		//resets guess
		guess = ""

		//player guess response
		for {
			fmt.Print("> ")
			fmt.Scan(&guess)
			if len(guess) < letterCount || len(guess) > letterCount{
				if len(guess) < letterCount {
					fmt.Println("Guess to small, try again!")
				} else {
					fmt.Println("Guess to big, try again!")
				}
				printTurn(x, guess, wordAnswer)
				guess = ""
			} else {
				break

			} 
		}
		fmt.Println("")
		//need error check if less than 5 letters

		//stores guess into results
		for y := 0; y < letterCount; y++{
			gameResults[x][y] = string(guess[y])
		}

		//Determines if guess matches selected word
		if guess == wordAnswer{
			winOrLose = true
			printTurn(x + 1, guess, wordAnswer)
			break
		} else if x == Turns - 1{
			winOrLose = false
			printTurn(x + 1, guess, wordAnswer)
			break
		} else {
			continue
		}
	}
	return winOrLose
}

// print results for all turns so far
func printTurn(currTurn int, guess string, answer string) {
	fmt.Println("Turn: " + strconv.Itoa(currTurn))
	if guess != "" {
		fmt.Println("Guess: " + guess)
		fmt.Println(answer)
		for _,x := range gameResults{
			//determines what letters match and do not match
			for i, y := range x{
				//checks if letters match
				if i == 0 {
					fmt.Print("[")
				}
				if string(answer[i]) == string(y) {
					//ANSI for green background, and resets format
					fmt.Print("\033[32m" +  string(y) + "\033[0m")
					
				} else if strings.Contains(answer, string(y)) {
					//ANSI for yellow background, and resets format
					fmt.Print("\033[43m" +  string(y) + "\033[0m")
				} else {
					fmt.Print(string(y))
				}

				if i != len(x) - 1{
					fmt.Print(" ")
				} 
				
			}
			fmt.Print("]")
			fmt.Println("")

		}
	} else {
		for _,x := range gameResults{
			fmt.Println(x)
		}
	}


}

func printWinLose(win bool){
	if win == true{
		fmt.Println("You win! Play again? (1 - Yes, 2 - No)")
	} else {
		fmt.Println("You loose! Play again? (1 - Yes, 0 - No)")
	}
}
