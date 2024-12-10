package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp

func readInput(input string) (int, int) {
	input = strings.TrimSpace(input)

	parts := re.Split(input, -1)
	if len(parts) != 2 {
		log.Fatalf("invalid line: %s", input)
	}
	left, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return left, right
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	re = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	input := ""
	for sc.Scan() {
		input += sc.Text()
	}

	ans := 0
	do := true
	matches := re.FindAllStringSubmatch(input, -1)
	fmt.Printf("%v\n", matches)
	for _, m := range matches {
		if m[0] == "do()" {
			do = true
			continue
		} else if m[0] == "don't()" {
			do = false
			continue
		}

		if !do {
			continue
		}

		n1, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		ans += n1 * n2
	}

	fmt.Printf("---------\nAnswer: %d\n", ans)
}
