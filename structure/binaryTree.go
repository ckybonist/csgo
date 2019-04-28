package structure

import "fmt"

type Tree interface {
	Grow(data int)
	Delete(key int) *Node
	dfs(node *Node, cb NodeCallback)
	Print()
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

type NodeCallback func(*Node)

func NewNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (root *Node) Grow(data int) {
	newNode := NewNode(data)

	var insert func(*Node)
	insert = func(node *Node) {
		atRight := data > node.data

		if atRight {
			if node.right == nil {
				node.right = newNode
				return
			}
			insert(node.right)
		} else {
			if node.left == nil {
				node.left = newNode
				return
			}
			insert(node.left)
		}
	}

	insert(root)
}

func (root *Node) dfs(cb NodeCallback) {
	var recursive func(*Node)
	recursive = func(node *Node) {
		if node == nil {
			return
		}

		recursive(node.left)
		cb(node)
		recursive(node.right)
	}

	var iterated func()
	iterated = func() {
		stack := []*Node{}
		curr := root

		for curr != nil || len(stack) != 0 {
			for curr != nil {
				if len(stack) == 0 {
					stack = append(stack, curr)
				} else {
					// Prepend()
					stack = append([]*Node{curr}, stack...)
				}
				curr = curr.left
			}
			// for _, e := range stack {
			// 	fmt.Println(e.data)
			// }
			// fmt.Println()

			// Pop()
			curr = stack[0]
			stack = stack[1:]
			cb(curr)

			curr = curr.right
		}
	}

	// recursive(root)
	iterated()
}

func (root *Node) Print() {
	print := func(node *Node) {
		fmt.Println(node.data)
	}

	root.dfs(print)
}

func TestBinaryTree() {
	tree := NewNode(5)
	tree.Grow(4)
	tree.Grow(1)
	tree.Grow(2)
	tree.Grow(7)

	tree.Print()
}
