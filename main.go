package main

import (
	"fmt"
	"sort"
)

func main() {
	GuessingGame()
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d?\n", i)
		fmt.Scan(&s)
		//fmt.Printf("%s\n", s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
