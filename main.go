package main

import (
	linked_list "data-structure-and-algorithms/linked-list"
	"data-structure-and-algorithms/queue"
	stack2 "data-structure-and-algorithms/stack"
	"fmt"
)

func main() {
	l := linked_list.NewLinkedList()
	l.Insert(12)
	l.Insert(13)
	l.Insert(14)
	l.Insert(15)
	l.InsertAt(20, 2)
	l.Traverse()
	fmt.Println("************* Print11 *************")
	l.Search(20)
	l.Search(100)
	l.DeleteNode(20)
	l.Traverse()
	l.InsertAt(20, 2)
	fmt.Println("************* Print22 *************")
	l.Traverse()
	l.DeleteNodeAt(0)
	l.DeleteNodeAt(6)
	l.DeleteNodeAt(1)
	fmt.Println("************* Print33 *************")
	l.Traverse()
	fmt.Println("************* Print33 *************")
	fmt.Println("************* Print33 *************")
	fmt.Println("************* Print33 *************")
	fmt.Println("************* Print33 *************")
	fmt.Println("************* Print33 *************")

	queue := queue.NewQueue()
	queue.List = l
	fmt.Println("************* Print Queue *************")
	queue.List.Traverse()
	queue.Enqueue(100)
	fmt.Println("************* Print Queue *************")
	queue.List.Traverse()
	queue.Dequeue()
	queue.List.Traverse()

	l2 := linked_list.NewLinkedList()
	l2.Insert(12)
	l2.Insert(13)
	l2.Insert(14)
	l2.Insert(15)
	l2.InsertAt(20, 2)

	stack := stack2.NewStack()
	stack.List = l2
	fmt.Println("************* Print stack *************")
	stack.List.Traverse()
	stack.Push(111)
	fmt.Println("************* Print stack *************")
	stack.List.Traverse()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	fmt.Println("************* Print stack *************")
	stack.List.Traverse()

}
