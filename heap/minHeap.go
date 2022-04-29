package heap

import "fmt"

type MinHeap struct {
	arr     []int
	size    int
	maxSize int
}

func NewMinHeap(maxSize int) *MinHeap {
	return &MinHeap{
		arr:     []int{},
		size:    0,
		maxSize: maxSize,
	}
}

func (m *MinHeap) Parent(index int) int {
	return (index - 1) / 2
}

func (m *MinHeap) LeftChild(index int) int {
	return 2*index + 1
}

func (m *MinHeap) RightChild(index int) int {
	return 2*index + 2
}

func (m *MinHeap) Insert(val int) {
	if m.size >= m.maxSize {
		fmt.Println("We should return error ")
	}
	m.arr = append(m.arr, val)
	m.size++
	m.upHeapify(m.size - 1)
}

func (m *MinHeap) upHeapify(index int) {
	for m.arr[index] < m.arr[m.Parent(index)] {
		m.swap(index, m.Parent(index))
		index = m.Parent(index)
	}
}

func (m *MinHeap) swap(index1 int, index2 int) {
	temp := m.arr[index1]
	m.arr[index1] = m.arr[index2]
	m.arr[index2] = temp
}

func (m *MinHeap) leaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *MinHeap) Remove() int {
	removedItem := m.arr[0]
	m.arr[0] = m.arr[m.size-1]
	m.arr = m.arr[:m.size-1]
	m.size--
	m.downHeapify(0)
	return removedItem
}

func (m *MinHeap) downHeapify(index int) {
	if m.leaf(index) {
		return
	}

	smallest := index
	leftChildIndex := m.LeftChild(index)
	rightRightIndex := m.RightChild(index)
	if leftChildIndex < m.size && m.arr[leftChildIndex] < m.arr[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < m.size && m.arr[rightRightIndex] < m.arr[smallest] {
		smallest = rightRightIndex
	}
	if smallest != index {
		m.swap(index, smallest)
		m.downHeapify(smallest)
	}
	return
}

func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	minHeap := NewMinHeap(len(inputArray))
	for _, val := range inputArray {
		minHeap.Insert(val)
	}
	fmt.Println("after insert ", minHeap.arr)
	for i := 0; i < len(inputArray); i++ {
		fmt.Println("remove", minHeap.Remove())
	}
}
