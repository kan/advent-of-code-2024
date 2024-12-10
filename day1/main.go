package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
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

	slices.Sort(left)
	slices.Sort(right)

	sumd := 0
	for i := range left {
		d := left[i] - right[i]
		if d < 0 {
			d *= -1
		}
		log.Printf("%d: %d - %d = %d\n", i, left[i], right[i], d)
		sumd += d
	}

	//fmt.Printf("left: %v\nright: %v\n", left, right)
	fmt.Printf("---------\nAnswer: %d\n", sumd)
}
