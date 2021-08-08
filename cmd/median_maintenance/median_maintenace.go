/*
Use two heaps to implement median maintenace.
Input is a streaming int array. Output is all kth medians from 1 to k.
(k is the length of the int arry)

median is (k+1)/2-th / k/2-th if k is odd / even.
MinIntArray contains all bigger numbers and Pop (Extract-Min) is the smallest of these bigger numbers.
MaxIntArray contains all smaller numbers and Pop (Extract-Mat) is the biggest of these smaller numbers.
*/
package main

import (
	"os"
	"strconv"
	"container/heap"
	"fmt"
	"bufio"
)

type MinIntHeap []int
func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxIntHeap []int
func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sum(arr []int) int {
	res := 0
	for _, n := range(arr) {
		res += n
	}
	return res
}

func medianSum(arr []int) int {
	if len(arr) < 1{
		return 0
	}
	// res
	medians := []int{}

	// init two heaps
	minHeap := &MinIntHeap{}
	heap.Init(minHeap)
	maxHeap := &MaxIntHeap{}
	heap.Init(maxHeap)

	if len(arr) <= 2 {
		medians = append(medians, arr[0])
		if len(arr) == 2 {
			medians = append(medians, arr[0])
		}
		return sum(medians)
	}
	for i, n := range(arr) {
		// init
		if i == 0 {
			medians = append(medians, arr[0])
			continue
		} else if i == 1 {
			if arr[0] < arr[1] {
				medians = append(medians, arr[0])
			} else {
				medians = append(medians, arr[1])
			}
			continue
		}
		if i == 2 {
			if arr[0] > arr[1] {
				heap.Push(minHeap, arr[0])
				heap.Push(maxHeap, arr[1])
			} else {
				heap.Push(minHeap, arr[1])
				heap.Push(maxHeap, arr[0])
			}
		}
		// push new n
		minOfBig := (*minHeap)[0]  // get the minimum of all big numbers
		maxOfSmall := (*maxHeap)[0]  // get the maximum of all small numbers
		if n <= minOfBig && n >= maxOfSmall {
			if minHeap.Len() < maxHeap.Len() {
				heap.Push(minHeap, n)
			} else {
				heap.Push(maxHeap, n)
			}
		} else if n >= minOfBig {
			heap.Push(minHeap, n)
		} else {
			heap.Push(maxHeap, n)
		}
		// rebalance two heaps & compute median
		if minHeap.Len() - maxHeap.Len() > 1 {
			x := heap.Pop(minHeap)
			heap.Push(maxHeap, x)
			medians = append(medians, (*minHeap)[0])
		} else if maxHeap.Len() - minHeap.Len() > 1 {
			x := heap.Pop(maxHeap)
			heap.Push(minHeap, x)
			medians = append(medians, (*maxHeap)[0])
		} else {
			if minHeap.Len() > maxHeap.Len() {
				medians = append(medians, (*minHeap)[0])
			} else if minHeap.Len() < maxHeap.Len() {
				medians = append(medians, (*maxHeap)[0])
			} else {
				medians = append(medians, (*maxHeap)[0])
			}
		}
	}
	return sum(medians)
}


func main() {
	// read txt
	arr := []int{}
	file, _ := os.Open("cmd/median_maintenance/testcase/problem11.3test.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		arr = append(arr, num)
	}
	sumOfMedians := medianSum(arr)
	fmt.Printf("%d (from %d)\n", sumOfMedians % 10000, sumOfMedians)

	arr = []int{}
	file, _ = os.Open("cmd/median_maintenance/testcase/problem11.3.txt")
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		arr = append(arr, num)
	}
	sumOfMedians = medianSum(arr)
	fmt.Printf("%d (from %d)\n", sumOfMedians % 10000, sumOfMedians)
}
