package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func walk(board [][]int, sx, sy, sdirc int) [][]int {
	x, y, dirc := sx, sy, sdirc

	for {
		nx, ny := x, y
		if dirc == 1 {
			ny = y - 1
		} else if dirc == 2 {
			nx = x + 1
		} else if dirc == 3 {
			ny = y + 1
		} else if dirc == 4 {
			nx = x - 1
		}

		if nx < 0 || ny < 0 || nx >= len(board[0]) || ny >= len(board) {
			break
		}

		if board[ny][nx] == 1 {
			dirc += 1
			if dirc > 4 {
				dirc = 1
			}
		} else {
			board[ny][nx] = 2
			x, y = nx, ny
		}
	}

	return board
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	board := [][]int{}
	dirc := 0
	bx := 0
	by := 0
	y := 0
	for sc.Scan() {
		line := sc.Text()
		cols := strings.Split(line, "")
		b := []int{}
		x := 0

		for _, col := range cols {
			if col == "." {
				b = append(b, 0)
			} else if col == "#" {
				b = append(b, 1)
			} else {
				b = append(b, 2)
				bx = x
				by = y
				if col == "^" {
					dirc = 1
				} else if col == ">" {
					dirc = 2
				} else if col == "v" {
					dirc = 3
				} else if col == "<" {
					dirc = 4
				} else {
					panic("invalid direction: " + col)
				}
			}
			x += 1
		}
		board = append(board, b)
		y += 1
	}

	for _, b := range board {
		fmt.Printf("%v\n", b)
	}
	fmt.Printf("(%d, %d) dirc: %d\n\n", bx, by, dirc)

	nb := walk(board, bx, by, dirc)
	ans := 0
	for _, b := range nb {
		fmt.Printf("%v\n", b)
		for _, bb := range b {
			if bb == 2 {
				ans += 1
			}
		}
	}

	fmt.Printf("Answer: %d\n", ans)
}
