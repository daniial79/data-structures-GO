package dataStructure

import "fmt"

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
