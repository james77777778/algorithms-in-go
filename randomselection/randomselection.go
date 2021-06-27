package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"bufio"
	"strconv"
	"math/rand"
)


func RandomPivot(A []int) int {
	randIdx := rand.Intn(len(A))
	return randIdx
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
	A[0], A[i - 1] = A[i - 1], A[0]
	return i - 1
}

func RSelect(A []int, n, i int) int {
	if n < 2 {
		return A[1]
	}

	// choose pivot p from A uniformly at random
	randIdx := RandomPivot(A)

	// partition
	// j is the new index of p
	j := Partition(A, randIdx)

	// recur
	if j == i {
		return A[j]
	} else if j > i {
		return RSelect(A[:j], j - 1, i)
	} else {
		return RSelect(A[j:], n - j, i - j)
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
	i := 5 - 1  // 5th of 10 ints
	f, err := os.Open("randomselection/testcase/problem6.5test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err := ReadInts(f)

	median := RSelect(A, len(A), i)
	fmt.Println(median)

	// testcase 2
	i = 50 - 1  // 50th of 100 ints
	f, err = os.Open("randomselection/testcase/problem6.5test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err = ReadInts(f)

	median = RSelect(A, len(A), i)
	fmt.Println(median)
}