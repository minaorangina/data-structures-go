package binarytree

import (
	"fmt"
	"strconv"
	"strings"
)

type binaryTree struct {
	root *node
}

func NewBinaryTree() *binaryTree {
	return &binaryTree{}
}

func (t *binaryTree) Insert(value int) {
	newNode := &node{value: value}
	if t.root == nil {
		t.root = newNode
		return
	}

	t.root.insert(value)
}

func (t *binaryTree) ContainsDFS(value int) bool {
	if t.root == nil {
		return false
	}

	return t.root.containsDFS(value)
}

func (t *binaryTree) ContainsBFS(value int) bool {
	queue := []node{}
	if t.root == nil {
		return false
	}

	queue = append(queue, *t.root)

	return containsBFS(queue, value)
}

func (t *binaryTree) InOrderTraversal() string {
	var final string
	if t.root == nil {
		return final
	}

	final = t.root.traverseInOrder()
	return strings.Trim(final, " ")
}

type node struct {
	value       int
	left, right *node
}

func (n *node) insert(incomingValue int) {
	if incomingValue < n.value {
		if n.left == nil {
			n.left = &node{value: incomingValue}
			return
		}
		n.left.insert(incomingValue)
	}
	if incomingValue > n.value {
		if n.right == nil {
			n.right = &node{value: incomingValue}
			return
		}
		n.right.insert(incomingValue)
	}
}

func (n *node) containsDFS(value int) bool {
	if value == n.value {
		return true
	}

	if value < n.value {
		if n.left == nil {
			return false
		}
		return n.left.containsDFS(value)
	}

	if n.right == nil {
		return false
	}
	return n.right.containsDFS(value)

}

func containsBFS(queue []node, value int) bool {
	fmt.Printf("QUEUE: %+v\n", queue)
	if len(queue) == 0 {
		return false
	}

	currentNode := queue[0]
	queue = queue[1:]
	if currentNode.value == value {
		return true
	}

	if currentNode.left != nil {
		queue = append(queue, *currentNode.left)
	}

	if currentNode.right != nil {
		queue = append(queue, *currentNode.right)
	}

	return containsBFS(queue, value)
}

func (n *node) traverseInOrder() string {
	var final string

	if n.left != nil {
		final += n.left.traverseInOrder()
		if string(final[len(final)-1]) != " " {
			final += " "
		}
	}

	final += strconv.Itoa(n.value)
	if string(final[len(final)-1]) != " " {
		final += " "
	}

	if n.right != nil {
		final += n.right.traverseInOrder()
		if string(final[len(final)-1]) != " " {
			final += " "
		}
	}

	return final
}
