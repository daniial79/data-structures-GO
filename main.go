package main

import (
	que "github.com/daniial79/data-structures-GO/dataStructures/queue"
)

func main() {

	mq := que.GenQueue()

	for i := 0; i <= 9; i++ {
		mq.Enqueue(i)
	}
	mq.PrintQueue()

}
