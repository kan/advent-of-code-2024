package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule map[string][]string

func checkRow(rules, rrules Rule, row []string) bool {
	for i, col := range row {
		m, ok := rules[col]
		if !ok {
			continue
		}
		for _, v := range row[i+1:] {
			flg := false
			for _, r := range m {
				if v == r {
					flg = true
					break
				}
			}
			if !flg {
				return false
			}
		}
		if i == 0 {
			continue
		}
		m2, ok := rrules[col]
		if !ok {
			continue
		}
		for _, v := range row[:i] {
			flg := false
			for _, r := range m2 {
				if v == r {
					flg = true
					break
				}
			}
			if !flg {
				return false
			}
		}
	}

	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	ans := 0
	mode := 0
	rules := Rule{}
	rrules := Rule{}
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			fmt.Printf("rules: %+v\n", rules)
			fmt.Printf("rrules: %+v\n", rrules)
			mode = 1
			continue
		}
		if mode == 0 {
			rows := strings.Split(line, "|")
			if _, ok := rules[rows[0]]; !ok {
				rules[rows[0]] = []string{}
			}
			rules[rows[0]] = append(rules[rows[0]], rows[1])
			if _, ok := rrules[rows[1]]; !ok {
				rrules[rows[1]] = []string{}
			}
			rrules[rows[1]] = append(rrules[rows[1]], rows[0])
		} else {
			row := strings.Split(line, ",")
			fmt.Printf("row: %+v check:%t\n", row, checkRow(rules, rrules, row))
			if checkRow(rules, rrules, row) {
				num, err := strconv.Atoi(row[len(row)/2])
				if err != nil {
					panic(err)
				}
				fmt.Printf("center: %d\n", num)
				ans += num
			}
		}
	}

	fmt.Printf("---------\nAnswer: %d\n", ans)
}
