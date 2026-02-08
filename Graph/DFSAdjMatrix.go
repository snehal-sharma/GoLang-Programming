
package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 1, 1, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	visited := make(map[int]bool, len(matrix))
	dfs(0, matrix, visited)
}

func dfs(node int, matrix [][]int, visited map[int]bool) {

	visited[node] = true

	fmt.Println(node)

	for neighbor := 0; neighbor < len(matrix); neighbor++ {
		if !visited[neighbor] && matrix[node][neighbor] == 1 {
			dfs(neighbor, matrix, visited)
		}
	}
}
