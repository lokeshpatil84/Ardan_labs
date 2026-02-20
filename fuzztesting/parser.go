// parser.go
package main

import "fmt"

func ParseNumber(s string) int {
	// This might panic on invalid input!
	result := 0
	for _, c := range s {
		result = result*10 + int(c-'0')
	}
	return result
}

func ParseNumberSafe(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("empty string")
	}
	result := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid character: %q", c)
		}
		result = result*10 + int(c-'0')
	}
	return result, nil
}
