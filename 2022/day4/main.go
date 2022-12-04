package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Range struct {
	left  int
	right int
}

func RangeFromString(s string) Range {
	var (
		left  int
		right int
	)
	fmt.Sscanf(s, "%d-%d", &left, &right)
	return Range{left, right}
}

func (r Range) Contains(other Range) bool {
	return r.left <= other.left && r.right >= other.right
}

func (r Range) Overlap(other Range) bool {
	return (r.right >= other.left && r.right <= other.right) ||
		(other.right >= r.left && other.right <= r.right)
}

func (r Range) String() string {
	return fmt.Sprintf("Range<%d-%d>", r.left, r.right)
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := scanner.Text()
		pairs := strings.Split(pair, ",")

		first := RangeFromString(pairs[0])
		second := RangeFromString(pairs[1])
		if first.Contains(second) || second.Contains(first) {
			total++
		}
	}

	fmt.Println(total)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := scanner.Text()
		pairs := strings.Split(pair, ",")

		first := RangeFromString(pairs[0])
		second := RangeFromString(pairs[1])
		if first.Overlap(second) {
			total++
		}
	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
