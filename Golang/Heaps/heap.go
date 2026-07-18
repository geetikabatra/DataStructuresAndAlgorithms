package main

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(item int) {
	h.array = append(h.array, item)
	h.heapifyUp(len(h.array) - 1)

}

func (h *MaxHeap) heapifyUp(index int) {
	for h.array[h.parent(index)] < h.array[index] {
		h.swap(h.parent(index), index)
		index = h.parent(index)
	}
}

func (h *MaxHeap) swap(ind1, ind2 int) {
	h.array[ind1], h.array[ind2] = h.array[ind2], h.array[ind1]
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}
	item := h.array[0]
	h.array[0] = 

	return item
}

func (h *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) left(index int) int {
	return (2*index + 1)
}

func (h *MaxHeap) right(index int) int {
	return (2*index + 2)
}

func main() {

}
