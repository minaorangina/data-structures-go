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

func (ll *linkedlist) isEmpty() bool {
	return ll.head == nil
}

func (ll *linkedlist) valueAt(location int) (int, error) {
	if location < 0 || location > ll.size() {
		return 0, fmt.Errorf("Index out of range")
	}
	currentNode := ll.head
	for i := 0; i < location; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data, nil
}

func (ll *linkedlist) front() (int, error) {
	if ll.isEmpty() {
		return 0, fmt.Errorf("List is empty")
	}
	return ll.head.data, nil
}

func (ll *linkedlist) back() (int, error) {
	if ll.isEmpty() {
		return 0, fmt.Errorf("List is empty")
	}
	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode.data, nil
}

func (ll *linkedlist) valueNFromBack(indexFromBack int) (int, error) {
	if ll.isEmpty() {
		return 0, fmt.Errorf("List is empty")
	}
	indexFromFront := ll.size() - 1 - indexFromBack
	val, err := ll.valueAt(indexFromFront)
	return val, err
}

func (ll *linkedlist) pushFront(value int) error {
	if ll.isEmpty() {
		return fmt.Errorf("List is empty")
	}
	// create new node
	newNode := &node{data: value}
	// copy head
	oldHead := ll.head
	// point new node to copy
	newNode.next = oldHead
	// point head at new node
	ll.head = newNode
	return nil
}

func (ll *linkedlist) popFront() error {
	if ll.isEmpty() {
		return fmt.Errorf("List is empty")
	}
	// point head to old head next
	ll.head = ll.head.next
	return nil
}

func (ll *linkedlist) pushBack(value int) error {
	if ll.isEmpty() {
		return fmt.Errorf("List is empty")
	}

	newNode := &node{data: value}

	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
	return nil
}

func (ll *linkedlist) popBack() error {
	if ll.isEmpty() {
		return fmt.Errorf("List is empty")
	}
	if ll.size() == 1 {
		return fmt.Errorf("Cannot pop back on a list with only one node")
	}

	currentNode := ll.head
	for currentNode.next != nil && currentNode.next.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = nil
	return nil
}

func (ll *linkedlist) insert(index, value int) error {
	if ll.isEmpty() {
		return fmt.Errorf("List is empty")
	}
	if index == 0 {
		return ll.pushFront(value)
	}

	currentNode := ll.head
	newNode := &node{data: value}

	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
	}
	nextNode := currentNode.next
	currentNode.next = newNode
	newNode.next = nextNode
	return nil
}
