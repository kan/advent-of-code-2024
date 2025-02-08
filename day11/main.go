package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twinkle(stones []string) []string {
	ns := []string{}
	for _, stone := range stones {
		if stone == "0" {
			ns = append(ns, "1")
		} else if len(stone)%2 == 0 {
			// 文字列の前半分を取得
			zs, _ := strconv.Atoi(stone[:len(stone)/2])
			ns = append(ns, strconv.Itoa(zs))
			// 文字列の後半分を取得
			ks, _ := strconv.Atoi(stone[len(stone)/2:])
			ns = append(ns, strconv.Itoa(ks))
		} else {
			n, _ := strconv.Atoi(stone)
			ns = append(ns, strconv.Itoa(n*2024))
		}
	}

	return ns
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	line := sc.Text()

	stones := strings.Split(line, " ")

	fmt.Printf("%v\n\n", stones)

	for i := 0; i < 25; i++ {
		stones = twinkle(stones)
		fmt.Printf("%v\n", stones)
	}

	fmt.Println()
	fmt.Printf("Answer: %d\n", len(stones))
}
