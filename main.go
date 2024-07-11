package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		err := runFile(args[0])
		if err != nil {
			fmt.Println("Couldn't read script contents")
		}
	} else {
		runPrompt()
	}
}

// runFile runs a glox script in a file
func runFile(fileName string) error {
	fmt.Println(fileName)
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	run(content)
	return nil
}

// runPrompt runs glox interactively
func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		scanner.Scan()
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}
		run(line)
	}
}

// run actually runs the interpreter
func run(content []byte) {
	tokens := bytes.Split(content, []byte{' ', '\n', '\t', '\r'})

	for _, token := range tokens {
		fmt.Println(string(token))
	}
}
