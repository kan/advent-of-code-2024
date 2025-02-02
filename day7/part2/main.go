package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func concat(a, b int) int {
	s := fmt.Sprintf("%d%d", a, b)
	d, _ := strconv.Atoi(s)
	return d
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	ans := 0
	for sc.Scan() {
		line := sc.Text()
		data := strings.Split(line, ": ")
		result, err := strconv.Atoi(data[0])
		if err != nil {
			panic(err)
		}
		nums := strings.Split(data[1], " ")
		calcs := []int{}
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			if len(calcs) == 0 {
				calcs = append(calcs, n)
				continue
			}
			ncalcs := []int{}
			for _, calc := range calcs {
				if calc+n <= result {
					ncalcs = append(ncalcs, calc+n)
				}
				if calc*n <= result {
					ncalcs = append(ncalcs, calc*n)
				}
				if concat(calc, n) <= result {
					ncalcs = append(ncalcs, concat(calc, n))
				}
			}
			calcs = ncalcs
		}
		if len(calcs) == 0 {
			continue
		}
		for _, calc := range calcs {
			if calc == result {
				ans += result
				break
			}
		}
	}

	fmt.Printf("Answer: %d\n", ans)
}
