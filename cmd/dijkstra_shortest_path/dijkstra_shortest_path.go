package main

import (
	"os"
	"log"
	"bufio"
	"errors"
	"strings"
	"strconv"

	"algorithms-in-go/structures/graph"
)

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

func main() {
	file, _ := os.Open("cmd/dijkstra_shortest_path/testcase/problem9.8test.txt")
	fileScanner := bufio.NewScanner(file)
	g, err := ReadEdge(fileScanner)
	if err != nil {
		log.Fatal(err)
	}
	g.PrintAdjList()
}
