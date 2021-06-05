package main

import (
	"fmt"
	"math/rand"
)

func RandomChoosePivot(n int) int {
	randIdx := rand.Intn(n)
	return randIdx
}

func Partition(A *[]int, n, pIdx int) int {
	// move p to the head
	(*A)[pIdx], (*A)[0] = (*A)[0], (*A)[pIdx]

	// init
	i, j := 1, 1

	// partition
	for j < n {
		if (*A)[j] < (*A)[0] {
			(*A)[j], (*A)[i] = (*A)[i], (*A)[j]
			i++
		}
		j++
	}

	// move p to i-1 idx
	(*A)[0], (*A)[i-1] = (*A)[i-1], (*A)[0]

	return i-1
}

func QuickSort(A []int, n int) []int {
	if n < 2 {
		return A
	}

	pIdx := RandomChoosePivot(n)

	medianIdx := Partition(&A, n, pIdx)

	QuickSort(A[:medianIdx], len(A[:medianIdx]))
	QuickSort(A[medianIdx + 1:], len(A[medianIdx + 1:]))

	return A
}

func main() {
	A := []int{3, 8, 2, 5, 1, 4, 7, 6}

	sortedA := QuickSort(A, len(A))

	fmt.Println(sortedA)
}