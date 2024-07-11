package main

import (
	"fmt"
	"github.com/benmuth/crafting-interpreters/glox/lox"
	"os"
)

func main() {
	args := os.Args[1:]

	itpr := new(lox.Interpreter)
	if len(args) > 1 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		err := itpr.RunFile(args[0])
		if err != nil {
			fmt.Println("Couldn't read script contents")
		}
	} else {
		itpr.RunPrompt()
	}
}
