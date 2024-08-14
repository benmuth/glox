package lox

import (
	"strconv"
)

var keywords map[string]int = map[string]int{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Scanner struct {
	// the source code
	source string

	start, current, line int

	tokens []Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source}
}

// scanTokens reads the source into a list of tokens
func (s *Scanner) scanTokens(itpr *Interpreter) []Token {
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.start = s.current
		s.scanToken(itpr)
	}
	s.tokens = append(s.tokens, Token{EOF, "", "", s.line})
	return s.tokens
}

// scanToken advances the scanner, consuming the source code one character at
// a time.
func (s *Scanner) scanToken(itpr *Interpreter) {
	c := s.advance()

	type_ := 0
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			type_ = BANG_EQUAL
		} else {
			type_ = BANG
		}
		s.addToken(type_)
	case '=':
		if s.match('=') {
			type_ = EQUAL_EQUAL
		} else {
			type_ = EQUAL
		}
		s.addToken(type_)
	case '<':
		if s.match('=') {
			type_ = LESS_EQUAL
		} else {
			type_ = LESS
		}
		s.addToken(type_)
	case '>':
		if s.match('=') {
			type_ = GREATER_EQUAL
		} else {
			type_ = GREATER
		}
		s.addToken(type_)
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line. We don't want to add tokens in
			// a comment.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else if s.match('*') {
			s.blockComment(itpr)
		} else {
			s.addToken(SLASH)
		}
	case '\n':
		s.line++
	case ' ', '\r', '\t': // Ignore other whitespace
	case '"':
		s.string(itpr)
	case 'o':
		// BUG: figure out whether "orange" is and should be a valid identifier
		if s.match('r') {
			s.addToken(OR)
		}
		// NOTE: should there be an "AND" identifier?
	default:
		if isDigit(c) {
			s.number(itpr)
		} else if isAlpha(c) {
			s.identifier(itpr)
		} else {
			err := itpr.error(s.line, "Unexpected character.")
			if err != nil {
				panic("TODO: Handle error")
			}
		}
	}
}

func (s *Scanner) identifier(itpr *Interpreter) {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	type_, ok := keywords[text]

	if !ok {
		type_ = IDENTIFIER
	}

	s.addToken(type_)
}

func (s *Scanner) number(itpr *Interpreter) {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	parsedFloat, err := strconv.ParseFloat(s.source[s.start:s.current], 32)
	if err != nil {
		// TODO: error handling
		panic(err)
	}
	s.addTokenLiteral(NUMBER, parsedFloat)
}

func (s *Scanner) string(itpr *Interpreter) {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		itpr.error(s.line, "Unterminated string.")
		return
	}

	// consume the closing "
	s.advance()

	// trim the surrounding quotes
	value := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(STRING, value)
}

func (s *Scanner) blockComment(itpr *Interpreter) {
	nestCount := 1
	for nestCount > 0 {
		for (s.peek() != '*' && s.peek() != '/') && !s.isAtEnd() {
			s.advance()
		}
		if s.isAtEnd() {
			itpr.error(s.line, "Unterminated block comment")
			break
		}
		if s.match('*') {
			if s.match('/') {
				nestCount--
			}
		}
		if s.match('/') {
			if s.match('*') {
				nestCount++
			}
		}
	}
}

// match checks if the current character matches expected. If it does, the
// character is consumed.
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

// peek looks at the current character without consuming it.
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

// peekNext looks at the next character without consuming it
func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isDigit(c) || isAlpha(c)
}

// advance moves the current position of the scanner forward and returns the
// character found there
func (s *Scanner) advance() byte {
	currChar := s.source[s.current]
	s.current++
	return currChar
}

// addToken adds the current token to the list of parsed tokens
func (s *Scanner) addToken(type_ int) {
	s.addTokenLiteral(type_, nil)
}

func (s *Scanner) addTokenLiteral(type_ int, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens,
		Token{Type: type_,
			Lexeme:  text,
			Literal: literal,
			Line:    s.line})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
