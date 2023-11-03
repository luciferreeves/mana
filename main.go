package main

import (
	"fmt"
	"mana/repl"
	"os"
	"os/user"
)

func main() {
	var user, err = user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to Mana REPL!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
