package binary_tree

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

type BinaryTree struct {
	Root *Node
}

// LevelOrderTraversal /*Also searching same behavior*/
func (tree *BinaryTree) LevelOrderTraversal() {
	if tree.Root == nil {
		return
	}
	var queue []*Node
	queue = append(queue, tree.Root)
	for len(queue) > 0 {
		node := queue[0]
		fmt.Println(node.Val)
		queue = queue[1:]
	}
}

func (tree *BinaryTree) PreOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node)
	tree.PreOrderTraversal(tree.Root.Left)
	tree.PreOrderTraversal(tree.Root.Right)
}

func (tree *BinaryTree) Insert(val int) {
	if tree.Root == nil {
		tree.Root = &Node{Val: val}
		return
	}
	tree.Root.insertNode(val)

}

func (tree *BinaryTree) DeleteTree() {
	tree.Root = nil
}

func (tree *BinaryTree) GetDeepestNode() *Node {
	queue := []*Node{}
	if tree.Root == nil {
		return nil
	}
	queue = append(queue, tree.Root)
	deepestNode := queue[0]
	for len(queue) > 0 {
		deepestNode = queue[0]
		queue = queue[1:]
		if deepestNode.Left != nil {
			queue = append(queue, deepestNode.Left)
		}
		if deepestNode.Right == nil {
			queue = append(queue, deepestNode.Right)
		}
	}
	return deepestNode
}

func (tree *BinaryTree) DeleteDeepestNode() {
	queue := []*Node{}
	if tree.Root == nil {
		return
	}
	queue = append(queue, tree.Root)
	var currentNode *Node
	var previousNode *Node
	for len(queue) > 0 {
		previousNode = currentNode
		currentNode = queue[0]
		queue = queue[1:]
		if currentNode.Left == nil {
			previousNode.Right = nil
			break
		}
		if currentNode.Right == nil {
			currentNode.Left = nil
			break
		}
		queue = append(queue, currentNode.Left)
		queue = append(queue, currentNode.Right)
	}
}

func (node *Node) insertNode(val int) {
	newNode := &Node{
		Val: val,
	}
	if node.Left == nil {
		node.Left = newNode
	} else if node.Right == nil {
		node.Right = newNode
	} else {
		queue := []*Node{}
		queue = append(queue, node)
		currentNode := queue[0]
		for len(queue) > 0 {
			currentNode = queue[0]
			queue = queue[1:]
			if currentNode.Left == nil {
				currentNode.Left = newNode
				break
			}
			if currentNode.Right == nil {
				currentNode.Right = newNode
				break
			}
			queue = append(queue, currentNode.Left)
			queue = append(queue, currentNode.Right)
		}
	}
}

//////////////////////////// Binary Search tree operations //////////////////

func (tree *BinaryTree) InsertInBinarySearchTree(val int) {
	if tree.Root == nil {
		tree.Root = &Node{Val: val}
		return
	}
	tree.Root.insertNodeInBinarySearchTree(tree.Root, val)
}

func (n *Node) insertNodeInBinarySearchTree(node *Node, val int) *Node {
	if node == nil {
		return &Node{
			Val: val,
		}
	}
	if val <= node.Val {
		node.Left = n.insertNodeInBinarySearchTree(node.Left, val)
	} else {
		node.Right = n.insertNodeInBinarySearchTree(node.Right, val)
	}
	return node
}

func (n *Node) deleteNodeInBST(node *Node, val int) *Node {
	if node == nil {
		fmt.Println("Node not found for value", val)
		return nil
	}
	if val < node.Val {
		node.Left = n.deleteNodeInBST(node.Left, val)
	} else if val > node.Val {
		node.Right = n.deleteNodeInBST(node.Right, val)
	} else {
		if node.Left != nil && node.Right != nil {
			temp := node
			minNodeForRight := node // get deepest node for right child
			node.Val = minNodeForRight.Val
			node.Right = n.deleteNodeInBST(temp.Right, minNodeForRight.Val)
		} else if node.Left != nil {
			node = node.Left
		} else if node.Right != nil {
			node = node.Right
		}
	}
	return node
}
