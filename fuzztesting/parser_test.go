package main

import "testing"

// Fuzz test - Go generates random strings
func FuzzParseNumber(f *testing.F) {
	// Add seed corpus (initial test cases)
	f.Add("123")
	f.Add("0")
	f.Add("9999")

	// Go will run this with thousands of random inputs
	f.Fuzz(func(t *testing.T, input string) {
		// This might find bugs with weird inputs!
		result, _ := ParseNumberSafe(input)

		// Basic sanity check
		if input == "123" && result != 123 {
			t.Errorf("expected 123, got %d", result)
		}
	})
}
