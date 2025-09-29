package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//Array of strings for available Wordle answers
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
	"XENON",
	"YOUTH",
	"ZIPPY",
	"ALPHA",
	"BRAVO",
}

//Map used to track available letters
var availableLetters map[rune]int = make(map[rune]int)

//Sorting order to print available letters in order as seen
//on wordle keyboard
var sortOrderavailableLetters = []rune{
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

//Would like to make it where game is dynamic. Players can set own turns or lettercountf
const Turns int = 6
const letterCount int = 5
var streaks int = 0

//2D array of strings to track each guess
var gameResults [Turns][letterCount]string

//used to set/reset game results and available letters after games
func setArrays() {
	for i := range gameResults {
		for j := range gameResults[i] {
			gameResults[i][j] = "_"
		}
	}
	availableLetters = map[rune]int{
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
	//Value used to track if player would like to play again
	var playAgain int = 0

	for true {
		//Game mode used to track menu selection
		var MenuSetting string
		var MenuSettingInt int
		var MenuSettingError error

		//tracks [win - t] or [lose - f] results
		var results bool

		//tracks if player would like to play again
		var playAgainResponse string
		var playAgainError error

		//Used to avoid printing menu
		//if playing again
		if playAgain != 1 {
			//prints menu options
			printMenu()

			//read in menu selection
			fmt.Scan(&MenuSetting)
			MenuSettingInt, MenuSettingError = strconv.Atoi(MenuSetting)
		}
		//If user selected 1 or decided to play again last game set by previous loop through
		if MenuSettingInt == 1 || playAgain == 1 {
			//reset playAgain value
			playAgain = 0

			//load game data structures
			setArrays()

			results = playGame()

			//Play again input handler
			for true {
				//print game results
				printWinLose(results)

				fmt.Print("> ")

				//record input
				fmt.Scan(&playAgainResponse)

				//convert input to int
				playAgain, playAgainError = strconv.Atoi(playAgainResponse)

				//Determines if input is correct
				if playAgainError != nil || playAgain >= 2 || playAgain < 0 {
					fmt.Println("No associated input. Please try again!")
					fmt.Println("")
					continue
				} else {
					//breaks loop if input is correct
					break
				}
			}
		//Streaks menu option
		} else if MenuSettingInt == 2 && MenuSettingError == nil {
			//show stats for 1 second
			fmt.Println("Streak: " + strconv.Itoa(streaks))
			time.Sleep(1 * time.Second)
		//Exit game menu option
		} else if MenuSettingInt == 0 && MenuSettingError == nil{
			break
		//Incorrect input handler
		} else {
			fmt.Println("No associated input. Please try again!")
		}
		//resets
		playAgainResponse = ""
		MenuSetting = ""
		fmt.Println("")
	}
}

func printMenu() {
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
// return 1 for win, 0 for lose
func playGame() bool {
	//randomizes games answer each time function is ran
	wordle := wordList[rand.Intn(len(wordList))]

	var winOrLose bool
	var guess string = ""

	//main game loop
	//play until out of turns
	for x := 0; x < Turns; x++ {
		//Print round x and available letters
		printTurn(x+1, guess, wordle)
		printAvailableLetters()
		//resets guess
		guess = ""

		//guess input handler
		for {
			fmt.Println("")
			fmt.Print("> ")

			//guess input
			fmt.Scan(&guess)

			//if quess has too many or too few letters
			if len(guess) < letterCount || len(guess) > letterCount {
				//print turn agian with error
				printTurn(x, guess, wordle)
				printAvailableLetters()
				if len(guess) < letterCount {
					fmt.Println("Guess to small, try again!")
				} else  if len(guess) > letterCount {
					fmt.Println("Guess to big, try again!")
				} 
				//reset
				guess = ""
			//Error check: Number input
			} else if isAllLetterInString(guess) == false{
				fmt.Println("Guess contains numbers. Try again!")
				printAvailableLetters()
				//reset
				guess = ""
			//breaks handler if no issues
			} else {
				break
			}
		}
		fmt.Println("")

		//stores guess into results
		for y := 0; y < letterCount; y++ {
			gameResults[x][y] = string(guess[y])
		}

		//Determines if guess matches wordle
		//Win and break play loop
		if strings.ToUpper(guess) == wordle {
			winOrLose = true
			printTurn(x+1, guess, wordle)
			break
		//out of guesses
		//loose and break play loop
		} else if x == Turns-1 {
			winOrLose = false
			printTurn(x+1, guess, wordle)
			break
		//Game not over. Continue game loop
		} else {
			checkGuessMatch(wordle, guess)
			continue
		}
	}
	//Print winning wordle
	fmt.Println("Answer: " + wordle)

	//return results
	return winOrLose
}

// print this rounds turn i.e. guess array
func printTurn(currTurn int, guess string, answer string) {
	//prints current turn number
	fmt.Println("Turn: " + strconv.Itoa(currTurn))

	//checks if user hasn't guessed yet
	if guess != "" {
		//prints previous guess
		fmt.Println("Guess: " + guess)

		//Prints game results simimlar to printing out native array
		//This allows me to color code each letter
		//Selects 1 guess at a time
		for _, x := range gameResults {
			//prints each letter of word x by color
			for i, y := range x {
				//prints first "[" at beginning of array
				if i == 0 {
					fmt.Print("[")
				}

				//https://www.dolthub.com/blog/2024-02-23-colors-in-golang/
				//prints each letter color coded
				if string(answer[i]) == strings.ToUpper(string(y)) {
					//ANSI for green background, and resets format
					//letter in correct position
					fmt.Print("\033[32m" + strings.ToUpper(string(y)) + "\033[0m")

				} else if strings.Contains(answer, strings.ToUpper(string(y))) {
					//ANSI for yellow background, and resets format
					//letter in word but incorrect possition
					fmt.Print("\033[33m" + strings.ToUpper(string(y)) + "\033[0m")
				} else {
					//letter is not a match
					fmt.Print(strings.ToUpper(string(y)))
				}

				//prints space for array display
				if i != len(x)-1 {
					fmt.Print(" ")
				}
			}
			//prints end bracket for each row/guess
			fmt.Print("]")
			fmt.Println("")

		}
	//prints out empty array before user has had first guess
	} else {
		for _, x := range gameResults {
			fmt.Println(x)
		}
	}
}


func printWinLose(win bool) {
	if win == true {
		fmt.Println("You win! Play again? (1 - Yes, 0 - No)")
		streaks += streaks + 1
	} else {
		fmt.Println("You loose! Play again? (1 - Yes, 0 - No)")
		streaks = 0
	}
}

//compare guess to answer
//Codes printing of available letters, set color coding
func checkGuessMatch(answer string, guess string) {
	//compare each letter of guess to answer
	for n, x := range guess {	
		//checks if letter is in guess
		if strings.Contains(answer, string(x)){
			//if guess[n] == answer[n] aka index match
			if rune(answer[n]) == x {
				//turn green - [1]
				//Letter in word and in right place
				availableLetters[x] = 1
			} else {
				//Letter in word but not in right place
				if availableLetters[x] != 1 {
					//turns yellow - [2]
					availableLetters[x] = 2
				}
			}
		//
		} else {
			//turn white - [-1]
			//letter not in word, and is now already been used
			availableLetters[x] = -1
		}
	}
}

//prints available letters - color coded
//set under guess array
func printAvailableLetters() {
	//sorts through letters and color code
	for _, x := range sortOrderavailableLetters {
		if availableLetters[x] == 0 {
			//what to print when letter wasn't chosen
			//white color
			fmt.Print("\033[47m" + string(x) + " " + "\033[0m")
		} else if availableLetters[x] == 2 {
			//what to print if letter was in word but not at right position
			//yellow
			fmt.Print("\033[43m" + string(x) + " " + "\033[0m")
		} else if availableLetters[x] == 1 {
			//what to print if letter was in word and in right position
			//green
			fmt.Print("\033[42m" + string(x) + " " + "\033[0m")
		} else if availableLetters[x] == -1 {
			//what to print when letter was choosen, but not in word
			fmt.Print(string(x) + " ")
		}

		//Prints new line at last letter for keyboard formatting
		if x == 'P' {
			fmt.Println("")
		} else if x == 'L' {
			fmt.Println("")
			fmt.Print("  ")
		}
	}
	fmt.Println("")
}

//checks if all characters are letters in a string
func isAllLetterInString (x string) bool{
	//check if any characters but letters in guess
	for _, y := range x {
		if unicode.IsLetter(y) == false {
			return false
		}
	}
	return true
}

