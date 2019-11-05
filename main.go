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

	index = 3
	val, err = ll.valueNFromBack(index)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("value[%d]FromBack: %d (expecting 4)\n", index, val)
	}

	index = 5
	val, err = ll.valueNFromBack(index)
	if err != nil {
		fmt.Println("valueNFromBack:", err)
	} else {
		fmt.Printf("value[%d]FromBack: %d (expecting error)\n", index, val)
	}

	err = ll.pushFront(78)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("pushFront %d\n", 78)
		ll.display()
	}

	err = ll.popFront()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("popFront")
		ll.display()
	}

	err = ll.pushBack(555)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("pushBack %d\n", 555)
		ll.display()
	}

	err = ll.popBack()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("popBack")
		ll.display()
	}

	err = ll.insert(2, 88)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert")
		ll.display()
	}

	err = ll.delete(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("delete @ index %d (expecting 88 to disappear)\n", 2)
		ll.display()
	}

	err = ll.delete(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("delete @ index %d (expecting 9 to disappear)\n", 2)
		ll.display()
	}

	err = ll.reverse()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("reverse(expecting 7, 9, 4)")
		ll.display()
	}

	deleted, err := ll.deleteFirstOccurrence(9)
	if err != nil {
		fmt.Println(err)
	} else if deleted {
		fmt.Println("delete first occurrence(expecting true: 7, 4)")
		ll.display()
	} else {
		fmt.Printf("\tdelete first occurrence did not delete (this was unexpected)\n")
	}

	deleted, err = ll.deleteFirstOccurrence(55)
	if err != nil {
		fmt.Println(err)
	} else if deleted {
		fmt.Printf("\tdelete first occurrence deleted (this was unexpected)\n")
	} else {
		fmt.Printf("delete first occurrence did not find occurrence of %d, as expected\n", 55)
		ll.display()
	}
}
