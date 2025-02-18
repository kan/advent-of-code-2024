package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const WIDTH = 101
const HEIGHT = 103

var rePos = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

type Robot struct {
	PosX, PosY     int
	SpeedX, SpeedY int
}

func (r *Robot) move() {
	r.PosX += r.SpeedX
	r.PosY += r.SpeedY

	if r.PosX < 0 {
		r.PosX += WIDTH
	} else if r.PosX >= WIDTH {
		r.PosX -= WIDTH
	}
	if r.PosY < 0 {
		r.PosY += HEIGHT
	} else if r.PosY >= HEIGHT {
		r.PosY -= HEIGHT
	}
}

func getRobot(txt string) (*Robot, error) {
	m := rePos.FindStringSubmatch(txt)
	x, err := strconv.Atoi(m[1])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(m[2])
	if err != nil {
		return nil, err
	}
	dx, err := strconv.Atoi(m[3])
	if err != nil {
		return nil, err
	}
	dy, err := strconv.Atoi(m[4])
	if err != nil {
		return nil, err
	}

	return &Robot{x, y, dx, dy}, nil
}

func getQuadrant(x, y int) int {
	if x == WIDTH/2 || y == HEIGHT/2 {
		return 0
	}

	if x < WIDTH/2 {
		if y < HEIGHT/2 {
			return 1
		} else {
			return 3
		}
	} else {
		if y < HEIGHT/2 {
			return 2
		} else {
			return 4
		}
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	qcnt := map[int]int{}

	for sc.Scan() {
		line := sc.Text()
		robo, err := getRobot(line)
		if err != nil {
			panic(err)
		}
		for i := 0; i < 100; i++ {
			robo.move()
		}
		q := getQuadrant(robo.PosX, robo.PosY)
		qcnt[q]++
	}

	fmt.Printf("%d, %d, %d, %d\n", qcnt[1], qcnt[2], qcnt[3], qcnt[4])

	fmt.Printf("Answer: %d\n", qcnt[1]*qcnt[2]*qcnt[3]*qcnt[4])
}
