package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// API URL - call for random word of length 5
// https://github.com/RazorSh4rk/random-word-api.git
var url = "https://random-word-api.herokuapp.com/word?length=5"

type apiResponse struct {
	wordle string
}

// Array of strings for available Wordle/answers
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

// Map used to track available letters
var availableLetters map[rune]int = make(map[rune]int)

// Sorting order to print available letters in order as seen
// on wordle keyboard
// Q W E R T Y U I O P
// A S D F G H J K L
//
//	Z X C V B N M
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

// Would like to make it where game is dynamic. Players can set own turns or lettercount
// default - 6 turns, 5 possible letter in wordle/answer
const turns int = 6
const letterCount int = 5

// streak counters
var streaksCounter int = 0

// 2D array of strings to track each guess
var guessArray [turns][letterCount]string

// Used to set/reset game results and available letters after games
func setArrays() {
	//load guess array with '_' for empty letter formatting
	for i := range guessArray {
		for j := range guessArray[i] {
			guessArray[i][j] = "_"
		}
	}
	//Available letters set to 0 as they have not been selected
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
	//0 = Don't play again. Load menu
	//1 = Play again. Don't load menu.
	var playAgain int = 0

	//Menu - runs infinite loop until break
	for true {
		//Menu setting used to track menu selection
		var MenuSettingInput string //Saves user input
		var MenuSettingError error  //Determine errors upon string -> int conversion into MenuSettingInt
		var MenuSettingInt int      //Menu user response converted to int

		//tracks [win - t] or [lose - f] results
		var results bool

		//response,error for input "Play again?" responses
		var playAgainResponse string
		var playAgainError error

		//0 = Don't play again. Load menu
		//1 = Play again. Don't load menu.
		if playAgain != 1 {
			//prints menu options
			printMenu()

			//read in menu selection, converts string input to int
			fmt.Scan(&MenuSettingInput)
			MenuSettingInt, MenuSettingError = strconv.Atoi(MenuSettingInput)
		}

		//If user selected 1 or decided to play again last game set by previous loop through
		if MenuSettingInt == 1 || playAgain == 1 {
			//reset playAgain value
			playAgain = 0

			//load game data structures
			setArrays()

			//Play's game. Returns results
			results = playGame()

			//Play again input handler
			for true {
				//print game results
				printWinLose(results)
				fmt.Print("> ")

				//record input
				fmt.Scan(&playAgainResponse)

				//convert input to int. Manage errors upon conversion.
				//Saves
				playAgain, playAgainError = strconv.Atoi(playAgainResponse)

				//Determines if input error free. If so, try again
				//i.e. not empty, and either 0 or 1
				//0 = Don't play again. Load menu
				//1 = Play again. Don't load menu.
				if playAgainError != nil || playAgain >= 2 || playAgain < 0 {
					fmt.Println("No associated input. Please try again!")
					fmt.Println("")
					continue
				} else {
					//breaks loop if input is correct
					break
				}
			}

		} else if MenuSettingInt == 2 && MenuSettingError == nil { //MenuSettingInt == 2 - Streaks Menu Option
			//print stats for 1 second
			fmt.Println("Streak: " + strconv.Itoa(streaksCounter))
			time.Sleep(1 * time.Second)
			//Exit game menu option
		} else if MenuSettingInt == 0 && MenuSettingError == nil { //MenuSettingInt == 0 and no errors - break loop. Program ends
			break

		} else { //Incorrect input
			fmt.Println("No associated input. Please try again!")
		}
		//reset input responses
		playAgainResponse = ""
		MenuSettingInput = ""
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

// function to play full game
// return 1 for win, 0 for lose
func playGame() bool {
	var wordle string
	apiResponse, err := http.Get("https://random-word-api.herokuapp.com/word?length=5")
	//if error, load local wordlist
	if err != nil {
		//randomizes anwer to current game from WordList array
		//https://www.geeksforgeeks.org/go-language/generating-random-numbers-in-golang/
		rand.Seed(time.Now().UnixNano())
		wordle = wordList[rand.Intn(len(wordList))]
	} else {
		defer apiResponse.Body.Close()
		body, err := io.ReadAll(apiResponse.Body)
		if err != nil {
			//randomizes anwer to current game from WordList array
			//https://www.geeksforgeeks.org/go-language/generating-random-numbers-in-golang/
			rand.Seed(time.Now().UnixNano())
			wordle = wordList[rand.Intn(len(wordList))]
		} else {
			wordle = strings.ToUpper(string(body))
		}
	}

	var winOrLose bool
	//Initializes guess to empty string
	var guess string = ""

	//main game loop. Plays until out of turns.
	for x := 0; x < turns; x++ {
		//Prints current turn, and available letters to choose from
		printTurn(x+1, guess, wordle)
		printAvailableLetters()
		//resets guess between turns
		guess = ""

		//guess input handler
		for {
			//Print response prompt
			fmt.Println("")
			fmt.Print("> ")

			//Read in guess from user
			fmt.Scan(&guess)

			//if guess has too many or too few letters - error
			if len(guess) < letterCount || len(guess) > letterCount {
				//print turn agian with error message
				printTurn(x, guess, wordle)
				printAvailableLetters()
				if len(guess) < letterCount {
					fmt.Println("Guess to small. Try again!")
				} else if len(guess) > letterCount {
					fmt.Println("Guess to big. Try again!")
				}
				//reset guess input
				guess = ""

			} else if isAllLetterInString(guess) == false { //Error check: Number input
				//print turn agian with error
				printTurn(x, guess, wordle)
				printAvailableLetters()
				fmt.Println("Guess contains numbers. Try again!")

				//reset guess input
				guess = ""

			} else { //No issue. Break loop.
				break
			}
		}
		//formatting
		fmt.Println("")

		//stores guess into guessArray
		for y := 0; y < letterCount; y++ {
			guessArray[x][y] = string(guess[y])
		}

		//Determines if guess matches wordle i.e. answer
		if strings.ToUpper(guess) == wordle { //Win and break play loop. Print turn for final time to review all guesses.
			winOrLose = true
			printTurn(x+1, guess, wordle)
			break
			//out of guesses
			//loose and break play loop
		} else if x == turns-1 { //Loose and break play loop. Print turn for final time to review all guesses.
			winOrLose = false
			printTurn(x+1, guess, wordle)
			break
			//Game not over. Continue game loop
		} else { //Determines which letters match if guess was not equal to wordle i.e. answer
			checkGuessMatch(wordle, strings.ToUpper(guess))
			continue
		}
	}
	//Print winning wordle once game ends
	fmt.Println("Answer: " + wordle)

	//return results
	//Win - true
	//Lose - false
	return winOrLose
}

// print this turn after guess
// loads guess array
// prints guess array in [_ _ _ _ _] format
// prints each letter in guess array using ANSI escape codes: green, yellow or white
// Turn: 1
// [_ _ _ _ _]
// [_ _ _ _ _]
// [_ _ _ _ _]
// [_ _ _ _ _]
// [_ _ _ _ _]
// [_ _ _ _ _]
// Q W E R T Y U I O P
// A S D F G H J K L
//
//	Z X C V B N M
func printTurn(currTurn int, guess string, answer string) {
	//prints current turn number
	fmt.Println("Turn: " + strconv.Itoa(currTurn))

	//checks if user hasn't guessed yet
	if guess != "" {
		//prints this rounds guess
		fmt.Println("Guess: " + strings.ToUpper(guess))

		//Loops through each guess in guessArray one letter at a time
		for _, x := range guessArray { //Retrieve each word guess in array
			xArrayToString := strings.Join(x[:], "")
			for i, y := range x { //prints each letter y of word x, by color
				//prints first "[" at beginning of array
				if i == 0 {
					fmt.Print("[")
				}

				//https://www.dolthub.com/blog/2024-02-23-colors-in-golang/
				//prints each letter using ANSI escape codes
				if string(answer[i]) == strings.ToUpper(string(y)) { //letter in correct position
					//ANSI for green background, and resets format
					fmt.Print("\033[32m" + strings.ToUpper(string(y)) + "\033[0m")

				} else if strings.Contains(string(answer), strings.ToUpper(string(y))) { //letter in guess but incorrect position
					//Determines if letter is in word more than once.
					//If so, only put yellow on first time letter appears in guess
					if strings.Count(xArrayToString, strings.ToUpper(string(y))) > 1 {
						if strings.Index(xArrayToString, strings.ToUpper(string(y))) == i { //checks if 'i' is currently currently at the first instance position of the letter to turn yellow
							fmt.Print("\033[33m" + strings.ToUpper(string(y)) + "\033[0m")
						} else { //mark any letter other than first instance to white
							fmt.Print(strings.ToUpper(string(y)))
						}
					} else { //letter is in word only once.
						if strings.Index(string(xArrayToString), strings.ToUpper(string(y))) == i { //mark letter yellow if only used in word once
							fmt.Print("\033[33m" + strings.ToUpper(string(y)) + "\033[0m")
						} else { //mark any letter other than first instance to white
							fmt.Print(strings.ToUpper(string(y)))
						}
					}
				} else { //letter is not in guess
					//mark letter white
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
		for _, x := range guessArray {
			fmt.Println(x)
		}
	}
}

// prints rewsults. Asks to play again.
// 1 - yes. 2 - no.
func printWinLose(win bool) {
	if win == true {
		fmt.Println("You win! Play again? (1 - Yes, 0 - No)")
		streaksCounter += streaksCounter + 1
	} else {
		fmt.Println("You loose! Play again? (1 - Yes, 0 - No)")
		streaksCounter = 0
	}
}

// Compare guess to wordle i.e. answer
// Codes printing of available letters
// 1 - green
// 2 - yellow
// -1 - white
func checkGuessMatch(answer string, guess string) {
	//compare each letter of guess to answer
	for n, x := range guess {
		//checks if letter is in guess
		if strings.Contains(answer, string(x)) {
			//if guess[n] == answer[n]
			if rune(answer[n]) == x { //Letter in wordle/answer and in right place
				//turn letter green - [1]
				availableLetters[x] = 1
			} else { //Letter in wordle/answer but not in right place
				//used to ensure green letter never turns yellow after turning green
				if availableLetters[x] != 1 {
					//turns letter yellow - [2]
					availableLetters[x] = 2
				}
			}
		} else { //letter not in wordle/answer, and is now already been used
			//turn white - [-1]
			availableLetters[x] = -1
		}
	}
}

// prints available letters based on color code
// 1 - green background
// 2 - yellow background
// -1 - white background
// Print colors in ANSI escape code
func printAvailableLetters() {
	//sorts through letters and color code
	for _, x := range sortOrderavailableLetters {
		if availableLetters[x] == 0 { //what to print when letter was never used in guess yet
			//white background color
			fmt.Print("\033[47m" + string(x) + " " + "\033[0m")
		} else if availableLetters[x] == 2 { // 2 - yellow background
			fmt.Print("\033[43m" + string(x) + " " + "\033[0m")
		} else if availableLetters[x] == 1 { // 1 - green background
			fmt.Print("\033[42m" + string(x) + " " + "\033[0m")
		} else if availableLetters[x] == -1 { // -1 - white background
			fmt.Print(string(x) + " ")
		}

		//Prints new line at last letter for keyboard formatting
		//Q W E R T Y U I O P - here
		//A S D F G H J K L - here
		//here - Z X C V B N M
		if x == 'P' {
			fmt.Println("")
		} else if x == 'L' {
			fmt.Println("")
			fmt.Print("  ")
		}
	}
	//formatting
	fmt.Println("")
}

// checks that only letters used in guess response
// true - only letters, correct response
// false - contains characters other than letters, error
func isAllLetterInString(x string) bool {
	for _, y := range x {
		if unicode.IsLetter(y) == false {
			return false
		}
	}
	return true
}
