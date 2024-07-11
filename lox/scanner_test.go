package lox

import (
	"slices"
	"testing"
)

func TestScanner(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   []Token
	}{
		{name: "parens", source: "()", want: []Token{{0, "(", nil, 0}, {1, ")", nil, 0}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewScanner(tc.source)
			tokens := s.scanTokens()
			// TODO: replace with google cmp package
			// this skips the EOF char in tokens
			if !slices.Equal(tc.want, tokens[:len(tokens)-1]) {
				t.Errorf("Failed to scan tokens. Want %+v, got %+v", tc.want, tokens)
			}
		})
	}
}
