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

func readInput(input string) []int {
	input = strings.TrimSpace(input)
	nums := []int{}

	parts := re.Split(input, -1)
	if len(parts) == 0 {
		log.Fatalf("invalid line: %s", input)
	}
	for _, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	return nums
}

func checkSafe(rs []int) bool {
	base := 0
	updown := 0
	for i, r := range rs {
		if i == 0 {
			base = r
			continue
		}

		d := base - r
		if d == 0 {
			return false
		}
		if d < -3 || d > 3 {
			return false
		}
		if d > 0 {
			if updown == 0 {
				updown = 1
			} else if updown == -1 {
				return false
			}
		}
		if d < 0 {
			if updown == 0 {
				updown = -1
			} else if updown == 1 {
				return false
			}
		}
		base = r
	}

	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	re = regexp.MustCompile(`\s+`)

	safe := 0
	for {
		if ok := sc.Scan(); !ok {
			break
		}

		input := sc.Text()
		if len(input) > 0 {
			rs := readInput(input)

			if checkSafe(rs) {
				safe++
				log.Printf("Safe: %v", rs)
			} else {
				for i := range rs {
					nrs := make([]int, len(rs))
					copy(nrs, rs)
					nrs = append(nrs[:i], nrs[i+1:]...)
					if checkSafe(nrs) {
						safe++
						log.Printf("FixSafe: %v -> %v", rs, nrs)
						break
					}
				}
			}
		}
	}

	fmt.Printf("---------\nAnswer: %d\n", safe)
}
