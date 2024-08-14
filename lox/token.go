package lox

import "fmt"

// Token stores the data for tokens.
type Token struct {
	// Type is the value of the "enum" for the Token.
	Type int
	// A Lexeme is the smallest unit of text that's meaningful to lox.
	Lexeme string
	// Literal is the value of the Type, if the Type can have different values,
	// like strings or floats.
	Literal any
	// Line is the line number the Token was found on.
	Line int
}

func (t Token) String() string {
	// TODO: check whether this fulfills the Stringer interface
	return fmt.Sprintf("%s %s %s", tokenNames[t.Type], t.Lexeme, t.Literal)
}
