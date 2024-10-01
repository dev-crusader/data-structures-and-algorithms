package datastructure

import "fmt"

type Heap struct {
	List []int
	size int
}

func initializeHeap(s int) *Heap {
	return &Heap{
		List: make([]int, s),
	}
}

func (h *Heap) isFull() bool {
	return h.size == len(h.List)
}

func (h *Heap) Insert(a int) {
	if !h.isFull() {
		h.List[h.size] = a
		h.size++
	} else {
		fmt.Println("Heap Size is full!!")
	}
}

func (h *Heap) getRightChild(index int) int {
	return 2*index + 2
}

func (h *Heap) getLeftChild(index int) int {
	return 2*index + 1
}

func (h *Heap) getParent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) BuildHeap() {
	for i := h.getParent(h.size); i >= 0; i-- {
		h.maxHeapify(i)
	}
}

func (h *Heap) maxHeapify(i int) {
	L := h.getLeftChild(i)
	R := h.getRightChild(i)
	var largest int

	if L < h.size && h.List[L] > h.List[i] {
		largest = L
	} else {
		largest = i
	}

	if R < h.size && h.List[R] > h.List[largest] {
		largest = R
	}

	if i != largest {
		h.List[i], h.List[largest] = h.List[largest], h.List[i]
		h.maxHeapify(largest)
	}
}

// HeapSort function uses maxheap to sort elements
func (h *Heap) HeapSort() []int {
	if h.size == 0 {
		fmt.Println("Heap is empty")
		return []int{}
	}

	result := make([]int, h.size)

	for i := h.size - 1; i >= 0; i-- {
		result[i] = h.List[0]
		h.List[i], h.List[0] = h.List[0], h.List[i]
		h.size--
		h.maxHeapify(0)
	}
	return result
}

func CallHeap() {
	t := initializeHeap(10)
	fmt.Println(t)
	for _, v := range []int{29, 45, 93, 61, 22, 62, 87, 5, 41, 14, 32} {
		t.Insert(v)
	}
	fmt.Println(t.List)
	t.BuildHeap()
	fmt.Println(t.List)
	result := t.HeapSort()
	fmt.Println(result)

}
