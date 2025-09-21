package main

import (
	"fmt"
	"math/rand"
	"strconv"
	// "math/rand"
)

var wordList = []string{
	"FETCH",
	"PARRY",
	"LOINS",
	"LUNGS",
	"TWIST",
}

// var streak int = 0
const Turns int = 6
const letterCount int = 5
var gameResults [Turns][letterCount]string

func init() {
	for i := range gameResults {
		for j := range gameResults[i]{
			gameResults[i][j] = "_"
		}
	}
}

func main() {
	var GameMode int

	for true {
		printIntro()
		fmt.Scan(&GameMode)
		if GameMode == 1 {
			//play game
			playGame()
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

func playGame() {
	//randomizes word
	wordAnswer := wordList[rand.Intn(len(wordList))]
	//sets how many turns the player should play. Default 6.
	var gameTurns int = len(wordList)

	//print turn
	for x := 0; x < gameTurns; x++ {
		var guess string
		printTurn(x + 1)
		fmt.Print("> ")
		fmt.Scan(&guess)
		fmt.Println("")
		if guess == wordAnswer {
			break
		}
	}

}

// print results for all turns so far
func printTurn(currTurn int) {
	fmt.Println("Turn: " + strconv.Itoa(currTurn))
	for _,x := range gameResults{
		fmt.Println(x)
	}
}
