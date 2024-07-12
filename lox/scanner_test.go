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
		{name: "keyword", source: `or else`,
			want: []Token{
				{OR, "or", nil, 0},
				{ELSE, "else", nil, 0},
			},
			wantError: false},
		{name: "string with newline",
			source: `"hello
world"`,
			want: []Token{
				{STRING, "\"hello\nworld\"", "hello\nworld", 1},
			},
			wantError: false},
		{name: "block comment", source: `"hello"/*string*/`,
			want: []Token{
				{STRING, `"hello"`, "hello", 0},
			},
			wantError: false},
		{name: "nested block comment", source: `for /* /* nested comment */ */`,
			want: []Token{
				{FOR, "for", nil, 0},
			},
			wantError: false},
		{name: "unterminated block comment", source: `1.0 /* /* /* broken comment * */ */`,
			want: []Token{
				{NUMBER, "1.0", float64(1), 0},
			},
			wantError: true},
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
