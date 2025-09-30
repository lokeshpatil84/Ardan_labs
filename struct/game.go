package main

import "fmt"

func main() {
	var i Item
	fmt.Println("i %#v \n",i)
	// Can be in any order ,can be partial
	i= Item{
		X:10,
		Y:20,
	}
    
	fmt.Printf("i %#v \n",i)
	
    
	// use %#v for debugging / logging 
	a,b := 1,"1"
	fmt.Printf("a: %#v b: %#v\n",a,b)

	fmt.Println((NewItem(20 ,40)))
	fmt.Println((NewItem(20 ,4000)))
}

// Types of "new" or factory functions
// 1. func NewItem() Item
// 2. func NewItem() *Item
// 3. func NewItem(x,y int) (Item,error)
// 4. func NewItem(x,y int) (*Item,error)

//value semantics: everyone has their own copy
// pointer semantics: everyone shares the same copy(heap , lock)
func NewItem(x,y int)(*Item,error){
    if x<0 || x>maxx || y<0 || y>maxy{
       return nil,fmt.Errorf("invalid coordinates %d,%d",x,y	)
}
   i := Item{X:x,Y:y}

   // the go compiler does escape analysis and will allocate i on the heap
   return &i,nil
}
const(
	maxx=600
	maxy=400
)
type Item struct {
	X int
	Y int
}