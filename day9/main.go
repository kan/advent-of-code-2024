package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gc(disk []int) []int {
	j := 0
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == -1 {
			continue
		}
		for j < i {
			if disk[j] == -1 {
				disk[j] = disk[i]
				disk[i] = -1
				break
			}
			j++
		}
	}
	return disk
}

func checksum(disk []int) int {
	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			break
		}
		sum += i * disk[i]
	}
	return sum
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	line := sc.Text()
	fmt.Printf("%s\n", line)

	id := 0
	gap := false
	disk := []int{}
	for _, c := range line {
		n, _ := strconv.Atoi(string(c))
		for i := 0; i < n; i++ {
			if gap {
				fmt.Print(".")
				disk = append(disk, -1)
			} else {
				fmt.Printf("%d", id)
				disk = append(disk, id)
			}
		}
		if !gap {
			id++
		}
		gap = !gap
	}
	fmt.Println()

	disk = gc(disk)
	for _, v := range disk {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", v)
		}
	}
	fmt.Println()

	fmt.Printf("Answer: %d\n", checksum(disk))
}
