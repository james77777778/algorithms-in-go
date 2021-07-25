package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"algorithms-in-go/structures/graph"
	"algorithms-in-go/structures/heap"
)

func min(a, b int) (int) {
	if a > b {
		return b
	}
	return a
}

func ReadEdge(scanner *bufio.Scanner) (g *graph.Graph, err error) {
	g = graph.NewGraph()
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Split(line, "\t")
		u, err := strconv.Atoi(lines[0])
		if err != nil {
			return g, errors.New(err.Error())
		}
		for _, s := range lines[1:] {
			subline := strings.Split(s, ",")
			v, err := strconv.Atoi(subline[0])
			if err != nil {
				return g, errors.New(err.Error())
			}
			weight, err := strconv.Atoi(subline[1])
			if err != nil {
				return g, errors.New(err.Error())
			}
			g.AddEdgeWithWeight(u, v, weight)
		}
	}
	return g, nil
}

func DijkstraShortestPath(g *graph.Graph, startNode int) (map[int]int) {
	// init
	allNodes := g.GetNodes()
	visited := make(map[int]struct{})
	distances := make(map[int]int)
	dijkstraHeap := heap.NewHeap()
	maxDistance := 100000  // a very big number and guarantee > all possible distances
	for _, v := range allNodes {
		distances[v] = maxDistance
	}

	// put start node
	for _, v := range allNodes {
		if v == startNode {
			dijkstraHeap.Insert(0, v)
			continue
		}
		weight, ok := g.Edges[graph.Edge{U: startNode, V: v}]
		if ok {
			dijkstraHeap.Insert(weight, v)
		} else {
			dijkstraHeap.Insert(maxDistance, v)
		}
	}

	for len(allNodes) != len(visited) {
		// Dijkstra's greedy criterion
		var minNode heap.Node
		var err error
		for true {
			minNode, err = dijkstraHeap.ExtractMin()
			if err != nil {
				log.Fatal(err)
			}
			_, ok := visited[minNode.Name]
			if !ok {
				break
			}
		}
		// add w to visited
		w := minNode.Name
		visited[w] = struct{}{}
		distances[w] = minNode.Key

		// update dijkstraHeap
		for _, adj := range g.AdjList[w] {
			_, ok := visited[adj]
			if !ok {
				dijkstraHeap.Delete(adj)
				newKey := min(distances[adj], distances[w] + g.Edges[graph.Edge{U: w, V: adj}])
				distances[adj] = newKey
				dijkstraHeap.Insert(newKey, adj)
			}
		}
	}
	return distances
}

func main() {
	file, _ := os.Open("cmd/dijkstra_shortest_path/testcase/problem9.8test.txt")
	fileScanner := bufio.NewScanner(file)
	g, err := ReadEdge(fileScanner)
	if err != nil {
		log.Fatal(err)
	}
	distances := DijkstraShortestPath(g, 1)
	allNodes := g.GetNodes()
	for _, n := range allNodes {
		fmt.Printf("vertices: %v, shortest path: %v\n", n, distances[n])
	}
}
