package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkChar(m [][]string, x, y, dx, dy int, ch string) bool {
	for i, j := y+dy, x+dx; i >= 0 && i < len(m) && j >= 0 && j < len(m[i]); i, j = i+dy, j+dx {
		return m[i][j] == ch
	}
	return false
}

func getXmas(m [][]string, x, y int) bool {
	if m[y][x] != "A" {
		return false
	}
	// 左上から右下
	if checkChar(m, x, y, -1, -1, "M") && checkChar(m, x, y, 1, 1, "S") {
		// 成立(MAS)
	} else if checkChar(m, x, y, -1, -1, "S") && checkChar(m, x, y, 1, 1, "M") {
		// 成立(SAM)
	} else {
		return false
	}
	// 右上から左下
	if checkChar(m, x, y, -1, 1, "M") && checkChar(m, x, y, 1, -1, "S") {
		// 成立(MAS)
	} else if checkChar(m, x, y, -1, 1, "S") && checkChar(m, x, y, 1, -1, "M") {
		// 成立(SAM)
	} else {
		return false
	}

	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var m [][]string
	for sc.Scan() {
		line := sc.Text()
		rows := strings.Split(line, "")
		m = append(m, rows)
	}

	ans := 0
	for y, cols := range m {
		for x := range cols {
			if getXmas(m, x, y) {
				ans++
			}
		}
	}
	fmt.Printf("%d x %d\n", len(m[0]), len(m))

	fmt.Printf("---------\nAnswer: %d\n", ans)
}
