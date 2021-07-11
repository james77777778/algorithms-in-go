package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"algorithms-in-go/structures/stack"
	"algorithms-in-go/structures/graph"
)

func ReadEdge(scanner *bufio.Scanner, isReversed bool) (g *graph.Graph, err error) {
	g = graph.NewGraph()
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Split(line, " ")
		if len(lines) != 2 {
			return g, errors.New(fmt.Sprintf("Cannot parse: %s", line))
		}
		n1, err := strconv.Atoi(lines[0])
		if err != nil {
			return g, errors.New(err.Error())
		}
		n2, err := strconv.Atoi(lines[1])
		if err != nil {
			return g, errors.New(err.Error())
		}
		if !isReversed {
			g.AddEdge(n1, n2)
		} else {
			g.AddEdge(n2, n1)
		}	
	}
	return g, nil
}

////////// Kosaraju's Two-Pass Algorithm (first pass)
func FirstDFSLoop(g *graph.Graph) (map[int]int){
	t := 0  // finishing time
	explored := map[int]bool{}  				// record the explored nodes
	finishing := map[int]int{}  				// record finishing time {finishing_time: node_id}

	nodes := g.GetNodes()
	// reverse iterate over nodes
	for i := len(nodes)-1; i >= 0; i-- {
		n := nodes[i]
		_, found := explored[n]
		if !found {
			// DFS with stack
			dfsStack := stack.NewStack()
			dfsStack.Push(n)
			explored[n] = true  				// mark n as explored
			for v := dfsStack.Peek(); dfsStack.Len() != 0; v = dfsStack.Peek() {
				noVisit := true
				adj := g.AdjList[v.(int)]
				sort.Ints(adj)  				// always choice smaller node
				for _, j := range adj {
					_, found = explored[j]
					if !found {
						dfsStack.Push(j)
						explored[j] = true
						noVisit = false
						break  					// jump out if there is node explored
					}
				}
				if noVisit {
					dfsStack.Pop()  			// pop the stack if there is no node explored
					t++
					finishing[t] = v.(int) 		// record finishing time
				}
			}
		}
	}
	return finishing
}

////////// Kosaraju's Two-Pass Algorithm (second pass)
func GetFinishingOrder(finishing map[int]int) (res []int) {
	res = make([]int, 0)
	for i := len(finishing); i > 0; i-- {
		res = append(res, finishing[i])
	}
	return res
}

func SecondDFSLoop(g *graph.Graph, finishing map[int]int) (sccs [][]int) {
	sccs = make([][]int, 0)  					// recored all scc
	explored := map[int]bool{}  				// record the explored nodes
	newOrder := GetFinishingOrder(finishing)
	for _, n := range newOrder {
		_, found := explored[n]
		if !found {
			// DFS with stack
			scc := make([]int, 0)  				// init scc
			scc = append(scc, n)  				// add leader node to scc
			dfsStack := stack.NewStack()
			dfsStack.Push(n)
			explored[n] = true  				// mark n as explored
			for v := dfsStack.Peek(); dfsStack.Len() != 0; v = dfsStack.Peek() {
				noVisit := true
				adj := g.AdjList[v.(int)]
				sort.Ints(adj)  				// always choice smaller node
				for _, j := range adj {
					_, found = explored[j]
					if !found {
						dfsStack.Push(j)
						explored[j] = true
						scc = append(scc, j)  	// add node to scc
						noVisit = false
						break  					// jump out if there is node explored
					}
				}
				if noVisit {
					dfsStack.Pop()  			// pop the stack if there is no node explored
				}
			}
			sccs = append(sccs, scc)  			// no more nodes in scc, add scc to sccs
		}
	}
	return sccs
}

func computeSize(sccs [][]int) (sizes []int) {
	sizes = make([]int, 0, 5)
	for _, scc := range sccs {
		sizes = append(sizes, len(scc))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[:5]
}

func main() {
	// case 1
	file, _ := os.Open("cmd/scc/testcase/problem8.10test1.txt")
	fileScanner := bufio.NewScanner(file)
	grev, _ := ReadEdge(fileScanner, true)
	finising := FirstDFSLoop(grev)

	file, _ = os.Open("cmd/scc/testcase/problem8.10test1.txt")
	fileScanner = bufio.NewScanner(file)
	g, _ := ReadEdge(fileScanner, false)
	sccs := SecondDFSLoop(g, finising)

	sccsSizes := computeSize(sccs)
	fmt.Println(sccsSizes)

	// case 2
	file, _ = os.Open("cmd/scc/testcase/problem8.10test2.txt")
	fileScanner = bufio.NewScanner(file)
	grev, _ = ReadEdge(fileScanner, true)
	finising = FirstDFSLoop(grev)

	file, _ = os.Open("cmd/scc/testcase/problem8.10test2.txt")
	fileScanner = bufio.NewScanner(file)
	g, _ = ReadEdge(fileScanner, false)
	sccs = SecondDFSLoop(g, finising)

	sccsSizes = computeSize(sccs)
	fmt.Println(sccsSizes)

	// case 3
	file, _ = os.Open("cmd/scc/testcase/problem8.10test3.txt")
	fileScanner = bufio.NewScanner(file)
	grev, _ = ReadEdge(fileScanner, true)
	finising = FirstDFSLoop(grev)

	file, _ = os.Open("cmd/scc/testcase/problem8.10test3.txt")
	fileScanner = bufio.NewScanner(file)
	g, _ = ReadEdge(fileScanner, false)
	sccs = SecondDFSLoop(g, finising)

	sccsSizes = computeSize(sccs)
	fmt.Println(sccsSizes)

	// case 4
	file, _ = os.Open("cmd/scc/testcase/problem8.10test4.txt")
	fileScanner = bufio.NewScanner(file)
	grev, _ = ReadEdge(fileScanner, true)
	finising = FirstDFSLoop(grev)

	file, _ = os.Open("cmd/scc/testcase/problem8.10test4.txt")
	fileScanner = bufio.NewScanner(file)
	g, _ = ReadEdge(fileScanner, false)
	sccs = SecondDFSLoop(g, finising)

	sccsSizes = computeSize(sccs)
	fmt.Println(sccsSizes)

	// case 5
	file, _ = os.Open("cmd/scc/testcase/problem8.10test5.txt")
	fileScanner = bufio.NewScanner(file)
	grev, _ = ReadEdge(fileScanner, true)
	finising = FirstDFSLoop(grev)

	file, _ = os.Open("cmd/scc/testcase/problem8.10test5.txt")
	fileScanner = bufio.NewScanner(file)
	g, _ = ReadEdge(fileScanner, false)
	sccs = SecondDFSLoop(g, finising)

	sccsSizes = computeSize(sccs)
	fmt.Println(sccsSizes)
}
