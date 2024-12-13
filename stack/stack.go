package stack

import (
	"errors"
)

type Stack struct{
	arr []int;
}

func (s *Stack) Empty() bool {
	return len(s.arr) == 0
}
func (s *Stack) Push(val int){
	s.arr = append(s.arr,val)
}

func (s *Stack) Pop(){
	if s.Empty() {
		return 
	}
	s.arr = s.arr[:len(s.arr)-1]
}
func (s *Stack) Top() (int,error) {
	if s.Empty() {
		return 0,errors.New("Stack is Empty")	
	}
	return s.arr[len(s.arr)-1],nil
}

