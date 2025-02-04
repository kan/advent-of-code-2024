package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gc(disk []int) []int {
	id := -1
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == -1 {
			continue
		}
		if id <= disk[i] && id != -1 {
			continue
		}
		id = disk[i]
		fl := 0
		for {
			if i-fl < 0 || disk[i] != disk[i-fl] {
				break
			}
			fl++
		}
		//log.Printf("id=%d, fl=%d\n", disk[i], fl)
	LOOP:
		for j := 0; j < i; j++ {
			if disk[j] == -1 {
				for k := j; k < j+fl; k++ {
					if disk[k] != -1 {
						j++
						continue LOOP
					}
				}
				for k := 0; k < fl; k++ {
					disk[j+k] = id
					disk[i-k] = -1
				}
				j += fl
				break
			}
		}
		i -= fl - 1
	}
	return disk
}

func checksum(disk []int) int {
	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
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
