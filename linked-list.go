package main

import "fmt"

// Node is a node within a linked list
type node struct {
	data int
	next *node
}

type linkedlist struct {
	head *node
}

// NewLinkedList initialises a new linked list
func NewLinkedList() linkedlist {
	return linkedlist{}
}

func (ll *linkedlist) addNode(val int) {
	newNode := &node{data: val}
	if ll.head == nil {
		ll.head = newNode
		fmt.Printf("head: %v\n", ll.head)
		return
	}

	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

func (ll *linkedlist) display() {
	if ll.head == nil {
		fmt.Printf("List is empty.\n")
		return
	}

	for currentNode := ll.head; currentNode != nil; currentNode = currentNode.next {
		fmt.Printf("-> %v\n", currentNode.data)
	}
}

func (ll *linkedlist) size() int {
	var size int
	if ll.head != nil {
		currentNode := ll.head
		for currentNode != nil {
			size++
			currentNode = currentNode.next
		}
	}
	return size
}

func (ll *linkedlist) isEmpty() {
	fmt.Println(ll.head == nil)
}
