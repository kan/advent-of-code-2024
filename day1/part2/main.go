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
	re = regexp.MustCompile(`\s+`)

	left := []int{}
	right := []int{}

	for {
		if ok := sc.Scan(); !ok {
			break
		}

		input := sc.Text()
		if len(input) > 0 {
			l, r := readInput(input)
			left = append(left, l)
			right = append(right, r)
		}
	}

	sums := 0
	for _, num := range left {
		cnt := 0
		for _, n := range right {
			if num == n {
				cnt++
			}
		}
		sums += num * cnt
	}

	//fmt.Printf("left: %v\nright: %v\n", left, right)
	fmt.Printf("---------\nAnswer: %d\n", sums)
}
