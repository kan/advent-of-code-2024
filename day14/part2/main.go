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

func getMap(robots []Robot, sec int) string {
	grid := make([][]int, HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		grid[i] = make([]int, WIDTH)
	}

	for _, robo := range robots {
		r := Robot{robo.PosX, robo.PosY, robo.SpeedX, robo.SpeedY}
		for i := 0; i < sec; i++ {
			r.move()
		}
		grid[r.PosY][r.PosX]++
	}

	res := ""
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if grid[y][x] > 0 {
				res += "#"
			} else {
				res += "."
			}
		}
		res += "\n"
	}

	return res
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	robots := []Robot{}
	for sc.Scan() {
		line := sc.Text()
		robo, err := getRobot(line)
		if err != nil {
			panic(err)
		}
		robots = append(robots, *robo)
	}

	mmap := map[string]int{}
	for i := 0; i < 100000; i++ {
		m := getMap(robots, i)
		if d, ok := mmap[m]; ok {
			fmt.Println(m)
			fmt.Printf("Time: %d -> %d\n", d, i)
			break
		}
		mmap[m] = i
	}
}
