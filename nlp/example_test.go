package nlp_test

import (
	"fmt"

	"github.com/ardanlabs/nlp"
)

func ExampleTokenize() {
	tokens := nlp.Tokenize("Who's on first?")
	fmt.Println(tokens)

	// Output:
	// [who on first]
}