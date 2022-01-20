package stack

import (
	linked_list "data-structure-and-algorithms/linked-list"
	"fmt"
)

type Stack struct {
	List *linked_list.LinkedList
}

func NewStack() *Stack {
	return &Stack{
		List: linked_list.NewLinkedList(),
	}
}

func (stack *Stack) Push(val int) {
	fmt.Println("Push ", val)
	stack.List.InsertAt(val, 0)
}

func (stack *Stack) Pop() {
	fmt.Println("Pop ", stack.List.Head.Value)
	stack.List.DeleteNodeAt(0)
}
