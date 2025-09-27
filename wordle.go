package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var wordList = []string{
	"APPLE",
	"BERRY",
	"CHARM",
	"DREAM",
	"EAGLE",
	"FLAME",
	"GRAPE",
	"HOUSE",
	"IVORY",
	"JOLLY",
	"KINGS",
	"LEMON",
	"MAGIC",
	"NOBLE",
	"OCEAN",
	"PLANT",
	"QUICK",
	"RIVER",
	"STONE",
	"TIGER",
	"VIVID",
	"WHEAT",
	"YIELD",
	"ZEBRA",
	"BLAZE",
	"CRANE",
	"DRIFT",
	"FROST",
	"GLINT",
	"HAVEN",
	"INLET",
	"JEWEL",
	"KNOLL",
	"LARCH",
	"MIRTH",
	"NEXUS",
	"OLIVE",
	"PRISM",
	"QUEST",
	"RIDGE",
	"SPARK",
	"THORN",
	"UNITY",
	"VALOR",
	"WILLOW",
	"XENON",
	"YOUTH",
	"ZIPPY",
	"ALPHA",
	"BRAVO",
}
var possibleLetters map[rune]int = make(map[rune]int)

var sortOrderPossibleLetters = []rune{
	'Q',
	'W',
	'E',
	'R',
	'T',
	'Y',
	'U',
	'I',
	'O',
	'P',
	'A',
	'S',
	'D',
	'F',
	'G',
	'H',
	'J',
	'K',
	'L',
	'Z',
	'X',
	'C',
	'V',
	'B',
	'N',
	'M',
}

var gameResults [Turns][letterCount]string

// var streak int = 0
const Turns int = 6
const letterCount int = 5

var streaks int = 0

func setArrays() {
	for i := range gameResults {
		for j := range gameResults[i] {
			gameResults[i][j] = "_"
		}
	}
	possibleLetters = map[rune]int{
		'Q': 0,
		'W': 0,
		'E': 0,
		'R': 0,
		'T': 0,
		'Y': 0,
		'U': 0,
		'I': 0,
		'O': 0,
		'P': 0,
		'A': 0,
		'S': 0,
		'D': 0,
		'F': 0,
		'G': 0,
		'H': 0,
		'J': 0,
		'K': 0,
		'L': 0,
		'Z': 0,
		'X': 0,
		'C': 0,
		'V': 0,
		'B': 0,
		'N': 0,
		'M': 0,
	}
}

