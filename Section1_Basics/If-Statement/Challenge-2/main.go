package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #2
//  Add one more user to the PassMe program below.
//
// EXAMPLE USERS
//  username: jack
//  password: 1888
//
//  username: inanc
//  password: 1879
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 1888
//    Access granted to "jack".
//
//  go run main.go inanc 1879
//    Access granted to "inanc".
//
//  go run main.go jack 1879
//    Invalid password for "jack".
//
//  go run main.go inanc 1888
//    Invalid password for "inanc".
// ---------------------------------------------------------

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user1, pwd1 = "jack", "1888"
	user2, pwd2 = "inanc", "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	actualUsername, actualPassword := args[1], args[2]

	if actualUsername != user1 || actualUsername != user2 {
		fmt.Printf(errUser, actualUsername)
	} else if actualUsername == user1 && actualPassword == pwd1 {
		fmt.Printf(accessOK, actualUsername)
	} else if actualUsername == user2 && actualPassword == pwd2 {
		fmt.Printf(accessOK, actualUsername)
	} else {
		fmt.Printf(errPwd, actualUsername)
	}
}
