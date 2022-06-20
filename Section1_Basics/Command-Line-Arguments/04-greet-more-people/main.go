package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	arguments := os.Args
	amountPeople := len(arguments) - 1
	if amountPeople == 0 {
		fmt.Println("There are no names here!")
		return
	}

	fmt.Printf("There are %d people!\n", amountPeople)
	names := arguments[1:]
	for _, name := range names {
		fmt.Printf("Hello great %s !\n", name)
	}
	fmt.Println("Nice to meet you all.")
}
