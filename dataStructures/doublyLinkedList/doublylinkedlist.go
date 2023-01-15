package dll

import (
	"fmt"
	"os"
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

func (dll *DoublyLinkedList) Pop() *Node {
	if dll.Length == 0 {
		fmt.Println("List is empty!")
		os.Exit(1)
	}

	var node *Node

	if dll.Length == 1 {
		node = dll.Head
		dll.Head = nil
		dll.Tail = nil
	} else {
		node = dll.Tail
		dll.Tail = node.Prev

		dll.Tail.Next = nil
		node.Prev = nil
	}

	dll.Length--
	return node
}

func (dll *DoublyLinkedList) Preppend(val int) {
	node := Node{Value: val}

	if dll.Length == 0 {
		dll.Head = &node
		dll.Tail = &node
	} else {
		node.Next = dll.Head
		dll.Head.Prev = &node

		dll.Head = &node
	}

	dll.Length++
}

func (dll *DoublyLinkedList) PopFirst() *Node {
	if dll.Length == 0 {
		fmt.Println("List is empty")
		os.Exit(1)
	}

	var node *Node

	if dll.Length == 1 {
		node = dll.Head
		dll.Head = nil
		dll.Tail = nil
	} else {
		node = dll.Head
		dll.Head = dll.Head.Next
		node.Next = nil
		dll.Head.Prev = nil
	}

	dll.Length--
	return node
}

func (dll *DoublyLinkedList) Get(index int) *Node {
	if index < 0 || index > dll.Length-1 {
		fmt.Println("Index out of range!")
		os.Exit(1)
	} else if dll.Length == 0 {
		fmt.Println("List is empty!")
		os.Exit(1)
	}

	var node *Node

	if index <= dll.Length/2 {
		temp := dll.Head
		for i := 1; i <= index; i++ {
			temp = temp.Next
		}
		node = temp
	} else {
		temp := dll.Tail
		for i := dll.Length - 1; index < i; i-- {
			temp = temp.Prev
		}
		node = temp
	}

	return node
}

func (dll *DoublyLinkedList) SetValue(index, val int) {
	targetNode := dll.Get(index)
	targetNode.Value = val
}

func (dll *DoublyLinkedList) Insert(index, val int) {
	if index == 0 {
		dll.Preppend(val)
	} else if index == dll.Length-1 {
		dll.Append(val)
	} else {
		targetIndex := dll.Get(index)

		node := Node{Value: val}

		node.Next = targetIndex
		node.Prev = targetIndex.Prev

		targetIndex.Prev.Next = &node
		targetIndex.Prev = &node

		dll.Length++
	}

}

func (dll *DoublyLinkedList) Remove(index int) *Node {
	var targetNode *Node

	if index == 0 {
		dll.PopFirst()
	} else if index == dll.Length-1 {
		dll.Pop()
	} else {
		targetNode = dll.Get(index)

		targetNode.Prev.Next = targetNode.Next
		targetNode.Next.Prev = targetNode.Prev

		targetNode.Next = nil
		targetNode.Prev = nil

	}

	dll.Length--
	return targetNode
}
