package heap

type Heap []int

func NewHeap() Heap {
	return make(Heap, 0, 10)
}

func (h *Heap) getLeftChildIndex(parentIdx int) int {
	return 2*parentIdx + 1
}

func (h *Heap) getRightChildIndex(parentIdx int) int {
	return 2*parentIdx + 2
}

func (h *Heap) getParentIndex(childIdx int) int {
	return (childIdx - 1) / 2
}

func (h *Heap) hasLeftChild(idx int) bool {
	return h.getLeftChildIndex(idx) < len(*h)
}

func (h *Heap) hasRightChild(idx int) bool {
	return h.getRightChildIndex(idx) < len(*h)
}

func (h *Heap) hasParent(idx int) bool {
	return h.getParentIndex(idx) >= 0
}

func (h *Heap) leftChild(idx int) int {
	childIdx := h.getLeftChildIndex(idx)
	return (*h)[childIdx]
}

func (h *Heap) rightChild(idx int) int {
	childIdx := h.getRightChildIndex(idx)
	return (*h)[childIdx]
}

func (h *Heap) parent(idx int) int {
	return (*h)[h.getParentIndex(idx)]
}

func (h *Heap) swap(a, b int) {
	a, b = b, a
}

func (h *Heap) peek() int {
	if len(*h) == 0 {
		panic("heap is empty")
	}
	return (*h)[0]
}

func (h *Heap) poll() int {
	item := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	h.heapifyDown()
	return item
}

func (h *Heap) heapifyDown() {
	idx := 0
	for h.hasLeftChild(idx) {
		smallerChildIdx := h.getLeftChildIndex(idx)
		if h.hasRightChild(idx) && h.rightChild(idx) < h.leftChild(idx) {
			smallerChildIdx = h.getRightChildIndex(idx)
		}

		if (*h)[idx] < (*h)[smallerChildIdx] {
			break
		}
		h.swap(idx, smallerChildIdx)
		idx = smallerChildIdx
	}
}

func (h *Heap) heapifyUp() {
	idx := len(*h) - 1
	for h.hasParent(idx) && h.parent(idx) > (*h)[idx] {
		h.swap(idx, h.getParentIndex(idx))
		idx = h.getParentIndex(idx)
	}
}

func (h *Heap) add(item int) {
	(*h) = append((*h), item)
	h.heapifyUp()
}
