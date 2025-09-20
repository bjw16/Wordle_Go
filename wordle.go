package main

import "fmt"

var wordList = []string{
	"FETCH",
	"PARRY",
	"LOINS",
	"LUNGS",
	"TWIST",
}
var result = map[int]string {

	1: "[ ][ ][ ][ ][ ]",
	2: "[ ][ ][ ][ ][ ]",
	3: "[ ][ ][ ][ ][ ]",
	4: "[ ][ ][ ][ ][ ]",
	5: "[ ][ ][ ][ ][ ]",
	6: "[ ][ ][ ][ ][ ]",
}

func main(){
	var GameMode int
	printIntro()
	for true{
		printGameMode()
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
	}
}

func printIntro(){
	//game_loop
	fmt.Println("WORDLE_GO")
	fmt.Println("Enter options below:")
}

func printGameMode(){
	fmt.Println("1 - Play Wordle")
	fmt.Println("2 - Streaks")
}

func playGame() {
	for _, x := range result{
		fmt.Println(x)
	}
	fmt.Println(result)
}