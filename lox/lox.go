// Package lox implements the lox programming language, as described in
// Crafting Interpreters, by Robert Nystrom.

package lox

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type Interpreter struct {
	hadError bool
}

// RunFile runs a glox script in a file
func (itpr *Interpreter) RunFile(fileName string) error {
	fmt.Println(fileName)
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	itpr.Run(content)
	if itpr.hadError {
		os.Exit(65)
	}
	return nil
}

// RunPrompt runs glox interactively
func (itpr *Interpreter) RunPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		scanner.Scan()
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}
		itpr.Run(line)
		itpr.hadError = false
	}
}

// Run actually runs the interpreter
func (itpr *Interpreter) Run(content []byte) {
	tokens := bytes.Split(content, []byte{' ', '\n', '\t', '\r'})

	for _, token := range tokens {
		fmt.Println(string(token))
	}
}

// error reports syntax errors
func (itpr *Interpreter) error(line int, message string) error {
	if err := itpr.report(line, "", message); err != nil {
		return err
	}
	return nil
}

// report writes an error message to stderr
func (itpr *Interpreter) report(line int, where, message string) error {
	_, err := io.WriteString(os.Stderr, fmt.Sprintf("[line %d] Error %s: %s\n", line, where, message))
	if err != nil {
		return err
	}
	itpr.hadError = true
	return nil
}
