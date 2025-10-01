package main

import (
	"fmt"
	"slices"
	
)

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


	i.Move(35,45)
	fmt.Println(i)

	p1:= player{
		Name: "lokesh",
	}

	fmt.Printf("p1: %+v\n",p1)

	fmt.Println(p1.Found(Gold))
	fmt.Println(p1.Found(Copper))
	fmt.Printf("p1: %+ v\n",)


	//----- interfaces
	ms:= []Mover{ &i,&p1}

  moveAll(ms,123,1230)
  for _,m:= range ms{
	fmt.Println(m)
  }


   fmt.Println(Key.String(Copper))
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

// Move moves i by delta x & delta y
// "i" is called "the receiver"
// i is apointer receiver
//value vs pointer receiver
// in general use value  semantics
//try to keep same semantics on all methods
//when yo must use pointer reveiver
 //if you have lock field
 //if you need to mutate the struct
 //decoding /unmarsally

func (i *Item) Move(dx, dy int){
	i.X+=dx
	i.Y+=dy
}


type player struct{
	Name string
	keys []Key
	Item
}

func (p *player) Found(key Key) error{
      switch key {
	  case Copper, Silver, Gold:
		//ok
	  default:
		return fmt.Errorf("invalid key %q",key)
	  }
	  if !slices.Contains(p.keys,key){
		  p.keys=append(p.keys,key)
	  }
	  return nil
}


/*
- Set of methods (and types)
- we define interface as "what you need ", not "what you provide"
- interface are small (stdlib avg ~2)
-if you have interface with more that 4 method ,think again

- start with concrete types,discover interfaces later
*/
type Mover interface{
     Move(int,int)
}
func moveAll(ms []Mover, dx,dy int){
   for _,m:= range ms{
	m.Move(dx,dy)
   }
}

// func Sort(s sortable){
	
// }

// type sortable interface{
// 	Sort(num []int)

//string implement the fmt.Stringer interface
func (k Key) String() string{
	 switch k{
	 case  Copper:
		return "Copper"
	 case Silver:
		return "Silver"
	 case Gold:
		return "Gold"
	 }
    return fmt.Sprintf("<Key %v>",k)
}

type Key byte

const(
	Copper  Key= iota + 1
	Silver
	Gold
)


// rule of thumb: Accept interface, return types 