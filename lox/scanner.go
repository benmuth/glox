package lox

// package com.craftinginterpreters.lox;

// import java.util.ArrayList;
// import java.util.HashMap;
// import java.util.List;
// import java.util.Map;

// import static com.craftinginterpreters.lox.TokenType.*;

// class Scanner {
//   private final String source;
//   private final List<Token> tokens = new ArrayList<>();
//   Scanner(String source) {
//     this.source = source;
//   }
// }

// List<Token> scanTokens() {
//     while (!isAtEnd()) {
//       // We are at the beginning of the next lexeme.
//       start = current;
//       scanToken();
//     }

//     tokens.add(new Token(EOF, "", null, line));
//     return tokens;
//   }

type Scanner struct {
	// the source code
	source string

	start, current, line int

	tokens []Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source}
}

func (s *Scanner) scanTokens() []Token {
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})
	return s.tokens
}

func (s *Scanner) scanToken() {
	panic("TODO: not implemented")
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
