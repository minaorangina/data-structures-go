package heap

import (
	"testing"

	"github.com/minaorangina/data-structures-go/utils"
)

func TestHeap(t *testing.T) {
	heap := NewHeap()
	utils.AssertNotNil(t, heap)

	heap.add(15)
	utils.AssertEqual(t, len(heap), 1)

	heap.add(17)
	utils.AssertEqual(t, len(heap), 2)
	utils.AssertEqual(t, heap.peek(), 15)

	heap.add(10)
	utils.AssertEqual(t, len(heap), 3)
	utils.AssertEqual(t, heap.peek(), 10)
}