func main() {
	var GameMode int
	var playAgain bool = false

	for true {
		var results bool
		var playAgainResponse int
		if playAgain != true {
			printIntro()
			fmt.Scan(&GameMode)
		}
		if GameMode == 1 || playAgain == true {
			//play game
			setArrays()
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
			fmt.Println("Streak: " + strconv.Itoa(streaks))
			time.Sleep(1 * time.Second)

		} else if GameMode == 0 {
			//error or exit
			break
		} else {
			fmt.Println("No associated input. Please try again!")
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

// function to play full game - module
func playGame() bool {
	//randomizes word
	wordAnswer := wordList[rand.Intn(len(wordList))]
	var winOrLose bool
	var guess string = ""

	//print turn
	for x := 0; x < Turns; x++ {
		//var playAgain int
		printTurn(x+1, guess, wordAnswer)
		printPossibleLetters()
		//resets guess
		guess = ""

		//player guess response
		for {
			fmt.Println("")
			fmt.Print("> ")
			fmt.Scan(&guess)
			if len(guess) < letterCount || len(guess) > letterCount {
				printTurn(x, guess, wordAnswer)
				if len(guess) < letterCount {
					fmt.Println("Guess to small, try again!")
				} else  if len(guess) > letterCount {
					fmt.Println("Guess to big, try again!")
				} 
				printPossibleLetters()
				guess = ""
			} else if isAllLetterInString(guess) == false{
				fmt.Println("Guess contains numbers. Try again!")
				printPossibleLetters()
				guess = ""
			} else {
				//breaks look if no issues
				break
			}
		}
		fmt.Println("")
		//need error check if less than 5 letters

		//stores guess into results
		for y := 0; y < letterCount; y++ {
			gameResults[x][y] = string(guess[y])
		}

		//Determines if guess matches selected word
		if guess == wordAnswer {
			winOrLose = true
			printTurn(x+1, guess, wordAnswer)
			break
		} else if x == Turns-1 {
			winOrLose = false
			printTurn(x+1, guess, wordAnswer)
			break
		} else {
			//else, it's not the end of game. Continue playing game
			makePossibleLettersChoosen(wordAnswer, guess)
			continue
		}
	}
	fmt.Println("Answer: " + wordAnswer)
	return winOrLose
}

// print results for all turns so far
func printTurn(currTurn int, guess string, answer string) {
	fmt.Println("Turn: " + strconv.Itoa(currTurn))
	if guess != "" {
		fmt.Println("Guess: " + guess)
		for _, x := range gameResults {
			//determines what letters match and do not match
			for i, y := range x {
				//checks if letters match
				if i == 0 {
					fmt.Print("[")
				}

				//https://www.dolthub.com/blog/2024-02-23-colors-in-golang/
				if string(answer[i]) == strings.ToUpper(string(y)) {
					//ANSI for green background, and resets format
					fmt.Print("\033[32m" + strings.ToUpper(string(y)) + "\033[0m")

				} else if strings.Contains(answer, strings.ToUpper(string(y))) {
					//ANSI for yellow background, and resets format
					fmt.Print("\033[33m" + strings.ToUpper(string(y)) + "\033[0m")
				} else {
					fmt.Print(strings.ToUpper(string(y)))
				}

				if i != len(x)-1 {
					fmt.Print(" ")
				}

			}
			fmt.Print("]")
			fmt.Println("")

		}
	} else {
		for _, x := range gameResults {
			fmt.Println(x)
		}
	}
}

func printWinLose(win bool) {
	if win == true {
		fmt.Println("You win! Play again? (1 - Yes, 2 - No)")
		streaks += streaks + 1
	} else {
		fmt.Println("You loose! Play again? (1 - Yes, 0 - No)")
		streaks = 0
	}
}

func printPossibleLetters() {
	for _, x := range sortOrderPossibleLetters {
		if possibleLetters[x] == 0 {
			//what to print when letter wasn't chosen
			fmt.Print("\033[47m" + string(x) + " " + "\033[0m")
		} else if possibleLetters[x] == 2 {
			//what to print if letter was in word but not at right position
			fmt.Print("\033[43m" + string(x) + " " + "\033[0m")
		} else if possibleLetters[x] == 1 {
			//what to print if letter was in word and in right position
			fmt.Print("\033[42m" + string(x) + " " + "\033[0m")
		} else if possibleLetters[x] == -1 {
			//what to print when letter was choosen, but not in word
			fmt.Print(string(x) + " ")
		}

		//Print formatting adjustments at certain circumstances
		if x == 'P' {
			fmt.Println("")
		} else if x == 'L' {
			fmt.Println("")
			fmt.Print("  ")
		}
	}
	fmt.Println("")
}

// Essentially, if letter is chosen, we will make it lowercase
// this will help with printing
func makePossibleLettersChoosen(answer string, userChoice string) {
	fmt.Println(possibleLetters)
	for n, x := range userChoice {	
		if strings.Contains(answer, string(x)){
			if rune(answer[n]) == x {
				//turn green
				//Letter in word and in right place
				possibleLetters[x] = 1
			} else {
				//turns yellow
				//Letter in word but not in right place
				if possibleLetters[x] != 1 {
					possibleLetters[x] = 2
				}
			}
		} else {
			//turn white
			//letter not in word, and is now selected
			possibleLetters[x] = -1
		}
	}
}

