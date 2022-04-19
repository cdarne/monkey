package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/cdarne/monkey/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, Welcome to the Monkey programming language!\nFeel free to type any command :)\n", u.Username)
	repl.Start(os.Stdin, os.Stdout)
}
