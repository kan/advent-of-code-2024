package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var rePos = regexp.MustCompile(`X=(\d+), Y=(\d+)`)

func getPos(txt string) (int, int, error) {
	m := rePos.FindStringSubmatch(txt)
	x, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(m[2])
	if err != nil {
		return 0, 0, err
	}

	return 10000000000000 + x, 10000000000000 + y, nil
}

var reDelta = regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)

func getDelta(txt string) (int, int, error) {
	m := reDelta.FindStringSubmatch(txt)
	x, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(m[2])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func getToken(ax, ay, bx, by, px, py int) (int, int) {
	det := ax*by - ay*bx
	if det == 0 {
		return 0, 0
	}

	if (px*by-py*bx)%det != 0 || (ax*py-ay*px)%det != 0 {
		return 0, 0
	}

	return (px*by - py*bx) / det, (ax*py - ay*px) / det
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	ans := 0

	for sc.Scan() {
		line := sc.Text()
		ax, ay, err := getDelta(line)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ax: %d, ay: %d\n", ax, ay)

		sc.Scan()
		line = sc.Text()
		bx, by, err := getDelta(line)
		if err != nil {
			panic(err)
		}
		fmt.Printf("bx: %d, by: %d\n", bx, by)

		sc.Scan()
		line = sc.Text()
		px, py, err := getPos(line)
		if err != nil {
			panic(err)
		}
		fmt.Printf("px: %d, py: %d\n", px, py)

		at, bt := getToken(ax, ay, bx, by, px, py)
		fmt.Printf("at: %d, bt: %d token:%d\n\n", at, bt, at*3+bt)
		ans += at*3 + bt

		sc.Scan()
	}

	fmt.Printf("Answer: %d\n", ans)
}
