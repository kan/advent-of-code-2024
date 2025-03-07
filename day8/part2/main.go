package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkAnti(board [][]rune, antiMap map[string]struct{}, c rune, cx, cy int) map[string]struct{} {
	for y, b := range board {
		for x, col := range b {
			if cx == x && cy == y {
				continue
			}
			if col != c {
				continue
			}
			ax := cx - (x - cx)
			ay := cy - (y - cy)
			for {
				if ax < 0 || ay < 0 || ax >= len(board[0]) || ay >= len(board) {
					// アンチノードが盤面外
					break
				}
				antiMap[fmt.Sprintf("%d,%d", ax, ay)] = struct{}{}
				ax -= (x - cx)
				ay -= (y - cy)
			}
		}
	}

	return antiMap
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	board := [][]rune{}
	for sc.Scan() {
		line := sc.Text()
		cols := strings.Split(line, "")
		b := []rune{}
		for _, col := range cols {
			b = append(b, []rune(col)[0])
		}
		board = append(board, b)
	}

	antiMap := map[string]struct{}{}
	for y, b := range board {
		for x, c := range b {
			fmt.Printf("%c", c)
			if c == '.' {
				continue
			}
			antiMap = checkAnti(board, antiMap, c, x, y)
		}
		fmt.Println()
	}
	fmt.Println()

	an, ln := 0, 0
	for y, b := range board {
		for x, c := range b {
			if c == '.' {
				if _, ok := antiMap[fmt.Sprintf("%d,%d", x, y)]; ok {
					fmt.Print("#")
					an++
					continue
				}
			} else {
				ln++
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Printf("Answer: %d(%d + %d)\n", an+ln, an, ln)
}
