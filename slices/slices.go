package main

import "fmt"

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

}
