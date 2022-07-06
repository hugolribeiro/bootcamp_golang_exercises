package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	args := os.Args
	amountArgs := len(args) - 1
	if amountArgs == 0 {
		fmt.Println("Give me args")
	} else if amountArgs == 1 {
		fmt.Printf("There is one: \"%s\"\n", args[1])
	} else if amountArgs == 2 {
		fmt.Printf("There are two: \"%s %s\"\n", args[1], args[2])
	} else {
		fmt.Printf("There are %d arguments\n", amountArgs)
	}
}
