package main

import (
	"fmt"
	"log"

	"algorithms-in-go/structures/heap"
)

func main() {
	h := heap.NewHeap()

	h.Insert(3, 1)
	h.Insert(5, 2)
	h.Insert(10, 3)
	h.Insert(1, 4)
	h.Insert(100, 5)
	h.Insert(7, 6)
	h.Insert(4, 7)

	for len := h.Len(); len > 0; len = h.Len() {
		min, err := h.ExtractMin()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(min)
	}

	fmt.Println("New Insert:")
	h.Insert(1, 8)
	h.Insert(10, 9)
	h.Insert(4, 10)

	for len := h.Len(); len > 0; len = h.Len() {
		min, err := h.ExtractMin()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(min)
	}
}