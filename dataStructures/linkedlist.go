package dataStructure

import (
	"fmt"
	"os"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func GenLinkedList() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (ll LinkedList) PrintList() {
	if ll.Length == 0 {
		fmt.Println("LinkedList is empty!")
	} else {
		temp := ll.Head

		for temp.Next != nil {
			fmt.Printf("%d => ", temp.Value)
			temp = temp.Next
		}
		fmt.Printf("%d => ", temp.Value)
		fmt.Println()
	}

}

func (ll *LinkedList) Append(val int) {
	node := Node{
		Value: val,
		Next:  nil,
	}
	if ll.Length == 0 {
		ll.Head = &node
		ll.Tail = &node
	} else {
		ll.Tail.Next = &node
		ll.Tail = &node
	}

	ll.Length++
}

func (ll *LinkedList) Pop() *Node {
	if ll.Length == 0 {
		fmt.Println("LinkedList is empty!")
		os.Exit(1)
	}

	var lastNode *Node

	if ll.Length == 1 {
		lastNode = ll.Head

		ll.Head = nil
		ll.Tail = nil

	} else {
		temp := ll.Head
		for temp.Next.Next != nil {
			temp = temp.Next
		}

		lastNode = temp.Next

		ll.Tail = temp
		ll.Tail.Next = nil
	}

	ll.Length--
	return lastNode

}

func (ll *LinkedList) Preppend(val int) {
	node := Node{
		Value: val,
		Next:  nil,
	}

	if ll.Length == 0 {
		ll.Head = &node
		ll.Tail = &node
	} else {
		node.Next = ll.Head
		ll.Head = &node
	}

	ll.Length++
}
