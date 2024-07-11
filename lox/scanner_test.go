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
		{name: "parens", source: "()", want: []Token{{0, "(", nil, 0}, {1, ")", nil, 0}}, wantError: false},
		{name: "error", source: "@", want: []Token{}, wantError: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			itpr := new(Interpreter)
			s := NewScanner(tc.source)
			tokens := s.scanTokens(itpr)
			// TODO: replace with google cmp package
			// this skips the EOF char in tokens
			if !slices.Equal(tc.want, tokens[:len(tokens)-1]) {
				t.Errorf("Failed to scan tokens. Want %+v, got %+v", tc.want, tokens)
			}
			if tc.wantError != itpr.hadError {
				t.Errorf("Expected error: %v, got error: %v", tc.wantError, itpr.hadError)
			}
		})
	}
}
