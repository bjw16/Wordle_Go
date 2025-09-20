package main

import (
	"fmt"
	// "math/rand"
)

var wordList = []string{
	"FETCH",
	"PARRY",
	"LOINS",
	"LUNGS",
	"TWIST",
}
var gameResults map[int]string
var streak int = 0

var Turns int

func init(){
	gameResults = make(map[int]string)
	for i := 0 ; i < Turns; i++ {
		gameResults [i] =  "[ ][ ][ ][ ][ ]"
	}
}

func main(){
	var GameMode int
	Turns = 6
	for true{
		printIntro()
		fmt.Scan(&GameMode)
		if GameMode == 1{
			//play game
			playGame()
		} else if GameMode == 2{
			//show stats
		} else if GameMode == 0{
			//error or exit
			break
		}
		fmt.Println("")
	}
}

func printIntro(){
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

func printGameMode(){
	
}

func playGame() {
	// wordAnswer := wordList[rand.Intn(len(wordList))]
	// var gameTurns int = len(wordList)
	// for x := 0; x < gameTurns; x++{
	// 	fmt.Println(gameResults[x])
	// }
	// fmt.Println(wordAnswer)
}
