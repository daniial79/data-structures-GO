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

func (ll *LinkedList) PopFirst() *Node {
	if ll.Length == 0 {
		fmt.Println("LinkedList is empty!")
		os.Exit(1)
	}

	var firstNode *Node

	if ll.Length == 1 {
		firstNode = ll.Head

		ll.Head = nil
		ll.Tail = nil
	} else {
		firstNode = ll.Head

		ll.Head = ll.Head.Next
		firstNode.Next = nil
	}

	ll.Length--
	return firstNode
}

func (ll *LinkedList) Get(index int) *Node {
	if index < 0 || index > ll.Length-1 {
		fmt.Println("Index out of range!")
		os.Exit(1)
	}

	var targetNode *Node

	if index == 0 {
		targetNode = ll.Head
	} else if index == ll.Length-1 {
		targetNode = ll.Tail
	} else {
		targetNode = ll.Head

		for i := 0; i < index; i++ {
			targetNode = targetNode.Next
		}
	}

	return targetNode
}

func (ll *LinkedList) SetValue(val, index int) {
	targetNode := ll.Get(index)
	targetNode.Value = val
}

func (ll *LinkedList) Insert(val, index int) {
	if index == 0 {
		ll.Preppend(val)
	} else if index == ll.Length-1 {
		ll.Append(val)
	} else {
		tarNode := ll.Get(index)
		tarNodePrev := ll.Get(index - 1)

		newNode := Node{
			Value: val,
			Next:  nil,
		}

		newNode.Next = tarNode
		tarNodePrev.Next = &newNode

		ll.Length++
	}

}
