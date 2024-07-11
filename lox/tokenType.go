package lox

// These represent the type of the token
const (
	// Single-character tokens.
	LEFT_PAREN = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

// Given the type of a token, this helps get the name of the token.
// Mostly for debugging.
var tokenNames map[int]string = map[int]string{
	// Single-character tokens.
	1:  "LEFT_PAREN",
	2:  "RIGHT_PAREN",
	3:  "LEFT_BRACE",
	4:  "RIGHT_BRACE",
	5:  "COMMA",
	6:  "DOT",
	7:  "MINUS",
	8:  "PLUS",
	9:  "SEMICOLON",
	10: "SLASH",
	11: "STAR",

	// One or two character tokens.
	12: "BANG",
	13: "BANG_EQUAL",
	14: "EQUAL",
	15: "EQUAL_EQUAL",
	16: "GREATER",
	17: "GREATER_EQUAL",
	18: "LESS",
	19: "LESS_EQUAL",

	// Literals.
	20: "IDENTIFIER",
	21: "STRING",
	22: "NUMBER",

	// Keywords.
	23: "AND",
	24: "CLASS",
	25: "ELSE",
	26: "FALSE",
	27: "FUN",
	28: "FOR",
	29: "IF",
	30: "NIL",
	31: "OR",
	32: "PRINT",
	33: "RETURN",
	34: "SUPER",
	35: "THIS",
	36: "TRUE",
	37: "VAR",
	38: "WHILE",

	39: "EOF",
}
