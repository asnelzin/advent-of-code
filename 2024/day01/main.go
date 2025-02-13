package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

func main() {
	part2()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var l, r int
		fmt.Sscanf(line, "%d %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	var sum int
	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println(sum)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var left []int
	right := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var l, r int
		n, err := fmt.Sscanf(line, "%d %d", &l, &r)
		if err != nil || n != 2 {
			log.Fatalf("could not scan line %q: %v", line, err)
		}
		left = append(left, l)
		right[r] += 1
	}

	var sum int
	for _, l := range left {
		r, ok := right[l]
		if !ok {
			r = 0
		}
		sum += l * r
	}
	fmt.Println(sum)
}
