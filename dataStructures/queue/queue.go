package queue

import (
	"fmt"
	"log"
)

type node struct {
	Value int
	Next  *node
}

type Queue struct {
	First  *node
	Last   *node
	Length int
}

func GenQueue() *Queue {
	return &Queue{}
}

func (q *Queue) PrintQueue() {
	if q.Length == 0 {
		log.Fatal("queue is empty")
	}

	temp := q.First
	fmt.Print("first => ")
	for temp.Next != nil {
		fmt.Printf("%d - ", temp.Value)
		temp = temp.Next
	}
	fmt.Printf("%d <= last\n", temp.Value)
}

func (q *Queue) Enqueue(val int) {
	node := node{Value: val}

	if q.Length == 0 {
		q.First = &node
		q.Last = &node
	} else {
		q.Last.Next = &node
		q.Last = &node
	}
	q.Length++

}
