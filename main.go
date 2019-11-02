package main

import "fmt"

func main() {
	ll := NewLinkedList()
	ll.addNode(9)
	ll.addNode(4)
	ll.addNode(2)
	ll.addNode(9)
	ll.addNode(7)
	ll.display()
	fmt.Println(ll.size())
	fmt.Println(ll.isEmpty())
	index := 4
	val, err := ll.valueAt(index)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("valueAt %d: %d\n", index, val)
	}

	val, err = ll.front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("front: %d\n", val)
	}

	val, err = ll.back()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("back: %d\n", val)
	}
}
