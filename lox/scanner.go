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
	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})
	return s.tokens
}

// scanToken advances the scanner, consuming the source code one character at
// a time
func (s *Scanner) scanToken(itpr *Interpreter) {
	c := s.advance()

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
	default:
		err := itpr.error(s.line, "Unexpected character.")
		if err != nil {
			panic("TODO: Handle error")
		}
	}
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

// TODO: figure out what literals are for
func (s *Scanner) addTokenLiteral(type_ int, literal error) {
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
