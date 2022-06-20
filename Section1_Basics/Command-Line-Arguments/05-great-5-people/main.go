package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet 5 People
//
//  Greet 5 people this time.
//
//  Please do not copy paste from the previous exercise!
//
//
// INPUT
//  bilbo balbo bungo gandalf legolas
//
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	arguments := os.Args
	amountNames := len(arguments) - 1
	if amountNames == 0 {
		fmt.Println("There are no names here!")
		return
	}

	names := arguments[1:]
	for _, name := range names {
		fmt.Printf("Hello great %s !\n", name)
	}
	fmt.Println("Nice to meet you all.")
}
