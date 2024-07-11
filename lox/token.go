package lox

import "fmt"

// Token stores the data for tokens
type Token struct {
	// type_ is the value of the "enum" for the Token
	type_ int
	// a lexeme is the smallest unit of text that's meaningful to lox
	lexeme string
	// TODO: figure out what Object literal is
	// final Object literal;
	literal error
	// the line number the Token was found on
	line int
}

func (t Token) String() string {
	// TODO: check whether this fulfills the Stringer interface
	return fmt.Sprintf("%s %s %e", tokenNames[t.type_], t.lexeme, t.literal)
}
