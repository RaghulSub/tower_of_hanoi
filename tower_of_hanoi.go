package main

import (
	"fmt"

	"github.com/RaghulSub/tower_of_hanoi/stack"
)


func main() {

	s := stack.Stack{}
	s.Push(10)
	s.Push(20)
	// s.Pop()
	top,err := s.Top()
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Top:",top)
	}
	s.Pop()
	
}
