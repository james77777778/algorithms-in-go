package heap

import (
	"fmt"
	"errors"
)


type (
	Heap struct {
		data   []int
		length int
	}
)

func NewHeap() *Heap {
	return &Heap{[]int{}, 0}
}

func (h *Heap) Len() int {
	return h.length
}

func (h *Heap) BubbleUp() {
	// index from 1 ~ length
	idx := h.length
	for {
		if idx <= 1 {
			break
		}
		parent_idx := idx/2-1
		child_idx := idx-1
		if h.data[parent_idx] > h.data[child_idx] {
			h.data[parent_idx], h.data[child_idx] = h.data[child_idx], h.data[parent_idx]
		}
		idx /= 2
	}
}

func (h *Heap) BubbleDown() {
	// index from 1 ~ length
	idx := 1
	for {
		if idx > h.length {
			break
		}
		parent_idx := idx - 1
		child_1_idx := idx * 2 - 1
		child_2_idx := idx * 2
		// two parents exist
		if child_1_idx < h.length && child_2_idx < h.length {
			if h.data[child_1_idx] < h.data[child_2_idx] {
				if h.data[parent_idx] > h.data[child_1_idx] {
					h.data[parent_idx], h.data[child_1_idx] = h.data[child_1_idx], h.data[parent_idx]
					idx = child_1_idx + 1
				} else {
					break
				}
			} else {
				if h.data[parent_idx] > h.data[child_2_idx] {
					h.data[parent_idx], h.data[child_2_idx] = h.data[child_2_idx], h.data[parent_idx]
					idx = child_2_idx + 1
				} else {
					break
				}
			}
		} else if child_1_idx < h.length && child_2_idx >= h.length {  // two parents exist
			if h.data[parent_idx] > h.data[child_1_idx] {
				h.data[parent_idx], h.data[child_1_idx] = h.data[child_1_idx], h.data[parent_idx]
				idx = child_1_idx + 1
			}  else {
				break
			}
		} else {
			break
		}
	}
}

func (h *Heap) Insert(key int) {
	h.data = append(h.data, key)
	h.length += 1
	h.BubbleUp()
}

func (h *Heap) ExtractMin() (int, error) {
	if h.length == 1 {
		h.length -= 1
		return h.data[0], nil
	}
	if h.length > 1 {
		last_idx := h.length - 1
		min := h.data[0]
		h.data[0], h.data[last_idx] = h.data[last_idx], h.data[0]
		h.length -= 1
		h.BubbleDown()
		return min, nil
	}
	return -1, errors.New("Heap size < 1, nothing can be extracted")
}

func (h *Heap) List() {
	fmt.Printf("%v\n", h.data[:h.length])
}
