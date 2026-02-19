package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ch1 , ch2 := make(chan string ,1),make(chan string,1)

	ctx ,cancel := context.WithTimeout(context.Background(),15 * time.Millisecond)

	defer cancel()
	go func(){
		time.Sleep(100 * time.Millisecond)
		ch1 <- "one"
	}()
   
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- "two"

	}()

	select{
	case <- ch1:
		fmt.Println("one")
	case <- ch2:
		fmt.Println("two")
	case <- ctx.Done():
		fmt.Println("timeout")		
	}
}