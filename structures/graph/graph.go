package graph

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

type Node struct {
	Id int
}

type Edge struct {
	u, v int
}

type Graph struct {
	Nodes []*Node
	AdjList map[int][]int
	Edges map[Edge]int
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: []*Node{},
		AdjList: map[int][]int{},
		Edges: map[Edge]int{},
	}
}

func (g *Graph) FindNode(id int) (found bool) {
	for _, n := range g.Nodes {
		if id == n.Id {
			return true
		}
	}
	return false
}

func (g *Graph) FindEdgeWeight(u, v int) (weight int, err error) {
	if weight, ok := g.Edges[Edge{u, v}]; ok {
		return weight, nil
	} else {
		return 0, errors.New(fmt.Sprintf("No edge (%v, %v)", u, v))
	}
}

func (g *Graph) AddNode(newId int) (err error) {
	// check no same id
	if g.FindNode(newId) {
		return errors.New(fmt.Sprintf("Can not add %v because there already has same id", newId))
	}
	// add the node
	g.Nodes = append(g.Nodes, &Node{
		Id: newId,
	})
	return nil
}

func (g *Graph) AddEdge(u, v int) (err error) {
	// add nodes
	if !g.FindNode(u) {
		g.Nodes = append(g.Nodes, &Node{
			Id: u,
		})
	}
	if !g.FindNode(v) {
		g.Nodes = append(g.Nodes, &Node{
			Id: v,
		})
	}
	// add the edge
	g.AdjList[u] = append(g.AdjList[u], v)
	g.Edges[Edge{u, v}] = 1

	return nil
}

func (g *Graph) AddEdgeWithWeight(u, v, weight int) (err error) {
	// add nodes
	if !g.FindNode(u) {
		g.Nodes = append(g.Nodes, &Node{
			Id: u,
		})
	}
	if !g.FindNode(v) {
		g.Nodes = append(g.Nodes, &Node{
			Id: v,
		})
	}
	// add the edge
	g.AdjList[u] = append(g.AdjList[u], v)
	g.Edges[Edge{u, v}] = weight

	return nil
}

func (g *Graph) GetNodes() (res []int) {
	res = make([]int, 0)
	for _, n := range g.Nodes {
		res = append(res, n.Id)
	}
	sort.Ints(res)
	return res
}

func (g *Graph) PrintAdjList() {
	// get assending order
	order := g.GetNodes()
	// print adj list by order
	for _, o := range order {
		lst := g.AdjList[o]
		// do not output if there is no node in adj list
		if len(lst) == 0 {
			continue
		}
		// output
		fmt.Printf("%v->", o)
		for i, n := range lst {
			w, err := g.FindEdgeWeight(o, n)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v(%v)", n, w)
			if i != len(lst) - 1 {
				fmt.Printf("->")
			}
		}
		fmt.Printf("\n")
	}
}