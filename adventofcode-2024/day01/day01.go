package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./day01 <FILE>")
		os.Exit(1)
	}
	fd, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("failed to open")
	}
	sc := bufio.NewScanner(fd)
	var left, right []int
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			continue
		}
		li, errl := strconv.Atoi(parts[0])
		ri, errr := strconv.Atoi(parts[1])
		if errl != nil || errr != nil {
			log.Fatalf("failed to parse")
		}
		left = append(left, li)
		right = append(right, ri)
	}
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		l := left[i]
		r := right[i]
		delta := max(l, r) - min(l, r)
		sum += delta
	}
	fmt.Println("part1: ", sum)

	similarity := 0
	/* freq := make(map[int]int)
	for _, v := range right {
		freq[v] += 1
	}
	for _, v := range left {
		similarity += v * freq[v]
	} */

	ln := left[0]
	lc := 0
	ri := 0
	for _, lv := range left {
		if lv == ln {
			lc += 1
		} else {
			for right[ri] < ln {
				ri += 1
			}
			rc := 0
			for right[ri] == ln {
				rc += 1
				ri += 1
			}
			similarity += ln * lc * rc

			ln = lv
			lc = 1
		}

	}
	fmt.Println("part2: ", similarity)
}
