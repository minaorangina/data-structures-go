package linkedlist

import "fmt"

// Node is a node within a linked list
type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

// NewLinkedList initialises a new linked list
func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (ll *LinkedList) AddNode(val int) {
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

func (ll *LinkedList) Display() {
	if ll.head == nil {
		fmt.Printf("List is empty.\n")
		return
	}

	for currentNode := ll.head; currentNode != nil; currentNode = currentNode.next {
		fmt.Printf("-> %v\n", currentNode.data)
	}
}

func (ll *LinkedList) Size() int {
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

func (ll *LinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) ValueAt(location int) (int, error) {
	if location < 0 || location > ll.Size() {
		return 0, fmt.Errorf("Index out of range")
	}
	currentNode := ll.head
	for i := 0; i < location; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data, nil
}

func (ll *LinkedList) Front() (int, error) {
	if ll.IsEmpty() {
		return 0, fmt.Errorf("List is empty")
	}
	return ll.head.data, nil
}

func (ll *LinkedList) Back() (int, error) {
	if ll.IsEmpty() {
		return 0, fmt.Errorf("List is empty")
	}
	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode.data, nil
}

func (ll *LinkedList) ValueNFromBack(indexFromBack int) (int, error) {
	if ll.IsEmpty() {
		return 0, fmt.Errorf("List is empty")
	}
	indexFromFront := ll.Size() - 1 - indexFromBack
	val, err := ll.ValueAt(indexFromFront)
	return val, err
}

func (ll *LinkedList) PushFront(value int) error {
	if ll.IsEmpty() {
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

func (ll *LinkedList) PopFront() error {
	if ll.IsEmpty() {
		return fmt.Errorf("List is empty")
	}
	// point head to old head next
	ll.head = ll.head.next
	return nil
}

func (ll *LinkedList) PushBack(value int) error {
	if ll.IsEmpty() {
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

func (ll *LinkedList) PopBack() error {
	if ll.IsEmpty() {
		return fmt.Errorf("List is empty")
	}
	if ll.Size() == 1 {
		return fmt.Errorf("Cannot pop back on a list with only one node")
	}

	currentNode := ll.head
	for currentNode.next != nil && currentNode.next.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = nil
	return nil
}

func (ll *LinkedList) Insert(index, value int) error {
	if ll.IsEmpty() {
		return fmt.Errorf("List is empty")
	}
	if index == 0 {
		return ll.PushFront(value)
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

func (ll *LinkedList) Delete(index int) error {
	if ll.IsEmpty() {
		return fmt.Errorf("List is empty")
	}
	if index == 0 {
		err := ll.PopFront()
		if err != nil {
			return err
		}
	}
	currentNode := ll.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
	return nil
}

func (ll *LinkedList) Reverse() error {
	if ll.IsEmpty() {
		return fmt.Errorf("List is empty")
	}
	if ll.Size() == 1 {
		return nil
	}
	cursor := ll.head
	currentNode := ll.head
	nextNode := ll.head.next

	for nextNode.next != nil {
		currentNode = nextNode
		nextNode = nextNode.next
		currentNode.next = cursor
		cursor = currentNode
	}
	nextNode.next = currentNode
	ll.head.next = nil
	ll.head = nextNode

	return nil
}

func (ll *LinkedList) DeleteFirstOccurrence(value int) (bool, error) {
	if ll.IsEmpty() {
		return false, fmt.Errorf("List is empty")
	}
	if ll.head.data == value {
		err := ll.PopFront()
		if err != nil {
			return false, err
		}
	}

	currentNode := ll.head
	for currentNode.next != nil {
		if currentNode.next.data == value {
			currentNode.next = currentNode.next.next
			return true, nil
		}
		currentNode = currentNode.next
	}
	return false, nil
}
