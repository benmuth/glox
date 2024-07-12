package lox

import "fmt"

// Token stores the data for tokens
type Token struct {
	// type_ is the value of the "enum" for the Token
	Type_ int
	// a lexeme is the smallest unit of text that's meaningful to lox
	Lexeme string
	// TODO: figure out what Object literal is
	// final Object literal;
	Literal any
	// the line number the Token was found on
	Line int
}

func (t Token) String() string {
	// TODO: check whether this fulfills the Stringer interface
	return fmt.Sprintf("%s %s %s", tokenNames[t.Type_], t.Lexeme, t.Literal)
}
