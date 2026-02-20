package main

import "fmt"


func parserNumber(s string) int {
	result := 0
	for _, c := range s{
		result = result *10 + int(c-'0')
	}
	return result
}


func parserNmberSafe(s string) (int ,error){
	if s =="" {
		return 0, fmt.Errorf("empty string")
	}
	result := 0

	for _ , c := range s {
		if c <'0' || c >'9'{
            return 0,fmt.Errorf("invalid character:%q",c)
		}
		result = result * 10 + int(c-'0')
	}
	return result,nil
}