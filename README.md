# Wordle_Go

A Go-based command line clone of the popular Wordle game, created as a project to learn Golang. 
Based on https://www.nytimes.com/games/wordle/index.html

## Features

- Command-line gameplay using ASCII interface.
- Guess the correct 5-letter word in 6 tries.
- Color-coded feedback for each guess:
  - **Green**: Correct letter in the correct position.
  - **Yellow**: Correct letter in the wrong position.


## How to Play

1. Install [Go](https://golang.org/dl/) if you haven't already.
2. Run the game with:
   ```
   go run wordle.go
   ```
3. Enter your guesses and try to solve the Wordle within 6 attempts!

## Game Visual

You can see an example of the game's output below.  

![Screen Recording 2025-09-28 at 11 43 05 PM](https://github.com/user-attachments/assets/2e731150-725c-499f-86b1-2281642d274d)




## Files

- `wordle.go` - Main game logic.
- `README.txt` - Original simple readme.
- `go.mod` - Go module definition.
- `wordle.exe` - (If present) Windows executable for the game.

## API
NOTE: If error with API, Wordle_Go switches to local wordlist
- https://github.com/RazorSh4rk/random-word-api.git - Used for pulling randomized words as potential Wordles/answers

## License

This project is for learning purposes and has no specific license.


## Note
Readme generated with Co-Pilot
