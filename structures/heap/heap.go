// Implement a heap constructed by many Node (int, graph.Edge) objects
// comparisons are based on "Key"
// for Dijkstra Shortest Path Algorithm
// Find() should be improved (now O(n))
package heap

import (
	"errors"
	"fmt"
)


type (
	Node struct {
		Key int
		Name int
	}
	Heap struct {
		Element []Node
		Length int
	}
)

func NewHeap() *Heap {
	return &Heap{[]Node{}, 0}
}

func (h *Heap) Len() int {
	return h.Length
}

func (h *Heap) BubbleUp() {
	// index from 1 ~ length
	idx := h.Length
	for {
		if idx <= 1 {
			break
		}
		parent_idx := idx/2-1
		child_idx := idx-1
		if h.Element[parent_idx].Key > h.Element[child_idx].Key {
			h.Element[parent_idx], h.Element[child_idx] = h.Element[child_idx], h.Element[parent_idx]
		}
		idx /= 2
	}
}

func (h *Heap) BubbleDown(pos int) {
	// index from 1 ~ length
	idx := pos
	for {
		if idx > h.Length {
			break
		}
		parent_idx := idx - 1
		child_1_idx := idx * 2 - 1
		child_2_idx := idx * 2
		// two parents exist
		if child_1_idx < h.Length && child_2_idx < h.Length {
			if h.Element[child_1_idx].Key < h.Element[child_2_idx].Key {
				if h.Element[parent_idx].Key > h.Element[child_1_idx].Key {
					h.Element[parent_idx], h.Element[child_1_idx] = h.Element[child_1_idx], h.Element[parent_idx]
					idx = child_1_idx + 1
				} else {
					break
				}
			} else {
				if h.Element[parent_idx].Key > h.Element[child_2_idx].Key {
					h.Element[parent_idx], h.Element[child_2_idx] = h.Element[child_2_idx], h.Element[parent_idx]
					idx = child_2_idx + 1
				} else {
					break
				}
			}
		} else if child_1_idx < h.Length && child_2_idx >= h.Length {  // two parents exist
			if h.Element[parent_idx].Key > h.Element[child_1_idx].Key {
				h.Element[parent_idx], h.Element[child_1_idx] = h.Element[child_1_idx], h.Element[parent_idx]
				idx = child_1_idx + 1
			}  else {
				break
			}
		} else {
			break
		}
	}
}

func (h *Heap) Insert(key int, name int) {
	h.Length += 1
	if len(h.Element) < h.Length {
		h.Element = append(h.Element, Node{Key: key, Name: name})
	} else {
		h.Element[h.Length - 1] = Node{Key: key, Name: name}
	}
	h.BubbleUp()
}

func (h *Heap) ExtractMin() (Node, error) {
	if h.Length == 1 {
		h.Length -= 1
		return h.Element[0], nil
	}
	if h.Length > 1 {
		last_idx := h.Length - 1
		min := h.Element[0]
		h.Element[0], h.Element[last_idx] = h.Element[last_idx], h.Element[0]
		h.Length -= 1
		h.BubbleDown(1)
		return min, nil
	}
	return Node{-1, -1}, errors.New("Heap size < 1, nothing can be extracted")
}

func (h *Heap) Find(name int) (int) {  // O(n) should be improved
	for i, element := range h.Element {
		if element.Name == name {
			return i
		}
	}
	return -1
}

func (h *Heap) Delete(name int) (error) {
	idx := h.Find(name)  // O(n) should be improved
	last_idx := h.Length - 1
	h.Element[idx], h.Element[last_idx] = h.Element[last_idx], h.Element[0]
	h.Length -= 1
	h.BubbleDown(idx+1)
	return nil
}

func (h *Heap) List() {
	for _, node := range h.Element[:h.Length] {
		fmt.Printf("(Key: %v, Name: %v)\n", node.Key, node.Name)
	}
}
