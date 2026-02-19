package nlp

import (
	"slices"
	"testing"
)


func TestTokenize(t *testing.T) {

	var cases = []struct {
             text string
			 tokens []string
	}{
		{"who is on firsr?", []string{"who","s","on","first"}},
		{"whats's on second?", []string{"what","s","on","second"}},
		{"",nil},
	}

	for _, tc := range cases{
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			if !slices.Equal(tc.tokens,tokens){
				t.Fatalf("expected %#v, got %#v",tc.tokens,tokens)
			}
		})
	}
}
