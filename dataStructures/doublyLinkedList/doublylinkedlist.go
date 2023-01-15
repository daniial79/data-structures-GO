package dll

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func GenDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) PrintList() {
	if dll.Length == 0 {
		fmt.Println("List is empty!")

	} else if dll.Length == 1 {
		fmt.Printf("NULL <=> %d <=> NULL\n", dll.Head.Value)
	} else {
		temp := dll.Head
		fmt.Print("NULL <=> ")
		for temp.Next != nil {
			fmt.Printf("%d <=> ", temp.Value)
			temp = temp.Next
			if temp.Next == nil {
				fmt.Printf("%d <=> NULL\n", temp.Value)
			}
		}
	}
}

func (dll *DoublyLinkedList) Append(val int) {
	node := Node{Value: val}
	if dll.Length == 0 {
		dll.Head = &node
		dll.Tail = &node
	} else {
		node.Prev = dll.Tail
		dll.Tail.Next = &node
		dll.Tail = &node
	}

	dll.Length++
}
