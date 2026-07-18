package main

import "fmt"

type Graph struct {
	vertices int
	adjList  map[int][]int
}

func NewGraph(v int) *Graph {

	return &Graph{
		vertices: v,
		adjList:  make(map[int][]int),
	}
}

// addedge
func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

// adjacency list
// type Edge struct {
// 	To     int
// 	Weight int
// }

// type WeightedGraph struct {
// 	vertices int
// 	adjList  map[int][]Edge
// }

// func NewWeightedGraph(v int) *WeightedGraph {
// 	return &WeightedGraph{
// 		vertices: v,
// 		adjList:  make(map[int][]Edge),
// 	}
// }

// func (g *WeightedGraph) AddEdge(u, v, w int) {
// 	g.adjList[u] = append(g.adjList[u], Edge{v, w})
// 	g.adjList[v] = append(g.adjList[v], Edge{u, w})
// }

// func (g *Graph) BFS(start int) []int {
// 	visited := make(map[int]bool)
// 	result := []int{}

// 	safeQueue := []int{start}

// 	visited[start] = true

// 	for len(safeQueue) > 0 {
// 		node := safeQueue[0]
// 		safeQueue = safeQueue[1:]
// 		result = append(result, node)

// 		for _, neigh := range g.adjList[node] {
// 			if !visited[neigh] {
// 				visited[neigh] = true
// 				safeQueue = append(safeQueue, neigh)
// 			}
// 		}
// 	}
// 	return result
// }

// recursive
func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	result := []int{}

	g.DFSHelper(start, visited, &result)

	return result
}

func (g *Graph) DFSHelper(node int, visited map[int]bool, result *[]int) {
	visited[node] = true
	*result = append(*result, node)

	for _, neigh := range g.adjList[node] {

		if !visited[neigh] {
			g.DFSHelper(neigh, visited, result)
		}

	}
}

func main() {
	g := NewGraph(5)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	result := g.DFS(0)

	for _, each := range result {
		fmt.Println(each)
	}

}
