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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	re = regexp.MustCompile(`\s+`)

	safe := 0
LOOP:
	for {
		if ok := sc.Scan(); !ok {
			break
		}

		input := sc.Text()
		if len(input) > 0 {
			rs := readInput(input)

			base := 0
			updown := 0
			for i, r := range rs {
				if i == 0 {
					base = r
					continue
				}

				d := base - r
				if d == 0 {
					log.Println("zero")
					continue LOOP
				}
				if d < -3 || d > 3 {
					log.Println("over3")
					continue LOOP
				}
				if d > 0 {
					if updown == 0 {
						updown = 1
					} else if updown == -1 {
						log.Println("up to down")
						continue LOOP
					}
				}
				if d < 0 {
					if updown == 0 {
						updown = -1
					} else if updown == 1 {
						log.Println("down to up")
						continue LOOP
					}
				}
				base = r
			}
			log.Printf("Safe: %v", rs)
			safe++
		}
	}

	fmt.Printf("---------\nAnswer: %d\n", safe)
}
