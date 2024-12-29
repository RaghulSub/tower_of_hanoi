package main

import (
	"fmt"
	"errors"
	"github.com/RaghulSub/tower_of_hanoi/stack"
)

type TowerOfHanoi struct {
	height int
	rods   []stack.Stack
}
func NewTowerOfHanoi(height int) (*TowerOfHanoi,error) {
	toh := TowerOfHanoi{}
	toh.height = height
	toh.rods = make([]stack.Stack, 3)
	for i:=0;i<3;i++ {
		toh.rods[i] = *stack.NewStack(height)
	}
	if height > 0 {
		for i:=height;i>0;i-- {
			toh.rods[0].Push(i)
		}
		return &toh,nil
	}
	return nil,errors.New("height should be greater than 0")
}
func (toh *TowerOfHanoi) MoveTopDisk(source int,destination int) (error) {
	if source == destination {
		return errors.New("invalid move")
	}
	sourceTop,srcErr := toh.rods[source].Top()
	if(srcErr != nil) {
		return errors.New("source rod is empty")
	}else{
		toh.rods[source].Pop()
	}
	if(!toh.rods[destination].Empty()) {
		destinationTop,DetErr := toh.rods[destination].Top()
		if(DetErr != nil) {
			return DetErr
		}
		if sourceTop > destinationTop {
			return errors.New("invalid move")
		}
	}
	toh.rods[destination].Push(sourceTop)

	return nil
}


func (toh *TowerOfHanoi) MoveDisks(n int, source int, destination int, buffer int,path *[][]int) { 
	if n <= 0 {
		return
	}
	if n == 1 {
		toh.MoveTopDisk(source,destination)
		*path = append(*path,[]int{source,destination});
		return
	}
	toh.MoveDisks(n-1,source,buffer,destination,path)
	*path = append(*path,[]int{source,destination});

	err := toh.MoveTopDisk(source,destination)
	if err != nil {
		fmt.Println("There is an error");
		return ;
	}

	toh.MoveDisks(n-1,buffer,destination,source,path)
}
func (toh *TowerOfHanoi) Solve() ([][]int) {
	var path [][]int;
	toh.MoveDisks(toh.height,0,2,1,&path);
	return path
}
func main() {

	toh ,err := NewTowerOfHanoi(3);
	if err != nil {
		fmt.Println("Some Error Occured");
	}
	path := toh.Solve();

	fmt.Println("path:",path);
}
