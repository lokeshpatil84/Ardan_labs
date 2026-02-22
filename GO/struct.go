package main

import "fmt"

type example struct {
	counter int64
	pi      float64
	flag    bool
}

func main() {
	// var e1 example

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.14,
	}
	fmt.Printf("%+v\n", e2)
}
