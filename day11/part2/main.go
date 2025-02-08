package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func twinkle(stones map[int]int) map[int]int {
	ns := map[int]int{}
	for stone, cnt := range stones {
		if stone == 0 {
			ns[1] += cnt
		} else {
			digits := 0
			temp := stone
			for temp > 0 {
				temp /= 10
				digits++
			}

			if digits%2 != 0 {
				ns[stone*2024] += cnt
				continue
			}

			half := digits / 2
			div := int(math.Pow10(half))

			ns[stone/div] += cnt
			ns[stone%div] += cnt
		}
	}

	return ns
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	line := sc.Text()

	data := strings.Split(line, " ")
	stones := map[int]int{}
	for _, d := range data {
		n, _ := strconv.Atoi(d)
		stones[n]++
	}

	fmt.Printf("%v\n\n", stones)

	for i := 0; i < 75; i++ {
		stones = twinkle(stones)
		fmt.Printf("%d\n", len(stones))
	}

	fmt.Println()
	ans := 0
	for _, cnt := range stones {
		ans += cnt
	}
	fmt.Printf("Answer: %d\n", ans)
}
