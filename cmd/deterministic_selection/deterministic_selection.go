package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"os"
	"strconv"
)

func GroupsOfK(A []int, k int) []int {
	var C []int
	i := 0
	for i < len(A) {
		if i + k < len(A) {
			group := make([]int, k)
			copy(group, A[i:i+k])
			sort.Ints(group)
			C = append(C, group[k/2])  // median of group
		} else {
			break
		}
		i++
	}
	if i != len(A) {
		group := make([]int, len(A)-i-1)
		copy(group, A[i:])
		sort.Ints(group)
		C = append(C, group[(len(A)-i-1)/2])
	}
	return C
}

func IndexOf(val int, data []int) (int) {
	for k, v := range data {
		if val == v {
			return k
		}
	}
	return -1  //not found
 }

func Partition(A []int, randIdx int) int {
	// move p to the head
	A[randIdx], A[0] = A[0], A[randIdx]

	// init
	n := len(A)
	i, j := 1, 1

	// partition
	// if A[j] < p then swap A[i] & A[j]
	// i is the position of the border that < p
	// j is the postiion to record the sorted array
	for j < n {
		if A[j] < A[0] {
			A[i], A[j] = A[j], A[i]
			i++
		}
		j++
	}

	// move p to i-1 index to finish partition
	A[0], A[i-1] = A[i-1], A[0]
	return i - 1
}

func MedianOfMedians(C []int, n int) int {
	if n <= 5 {
		sort.Ints(C)
		return C[n/2]
	} else {
		C_ := GroupsOfK(C, 5)
		return MedianOfMedians(C_, n/5)
	}
}

func DSelect(A []int, n, i int) int {
	if n < 2 {
		return A[1]
	}
	// bread A into groups of 5, sort each group, find middle elements
	C := GroupsOfK(A, 5)
	
	// recursively computes median of C
	var p int
	p = MedianOfMedians(C, n/5)
	idx := IndexOf(p, A)

	// partition
	// j is the new index of p
	j := Partition(A, idx)

	// recur
	if j == i {
		return A[j]
	} else if j > i {
		return DSelect(A[:j], j-1, i)
	} else {
		return DSelect(A[j:], n-j, i-j)
	}
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	// testcase 1
	i := 5 - 1
	f, err := os.Open("cmd/deterministic_selection/testcase/problem6.5test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err := ReadInts(f)

	median := DSelect(A, len(A), i)
	fmt.Println(median)

	// testcase 2
	i = 50 - 1
	f, err = os.Open("cmd/deterministic_selection/testcase/problem6.5test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err = ReadInts(f)

	median = DSelect(A, len(A), i)
	fmt.Println(median)
}
