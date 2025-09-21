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
var Turns int
var letterCount int
var gameResults [][]string

func init() {
	Turns = 6
	letterCount = 5
	gameResults := make([][]string, Turns, 100)
	for i := range gameResults {
		gameResults[i] = append(gameResults[i], "     ")
	}
}

func main() {
	var GameMode int
	Turns = 6
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
		printTurn(x)
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
	fmt.Println("Turn: " + strconv.Itoa(currTurn+1))
	fmt.Println(gameResults[1])
	// for i := 0; i < Turns; i++ {
	// 	fmt.Println(gameResults[i])
	// 	// for v2 := range gameResults[v1] {
	// 	// 	fmt.Print("[")
	// 	// 	fmt.Print(v2)
	// 	// 	fmt.Print("]")
	// 	// }
	// }
}
