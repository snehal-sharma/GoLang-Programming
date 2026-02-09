
package main

import "fmt"

func main() {
	graph := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'1', '0', '0', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	count := NumberOfIslands(graph)
	fmt.Println(count)
}

func NumberOfIslands(graph [][]byte) int {
	rows := len(graph)
	cols := len(graph[0])

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if graph[r][c] == '1' {
				count++
				dfs(r, c, graph)
			}
		}
	}
	return count
}

func dfs(r, c int, graph [][]byte) {
	rows := len(graph)
	cols := len(graph[0])
	if r < 0 || c < 0 || r >= rows || c >= cols || graph[r][c] == '0' {
		return
	}

	graph[r][c] = '0'

	dfs(r+1, c, graph)
	dfs(r-1, c, graph)
	dfs(r, c+1, graph)
	dfs(r, c-1, graph)
}
