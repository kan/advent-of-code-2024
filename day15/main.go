package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getRoboPos(grid [][]string) (int, int) {
	for y, row := range grid {
		for x, col := range row {
			if col == "@" {
				return x, y
			}
		}
	}
	return -1, -1
}

func moveUp(grid [][]string) {
	x, y := getRoboPos(grid)
	if y == 0 {
		// 最上部
		return
	}
	box := 0
	for i := 1; y-i >= 0; i++ {
		if grid[y-i][x] == "#" {
			// 壁なので移動できない
			return
		} else if grid[y-i][x] == "O" {
			// 箱は押す
			box = 1
		} else if grid[y-i][x] == "." {
			// 隙間とロボの間に箱があれば押す
			grid[y][x] = "."
			if box == 1 {
				grid[y-i][x] = "O"
			}
			grid[y-1][x] = "@"
			return
		}
	}
}

func moveDown(grid [][]string) {
	x, y := getRoboPos(grid)
	if y == len(grid)-1 {
		// 最下部
		return
	}
	box := 0
	for i := 1; y+i < len(grid); i++ {
		if grid[y+i][x] == "#" {
			// 壁なので移動できない
			return
		} else if grid[y+i][x] == "O" {
			// 箱は押す
			box = 1
		} else if grid[y+i][x] == "." {
			// 隙間とロボの間に箱があれば押す
			grid[y][x] = "."
			if box == 1 {
				grid[y+i][x] = "O"
			}
			grid[y+1][x] = "@"
			return
		}
	}
}

func moveLeft(grid [][]string) {
	x, y := getRoboPos(grid)
	if x == 0 {
		// 最左部
		return
	}
	box := 0
	for i := 1; x-i >= 0; i++ {
		if grid[y][x-i] == "#" {
			// 壁なので移動できない
			return
		} else if grid[y][x-i] == "O" {
			// 箱は押す
			box = 1
		} else if grid[y][x-i] == "." {
			// 隙間とロボの間に箱があれば押す
			grid[y][x] = "."
			if box == 1 {
				grid[y][x-i] = "O"
			}
			grid[y][x-1] = "@"
			return
		}
	}
}

func moveRight(grid [][]string) {
	x, y := getRoboPos(grid)
	if x == len(grid[0])-1 {
		// 最右部
		return
	}
	box := 0
	for i := 1; x+i < len(grid[0]); i++ {
		if grid[y][x+i] == "#" {
			// 壁なので移動できない
			return
		} else if grid[y][x+i] == "O" {
			// 箱は押す
			box = 1
		} else if grid[y][x+i] == "." {
			// 隙間とロボの間に箱があれば押す
			grid[y][x] = "."
			if box == 1 {
				grid[y][x+i] = "O"
			}
			grid[y][x+1] = "@"
			return
		}
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	grid := [][]string{}
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			break
		}
		grid = append(grid, strings.Split(line, ""))
	}
	moveCode := ""
	for sc.Scan() {
		moveCode += sc.Text()
	}

	for _, b := range grid {
		fmt.Printf("%v\n", b)
	}

	for _, c := range moveCode {
		switch c {
		case '^':
			moveUp(grid)
		case 'v':
			moveDown(grid)
		case '<':
			moveLeft(grid)
		case '>':
			moveRight(grid)
		}
	}

	fmt.Println()
	for _, b := range grid {
		fmt.Printf("%v\n", b)
	}

	gps := 0
	for y, row := range grid {
		for x, col := range row {
			if col == "O" {
				gps += y*100 + x
			}
		}
	}

	fmt.Printf("Answer: %d\n", gps)
}
