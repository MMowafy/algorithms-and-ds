package queue

import (
	linked_list "data-structure-and-algorithms/linked-list"
	"fmt"
)

type Queue struct {
	List *linked_list.LinkedList
}

func NewQueue() *Queue {
	return &Queue{
		List: linked_list.NewLinkedList(),
	}
}

func (queue *Queue) Enqueue(val int) {
	fmt.Println("Enqueue ", val)
	queue.List.Insert(val)
}

func (queue *Queue) Dequeue() {
	fmt.Println("Dequeue ", queue.List.Head.Value)
	queue.List.DeleteNodeAt(0)
}
