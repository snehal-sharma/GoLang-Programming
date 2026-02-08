
package main

import "fmt"

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {},
		3: {},
	}

	visited := make(map[int]bool)
	dfs(0, graph, visited)
}

func dfs(node int, graph map[int][]int, visited map[int]bool) {
	if visited[node] {
		return
	}

	visited[node] = true

	fmt.Println(node)

	for _, neighbor := range graph[node] {
		dfs(neighbor, graph, visited)
	}
}
