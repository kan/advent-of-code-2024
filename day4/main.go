package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(s []string) []string {
	r := make([]string, len(s))
	copy(r, s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func checkXmas(data []string) bool {
	s := ""
	for _, c := range data {
		if c == "X" || c == "M" || c == "A" || c == "S" {
			s += c
		}
		if len(s) == 4 {
			break
		}
	}

	return s == "XMAS"
}

func getXmas(m [][]string, x, y int) int {
	ans := 0
	// 右 →
	if checkXmas(m[y][x:]) {
		ans++
	}
	// 左 ←
	if checkXmas(reverse(m[y][0 : x+1])) {
		ans++
	}
	// 下 ↓
	s := []string{}
	for i := y; i < len(m); i++ {
		s = append(s, m[i][x])
	}
	if checkXmas(s) {
		ans++
	}
	// 上 ↑
	s = []string{}
	for i := y; i >= 0; i-- {
		s = append(s, m[i][x])
	}
	if checkXmas(s) {
		ans++
	}
	// 右下
	s = []string{}
	for i, j := y, x; i < len(m) && j < len(m[i]); i, j = i+1, j+1 {
		s = append(s, m[i][j])
	}
	if checkXmas(s) {
		ans++
	}
	// 左下
	s = []string{}
	for i, j := y, x; i < len(m) && j >= 0; i, j = i+1, j-1 {
		s = append(s, m[i][j])
	}
	if checkXmas(s) {
		ans++
	}
	// 右上
	s = []string{}
	for i, j := y, x; i >= 0 && j < len(m[i]); i, j = i-1, j+1 {
		s = append(s, m[i][j])
	}
	if checkXmas(s) {
		ans++
	}
	// 左上
	s = []string{}
	for i, j := y, x; i >= 0 && j >= 0; i, j = i-1, j-1 {
		s = append(s, m[i][j])
	}
	if checkXmas(s) {
		ans++
	}
	return ans
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
			ans += getXmas(m, x, y)
		}
	}
	fmt.Printf("%d x %d\n", len(m[0]), len(m))

	fmt.Printf("---------\nAnswer: %d\n", ans)
}
