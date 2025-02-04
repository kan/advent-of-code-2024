package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRoute(board [][]int, x, y, p int) int {
	if board[y][x] != p {
		return 0
	}
	fmt.Printf("x: %d, y: %d, p: %d\n", x, y, p)
	if p == 9 {
		return 1
	}
	r := 0
	if y > 0 {
		// up
		r += getRoute(board, x, y-1, p+1)
	}
	if y < len(board)-1 {
		// down
		r += getRoute(board, x, y+1, p+1)
	}
	if x > 0 {
		// left
		r += getRoute(board, x-1, y, p+1)
	}
	if x < len(board[y])-1 {
		// right
		r += getRoute(board, x+1, y, p+1)
	}
	return r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	board := [][]int{}
	for sc.Scan() {
		line := sc.Text()
		cols := strings.Split(line, "")
		b := []int{}

		for _, col := range cols {
			if col == "." {
				b = append(b, -1)
			} else {
				n, _ := strconv.Atoi(col)
				b = append(b, n)
			}
		}
		board = append(board, b)
	}

	for _, b := range board {
		fmt.Printf("%v\n", b)
	}

	ans := 0

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == 0 {
				r := getRoute(board, x, y, 0)
				fmt.Printf("x: %d, y: %d, r: %d\n", x, y, r)
				ans += r
			}
		}
	}

	fmt.Printf("Answer: %d\n", ans)
}
