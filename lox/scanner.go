package lox

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
// a time
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
			// A comment goes until the end of the line.
			// We don't want to add these tokens.
			// We also don't want to consume the new line here, because when it's
			// consumed, we increment the line counter.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ', '\r', '\t':
		// Ignore whitespace
	case '\n':
		s.line++
	case '"':
		s.string(itpr)
	default:
		err := itpr.error(s.line, "Unexpected character.")
		if err != nil {
			panic("TODO: Handle error")
		}
	}
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
	value := s.source[s.start+1: s.current-1]
	s.addTokenLiteral(STRING, value)
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

// peek looks ahead one character without consuming it.
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
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
	s.addTokenLiteral(type_, "")
}

// TODO: figure out what literals are for
func (s *Scanner) addTokenLiteral(type_ int, literal string) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens,
		Token{type_: type_,
			lexeme:  text,
			literal: literal,
			line:    s.line})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
