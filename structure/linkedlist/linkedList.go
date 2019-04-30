package linkedlist

import (
	"fmt"
	"strings"
)

// Node structure contains key(data) and next node pointer
type Node struct {
	key  string
	next *Node
}

type List interface {
	Reverse() *Node
	Print()
}

// Build a list from array
func Build(data []string) *Node {
	if len(data) == 0 {
		fmt.Println("Data is empty, return nil")
		return nil
	}

	list := &Node{data[0], nil}
	root := list
	for i, e := range data {
		if i > 0 {
			root.next = &Node{e, nil}
			root = root.next
		}
	}

	return list
}

// Reverse whole list
func (list *Node) Reverse() *Node {
	if list == nil {
		fmt.Println("Node is nil!")
	}

	curr := list

	if list.next != nil {
		following := curr.next
		curr.next = nil
		for following != nil {
			nf := following.next // next following
			following.next = curr
			
			curr = following
			following = nf
		}
	}

	return curr
}

// Print key of each node
func (list *Node) Print() {
	node := list
	fmt.Print("{ ")
	for node != nil {
		fmt.Printf("%v ", node.key)
		node = node.next
	}
	fmt.Println("}")
}

// Test is a driver function
func Test() {
	// data := []string{"a", "b", "c", "d", "e"}
	data := strings.Split("kikiyoyo", "")

	list := Build(data)
	list.Print()

	list = list.Reverse()
	list.Print()
}
