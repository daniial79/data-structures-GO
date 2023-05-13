package linkedList

import (
	"errors"
	"fmt"
)

type node struct {
	Value int
	Next  *node
}

type LinkedList struct {
	Head   *node
	Tail   *node
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
	node := node{
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

func (ll *LinkedList) Pop() (lastNode *node, err error) {
	if ll.Length == 0 {
		err = errors.New("LinkedList is empty!")
		return nil, err
	}

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
	return lastNode, nil

}

func (ll *LinkedList) Preppend(val int) {
	node := node{
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

func (ll *LinkedList) PopFirst() (firstNode *node, err error) {
	if ll.Length == 0 {
		err = errors.New("list is empty")
		return nil, err
	}

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
	return firstNode, nil
}

func (ll *LinkedList) Get(index int) (targetNode *node, err error) {
	if index < 0 || index > ll.Length-1 {
		err = errors.New("Index out of range")
		return nil, err
	}

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

	return targetNode, nil
}

func (ll *LinkedList) SetValue(val, index int) (err error) {
	targetNode, err := ll.Get(index)
	if err != nil {
		return err
	}
	targetNode.Value = val
	return nil
}

func (ll *LinkedList) Insert(val, index int) error {
	if index == 0 {
		ll.Preppend(val)
	} else if index == ll.Length-1 {
		ll.Append(val)
	} else {
		tarNode, err := ll.Get(index)
		if err != nil {
			return err
		}

		tarNodePrev, err := ll.Get(index - 1)
		if err != nil {
			return err
		}

		newNode := node{
			Value: val,
			Next:  nil,
		}

		newNode.Next = tarNode
		tarNodePrev.Next = &newNode

		ll.Length++
	}

	return nil
}

func (ll *LinkedList) Remove(index int) (node *node, err error) {
	if index == 0 {
		node, err = ll.PopFirst()
		if err != nil {
			return nil, err
		}
	} else if index == ll.Length-1 {
		node, err = ll.Pop()
		if err != nil {
			return nil, err
		}
	} else {
		node, err = ll.Get(index)
		if err != nil {
			return nil, err
		}
		prevNode, _ := ll.Get(index - 1)

		prevNode.Next = node.Next

		node.Next = nil
	}

	ll.Length--

	return node, nil

}

func (ll *LinkedList) Reverse() {
	temp := ll.Head
	ll.Head = ll.Tail
	ll.Tail = temp
	after := temp.Next
	var before *node

	for i := 1; i <= ll.Length; i++ {
		after = temp.Next
		temp.Next = before
		before = temp
		temp = after
	}

}
