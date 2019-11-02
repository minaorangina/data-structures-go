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
	fmt.Println(ll.valueAt(4))

	val, ok := ll.front()
	if !ok {
		fmt.Println("List is empty")
	} else {
		fmt.Printf("front: %d\n", val)
	}
}
