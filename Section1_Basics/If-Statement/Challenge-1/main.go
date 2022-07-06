package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

const (
	usage           = "Usage: [username] [password]"
	errUser         = "Access denied for %q.\n"
	errPwd          = "Invalid password for %q.\n"
	accessOK        = "Access granted to %q.\n"
	correctUsername = "jack"
	correctPassword = "1888"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	username, password := args[1], args[2]

	if username != correctUsername {
		fmt.Printf(errUser, username)
		return
	}

	if password != correctPassword {
		fmt.Printf(errPwd, password)
		return
	}

	fmt.Printf(accessOK, username)

}
