package nlp

import (
	"fmt"
)

func ExampleTokenize() {
	tokens := Tokenize("Who's on first?")
	fmt.Println(tokens)

	// Output:
	// [who on first]
}
