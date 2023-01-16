package stack

import (
	"errors"
	"fmt"
)

type node struct {
	Value int
	Rear  *node
}

type Stack struct {
	Top    *node
	Height int
}

func GenStack() *Stack {
	return &Stack{}
}

func (s *Stack) PrintStack() {
	if s.Height == 0 {
		fmt.Println("stack is empty")
	} else {
		temp := s.Top
		for temp.Rear != nil {
			fmt.Printf(" %d\n", temp.Value)
			fmt.Print("||\n")
			fmt.Print("\\/\n")
			temp = temp.Rear
		}
		fmt.Println("NULL")
	}
}

func (s *Stack) Push(val int) {
	node := node{Value: val, Rear: nil}
	if s.Height == 0 {
		s.Top = &node
	} else {
		node.Rear = s.Top
		s.Top = &node
	}
	s.Height++
}

func (s *Stack) Pop() (node *node, err error) {

	if s.Height == 0 {
		err = errors.New("stack is empty!")
		node = nil
	} else {
		err = nil
		node = s.Top
		s.Top = s.Top.Rear
		node.Rear = nil
	}

	return node, err
}
