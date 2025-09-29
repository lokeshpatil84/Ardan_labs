package main

import (
	"fmt"
	"sort"
)

func main() {
	cart := []string{"apple", "bannana", "grape"}

	fmt.Println(cart)
	fmt.Println(len(cart))
	fmt.Println(cart[0])

	//index + value
	for i, c := range cart {
		fmt.Println(i, c)
	}

	//only value
	for _, c := range cart {
		fmt.Println(c)
	}

	cart = append(cart, "orange")
	fmt.Println(cart)

	//slicing operator ,half =open
	fruit := cart[:2]
	fmt.Println(fruit)
	//---------------------------------
	fruit1 := cart[1:2]
	fmt.Println(fruit1)

	//---------------------------------
	var s []int
	for i := range 10000 {
		s = appendInt(s, i)
	}
	fmt.Println(s[:10])
	//---------------------------------
	// Ecercise: cncat ,without using a "for" loop
	out := concat([]string{"a", "b"}, []string{"c"})
	fmt.Println(out)

	//---------------------------------
	// Ecercise: Medain
	values:= []float64{1,6,3,4,5}
    value1:= []float64{1,2,3,5}
	fmt.Println(median(values))
	fmt.Println(median(value1))
	fmt.Println(values)

	//---------------------------------

	players:= []player{
		{"lokesh",10000000},
		{"dhoni",40},
	}


	
	// add bonus
	//pointer semantics "or" loop
	for i:= range players{
		players[i].Score+=100
	}
	fmt.Println(players)
}

type player struct{
	Name string
	Score int
}

func median(values []float64) float64 {
	  
    vals := make([]float64, len(values))
	copy(vals, values)
	   sort.Float64s(vals)
	    length:= len(vals)/2
	   if len(vals)%2==1{
		   return vals[length]
	   }

	  mid := (vals[length-1] + vals[length])/2
	  return mid

	}


func concat(a, b []string) []string {
	cont := make([]string, len(a)+len(b))
	copy(cont, a)
	copy(cont[len(a):], b)

	return cont
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) == cap(s) {
		size := 2 * (len(s) + 1)
		fmt.Println(cap(s), "->", size)
		new := make([]int, size)
		copy(new, s)
		s = new[:len(s)]
	}
	s = s[:len(s)+1]
	s[i] = v
	return s
}
