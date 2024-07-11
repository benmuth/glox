package lox

import (
	"slices"
	"testing"
)

func TestScanner(t *testing.T) {
	tests := []struct {
		name      string
		source    string
		want      []Token
		wantError bool
	}{
		{name: "parens", source: "(( )){} // grouping stuff",
			want: []Token{
				{LEFT_PAREN, "(", nil, 0},
				{LEFT_PAREN, "(", nil, 0},
				{RIGHT_PAREN, ")", nil, 0},
				{RIGHT_PAREN, ")", nil, 0},
				{LEFT_BRACE, "{", nil, 0},
				{RIGHT_BRACE, "}", nil, 0},
			},
			wantError: false},
		{name: "unexpected character", source: "@", want: []Token{}, wantError: true},
		{name: "operators", source: "!*+-/=<> <= == // operators",
			want: []Token{
				{BANG, "!", nil, 0},
				{STAR, "*", nil, 0},
				{PLUS, "+", nil, 0},
				{MINUS, "-", nil, 0},
				{SLASH, "/", nil, 0},
				{EQUAL, "=", nil, 0},
				{LESS, "<", nil, 0},
				{GREATER, ">", nil, 0},
				{LESS_EQUAL, "<=", nil, 0},
				{EQUAL_EQUAL, "==", nil, 0},
			},
			wantError: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			itpr := new(Interpreter)
			s := NewScanner(tc.source)
			tokens := s.scanTokens(itpr)
			// exclude eof
			tokens = tokens[:len(tokens)-1]
			// TODO: replace with google cmp package
			// this skips the EOF char in tokens
			if !slices.Equal(tc.want, tokens) {
				t.Errorf("Failed to scan tokens. Want %v, got %v", tc.want, tokens)
			}
			if tc.wantError != itpr.hadError {
				t.Errorf("Expected error: %v, got error: %v", tc.wantError, itpr.hadError)
			}
		})
	}
}
