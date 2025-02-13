package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
outer:
	for scanner.Scan() {
		line := scanner.Text()
		var seq []int
		for _, d := range strings.Split(line, " ") {
			n, err := strconv.Atoi(d)
			if err != nil {
				log.Fatalf("could not convert %s to int: %v", d, err)
			}
			seq = append(seq, n)
		}

		diff := abs(seq[1] - seq[0])
		if diff < 1 || diff > 3 {
			continue
		}

		inc := true
		if seq[1]-seq[0] < 0 {
			inc = false
		}
		for i := 2; i < len(seq); i++ {
			d := seq[i] - seq[i-1]
			if (d < 0 && inc) || (d > 0 && !inc) || (abs(d) < 1 || abs(d) > 3) {
				continue outer
			}
		}
		sum += 1
	}
	fmt.Println(sum)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var seq []int
		for _, d := range strings.Split(line, " ") {
			n, err := strconv.Atoi(d)
			if err != nil {
				log.Fatalf("could not convert %s to int: %v", d, err)
			}
			seq = append(seq, n)
		}

		if check(seq) {
			sum += 1
			continue
		}

		for i := 0; i < len(seq); i++ {
			try := make([]int, len(seq))
			copy(try, seq)
			try = append(try[:i], try[i+1:]...)
			if check(try) {
				sum += 1
				break
			}
		}
	}
	fmt.Println(sum)
}

func check(seq []int) bool {
	fmt.Println(seq)
	diff := abs(seq[1] - seq[0])
	if diff < 1 || diff > 3 {
		return false
	}

	inc := true
	if seq[1]-seq[0] < 0 {
		inc = false
	}

	for i := 2; i < len(seq); i++ {
		d := seq[i] - seq[i-1]
		if (d < 0 && inc) || (d > 0 && !inc) || (abs(d) < 1 || abs(d) > 3) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
