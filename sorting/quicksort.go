package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"bufio"
	"strconv"
	"math/rand"
	"time"
)

var Count int

const (
	first = iota
	last = iota
	medianOf3 = iota
	random = iota
)

func RandomChoosePivot(n int) int {
	randIdx := rand.Intn(n)
	return randIdx
}

func Partition(A *[]int, n, randomIdx int) int {
	// move p to the head
	(*A)[randomIdx], (*A)[0] = (*A)[0], (*A)[randomIdx]

	// init
	i, j := 1, 1

	// partition
	for j < n {
		if (*A)[j] < (*A)[0] {
			(*A)[j], (*A)[i] = (*A)[i], (*A)[j]
			i++
		}
		j++

		// count comparison
		Count++
	}

	// move p to i-1 idx
	(*A)[0], (*A)[i-1] = (*A)[i-1], (*A)[0]

	return i-1
}

func QuickSort(A []int, n int, mode int) []int {
	if n < 2 {
		return A
	}

	idx := 0
	if mode == first {
		idx = 0
	} else if mode == last {
		idx = n - 1
	} else if mode == medianOf3 {
		medianIdx := MedianOf3(A[0], A[n - 1], A[n / 2])
		if medianIdx == 0 {
			idx = 0
		} else if medianIdx == 1 {
			idx = n - 1
		} else {
			idx = n / 2
		}
	} else {
		idx = RandomChoosePivot(n)
	}

	pivotIdx := Partition(&A, n, idx)

	QuickSort(A[:pivotIdx], len(A[:pivotIdx]), mode)
	QuickSort(A[pivotIdx + 1:], len(A[pivotIdx + 1:]), mode)

	return A
}

func MedianOf3(a, b, c int) int {
	if (a > b) != (a > c) {
		return 0
	} else if (b < a) != (b < c) {
		return 1
	} else {
		return 2
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
	rand.Seed(time.Now().UTC().UnixNano())
	// testcase 1
	fmt.Println("Testcase 1:")
	f, err := os.Open("sorting/testcase/problem5.6test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err := ReadInts(f)

	Count = 0
	QuickSort(A, len(A), first)
	fmt.Printf("%v comparisons\n", Count)

	Count = 0
	f, err = os.Open("sorting/testcase/problem5.6test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err = ReadInts(f)
	QuickSort(A, len(A), last)
	fmt.Printf("%v comparisons\n", Count)

	Count = 0
	f, err = os.Open("sorting/testcase/problem5.6test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err = ReadInts(f)
	QuickSort(A, len(A), medianOf3)
	fmt.Printf("%v comparisons\n", Count)

	Count = 0
	f, err = os.Open("sorting/testcase/problem5.6test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err = ReadInts(f)
	QuickSort(A, len(A), random)
	fmt.Printf("%v comparisons\n", Count)

	// testcase 2
	fmt.Println("\nTestcase 2:")
	f, err = os.Open("sorting/testcase/problem5.6test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err = ReadInts(f)

	Count = 0
	QuickSort(A, len(A), first)
	fmt.Printf("%v comparisons\n", Count)

	Count = 0
	f, err = os.Open("sorting/testcase/problem5.6test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err = ReadInts(f)
	QuickSort(A, len(A), last)
	fmt.Printf("%v comparisons\n", Count)

	Count = 0
	f, err = os.Open("sorting/testcase/problem5.6test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err = ReadInts(f)
	QuickSort(A, len(A), medianOf3)
	fmt.Printf("%v comparisons\n", Count)

	Count = 0
	f, err = os.Open("sorting/testcase/problem5.6test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	A, err = ReadInts(f)
	QuickSort(A, len(A), random)
	fmt.Printf("%v comparisons\n", Count)
}