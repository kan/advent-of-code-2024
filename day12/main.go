package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dfs(grid [][]string, x, y int, target string, visited [][]bool) (int, int) {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return 0, 1
	}

	if visited[y][x] {
		return 0, 0
	}

	if grid[y][x] != target {
		return 0, 1
	}

	visited[y][x] = true

	count := 1
	fence := 0
	c, f := dfs(grid, x, y-1, target, visited)
	count += c
	fence += f
	c, f = dfs(grid, x, y+1, target, visited)
	count += c
	fence += f
	c, f = dfs(grid, x-1, y, target, visited)
	count += c
	fence += f
	c, f = dfs(grid, x+1, y, target, visited)
	count += c
	fence += f

	return count, fence
}

func getConnect(grid [][]string, x, y int) (int, int) {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return 0, 0
	}

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	target := grid[y][x]

	count, fence := dfs(grid, x, y, target, visited)
	for y, row := range visited {
		for x, col := range row {
			if col {
				grid[y][x] = ""
			}
		}
	}

	return count, fence
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	grid := [][]string{}
	for sc.Scan() {
		line := sc.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	for _, b := range grid {
		fmt.Printf("%v\n", b)
	}

	ans := 0
	for y, row := range grid {
		for x, col := range row {
			if col != "" {
				c, f := getConnect(grid, x, y)
				ans += c * f
				fmt.Printf("%s: %d * %d = %d\n", col, c, f, c*f)
			}
		}
	}

	fmt.Printf("Answer: %d\n", ans)
}
