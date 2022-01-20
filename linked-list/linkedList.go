package linked_list

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewNode() *Node {
	return &Node{}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Insert(val int) {
	newNode := &Node{
		Value: val,
	}
	if list.Size == 0 {
		list.Head = newNode
		list.Tail = newNode
		list.Size++
		return
	}
	list.Tail.Next = newNode
	list.Tail = newNode
	list.Size++
}

func (list *LinkedList) InsertAt(val int, pos int) {
	newNode := &Node{
		Value: val,
	}
	if pos == 0 {
		newNode.Next = list.Head
		list.Head = newNode
		list.Size++

	} else if pos >= list.Size {
		list.Tail.Next = newNode
		list.Tail = newNode
		list.Size++
	} else {
		tempNode := list.Head
		for i := 0; i < pos-1; i++ {
			tempNode = tempNode.Next
		}
		newNode.Next = tempNode.Next
		tempNode.Next = newNode
		list.Size++
	}

}

func (list *LinkedList) Traverse() {
	tempNode := list.Head
	for i := 0; i < 100; i++ {
		if tempNode == nil {
			return
		}
		fmt.Println("Traverse Node ", tempNode.Value)
		tempNode = tempNode.Next
	}
}

func (list *LinkedList) Search(val int) {
	tempNode := list.Head
	for i := 0; i < list.Size; i++ {
		if tempNode.Value == val {
			fmt.Println("Node is found with val", tempNode.Value, i)
			return
		}
		tempNode = tempNode.Next
	}
	fmt.Println("Node is not found with val", val)
}

func (list *LinkedList) DeleteList() {
	list.Head = nil
	list.Tail = nil
}

func (list *LinkedList) DeleteNode(val int) {
	tempNode := list.Head
	previousNode := list.Head
	for i := 0; i < list.Size; i++ {
		if tempNode == nil {
			fmt.Println("Node not found")
			return
		}
		if tempNode.Value == val {
			fmt.Println("Node is deleted with val ", tempNode.Value)
			previousNode.Next = tempNode.Next
			list.Size--
			return
		}
		previousNode = tempNode
		tempNode = tempNode.Next
	}
}

func (list *LinkedList) DeleteNodeAt(pos int) {
	if pos == 0 {
		list.Head = list.Head.Next
		list.Size--
		if list.Size == 0 {
			list.Tail = nil
		}
	} else if pos >= list.Size {
		fmt.Println("heeelelel ", pos)
		tempNode := list.Head
		for i := 0; i < list.Size-2; i++ {
			tempNode = tempNode.Next
		}
		if tempNode == list.Head {
			list.Head = nil
			list.Tail = nil
			list.Size--
			return
		}
		fmt.Println(tempNode.Value)
		tempNode.Next = nil
		list.Tail = tempNode
		list.Size--
	} else {
		tempNode := list.Head
		for i := 0; i < pos-1; i++ {
			tempNode = tempNode.Next
		}
		tempNode.Next = tempNode.Next.Next
		list.Size--
	}
}
