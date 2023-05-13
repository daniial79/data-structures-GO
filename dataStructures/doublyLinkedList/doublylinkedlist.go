package dll

import (
	"errors"
	"fmt"
)

type node struct {
	Value int
	Next  *node
	Prev  *node
}

type DoublyLinkedList struct {
	Head   *node
	Tail   *node
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
	node := node{Value: val}
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

func (dll *DoublyLinkedList) Pop() (node *node, err error) {
	if dll.Length == 0 {
		err = errors.New("list is empty")
		return nil, err
	}

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
	return node, nil
}

func (dll *DoublyLinkedList) Preppend(val int) {
	node := node{Value: val}

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

func (dll *DoublyLinkedList) PopFirst() (node *node, err error) {
	if dll.Length == 0 {
		err = errors.New("list is empty")
		return nil, err
	}

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
	return node, nil
}

func (dll *DoublyLinkedList) Get(index int) (node *node, err error) {
	if index < 0 || index > dll.Length-1 {
		err = errors.New("index out of range")
		return nil, err
	}

	if dll.Length == 0 {
		err = errors.New("list is empty")
		return nil, err
	}

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

	return node, nil
}

func (dll *DoublyLinkedList) SetValue(index, val int) (err error) {
	targetNode, err := dll.Get(index)
	if err != nil {
		return err
	}
	targetNode.Value = val
	return nil
}

func (dll *DoublyLinkedList) Insert(index, val int) (err error) {
	if index == 0 {
		dll.Preppend(val)
	} else if index == dll.Length-1 {
		dll.Append(val)
	} else {
		targetIndex, err := dll.Get(index)

		if err != nil {
			return err
		}

		node := node{Value: val}

		node.Next = targetIndex
		node.Prev = targetIndex.Prev

		targetIndex.Prev.Next = &node
		targetIndex.Prev = &node

		dll.Length++
	}

	return nil

}

func (dll *DoublyLinkedList) Remove(index int) (targetNode *node, err error) {
	if index == 0 {
		targetNode, err = dll.PopFirst()
		if err != nil {
			return nil, err
		}
	} else if index == dll.Length-1 {
		targetNode, err = dll.Pop()
		if err != nil {
			return nil, err
		}
	} else {
		targetNode, _ = dll.Get(index)

		targetNode.Prev.Next = targetNode.Next
		targetNode.Next.Prev = targetNode.Prev

		targetNode.Next = nil
		targetNode.Prev = nil

	}

	dll.Length--
	return targetNode, nil
}
