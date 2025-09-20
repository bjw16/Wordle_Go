package main

import "fmt"

var wordList = []string{
	"FETCH",
	"PARRY",
	"LOINS",
	"LUNGS",
	"TWIST",
}
func main(){
	for _, x := range wordList{
		fmt.Println(x)
	}
}