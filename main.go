package main

import (
	"fmt"
	"os"
	"os/user"
	"mana/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to Mana REPL!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
