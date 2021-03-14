package linkedlist

import (
	"fmt"
	"testing"

	"github.com/minaorangina/data-structures-go/utils"
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()
	ll.AddNode(9)
	ll.AddNode(4)
	ll.AddNode(2)
	ll.AddNode(9)
	ll.AddNode(7)
	ll.Display()

	fmt.Printf("size? %d\n", ll.Size())
	fmt.Printf("empty? %t\n", ll.IsEmpty())

	index := 4
	val, err := ll.ValueAt(index)
	utils.AssertNoError(t, err)
	fmt.Printf("valueAt %d: %d\n", index, val)

	val, err = ll.Front()
	utils.AssertNoError(t, err)
	fmt.Printf("front: %d\n", val)

	val, err = ll.Back()
	utils.AssertNoError(t, err)
	fmt.Printf("back: %d\n", val)

	index = 3
	val, err = ll.ValueNFromBack(index)
	utils.AssertNoError(t, err)
	fmt.Printf("value[%d]FromBack: %d (expecting 4)\n", index, val)

	index = 5
	val, err = ll.ValueNFromBack(index)
	utils.AssertErrored(t, err)
	fmt.Printf("value[%d]FromBack: %d (expecting error)\n", index, val)

	err = ll.PushFront(78)
	utils.AssertNoError(t, err)
	fmt.Printf("pushFront %d\n", 78)
	ll.Display()

	err = ll.PopFront()
	utils.AssertNoError(t, err)
	fmt.Println("popFront")
	ll.Display()

	err = ll.PushBack(555)
	utils.AssertNoError(t, err)
	fmt.Printf("pushBack %d\n", 555)
	ll.Display()

	err = ll.PopBack()
	utils.AssertNoError(t, err)
	fmt.Println("popBack")
	ll.Display()

	err = ll.Insert(2, 88)
	utils.AssertNoError(t, err)
	fmt.Println("insert")
	ll.Display()

	err = ll.Delete(2)
	utils.AssertNoError(t, err)
	fmt.Printf("delete @ index %d (expecting 88 to disappear)\n", 2)

	err = ll.Delete(0)
	utils.AssertNoError(t, err)
	fmt.Printf("delete @ index %d (expecting 9 to disappear)\n", 2)
	ll.Display()

	err = ll.Reverse()
	utils.AssertNoError(t, err)
	fmt.Println("reverse(expecting 7, 9, 4)")
	ll.Display()

	deleted, err := ll.DeleteFirstOccurrence(9)
	utils.AssertNoError(t, err)
	if deleted {
		fmt.Println("delete first occurrence(expecting true: 7, 4)")
		ll.Display()
	} else {
		fmt.Printf("\tdelete first occurrence did not delete (this was unexpected)\n")
	}

	deleted, err = ll.DeleteFirstOccurrence(55)
	utils.AssertNoError(t, err)
	if deleted {
		fmt.Printf("\tdelete first occurrence deleted (this was unexpected)\n")
	} else {
		fmt.Printf("delete first occurrence did not find occurrence of %d, as expected\n", 55)
		ll.Display()
	}
}
