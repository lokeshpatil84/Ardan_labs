package nlp

import (
	"os"
	"slices"
	"testing"
	
)


type tokCase struct{
	Test string
	Tokens []string
	Name string
}


func loadTokenizeCases(t *testing.T) []tokCase {
     file , err := os.Open("testdata/tokenize_cases.toml")
	 require.NoError(t, err)
	 defer file.Close()
     
	 var data struct{
		Cases []tokCase `toml:"case"`
	 }
	 dec := toml.NewDecoder(file)
	 _ , err := dec.decode(&data)
	 require.NoError(t, err)
	 return data.Cases
}
func TestTokenize(t *testing.T) {

	// var cases = []struct {
    //          text string
	// 		 tokens []string
	// }{
	// 	{"who is on firsr?", []string{"who","s","on","first"}},
	// 	{"whats's on second?", []string{"what","s","on","second"}},
	// 	{"",nil},
	// }
    cases := loadTokenizeCases(t)
	for _, tc := range cases{
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			if !slices.Equal(tc.tokens,tokens){
				t.Fatalf("expected %#v, got %#v",tc.tokens,tokens)
			}
		})
	}
}
