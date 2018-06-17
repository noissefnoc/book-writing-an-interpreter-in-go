package main

import (
	user2 "os/user"
	"fmt"
	"github.com/noissefnoc/book-writing-an-interpreter-in-go/monkey/repl"
	"os"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programing language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
