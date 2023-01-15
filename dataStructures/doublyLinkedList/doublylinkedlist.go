package dll

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
