package main

import (
	"fmt"
	"time"
)

func main(){
	go fmt.Println("Hello, World!")
	fmt.Printf("main")

	for i := range 3{
		go func ()  {
			fmt.Println("goroutine:", i)
		}()
		
	}
	time.Sleep(10*time.Millisecond)

	ch := make(chan int)
	go func ()  {
		ch <-7 
	}()
	
	v := <- ch
	fmt.Println(v)


fmt.Println(sleepSort([]int {20,30,10}))
	
}


func sleepSort(values []int) []int {
    ch := make(chan int)
	for _ , n := range values {
		go func(){
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n

		}()
	}
	var out []int
	for range values {
		n := <- ch
		out = append(out, n)

}
	return out
}