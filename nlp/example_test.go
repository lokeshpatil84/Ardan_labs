package nlp

import (
	"fmt"

	
)

func ExampleTokenize() {
	tokens := nlp.Tokenize("Who's on first?")
	fmt.Println(tokens)

	// Output:
	// [who on first]
}