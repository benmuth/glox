package lox

import "fmt"

type Token struct {
	_type  int
	lexeme string
	// TODO: figure out what Object literal is
	// final Object literal;
	literal error
	line    int
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %e", tokenNames[t._type], t.lexeme, t.literal)
}
