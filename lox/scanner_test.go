package lox

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
				{LEFT_PAREN, "(", "", 0},
				{LEFT_PAREN, "(", "", 0},
				{RIGHT_PAREN, ")", "", 0},
				{RIGHT_PAREN, ")", "", 0},
				{LEFT_BRACE, "{", "", 0},
				{RIGHT_BRACE, "}", "", 0},
			},
			wantError: false},
		{name: "unexpected character", source: "@", want: []Token{}, wantError: true},
		{name: "operators", source: "!*+-/=<> <= == // operators",
			want: []Token{
				{BANG, "!", "", 0},
				{STAR, "*", "", 0},
				{PLUS, "+", "", 0},
				{MINUS, "-", "", 0},
				{SLASH, "/", "", 0},
				{EQUAL, "=", "", 0},
				{LESS, "<", "", 0},
				{GREATER, ">", "", 0},
				{LESS_EQUAL, "<=", "", 0},
				{EQUAL_EQUAL, "==", "", 0},
			},
			wantError: false},
		{name: "string", source: `"hello world" // string`,
			want: []Token{
				{STRING, "\"hello world\"", "hello world", 0},
			},
			wantError: false},
		{name: "number", source: `1.25 // number`,
			want: []Token{
				{NUMBER, "1.25", 1.25, 0},
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
			if diff := cmp.Diff(tc.want, tokens); diff != "" {
				t.Errorf("Token mismatch: (-want +got):\n%s", diff)
			}

			if tc.wantError != itpr.hadError {
				t.Errorf("Expected error: %v, got error: %v", tc.wantError, itpr.hadError)
			}
		})
	}
}
