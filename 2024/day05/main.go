package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type set map[int]bool

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	rules := make(map[int]set)
	var pages [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			order := strings.Split(line, "|")
			if len(order) != 2 {
				log.Fatalf("rule is not correct: %s", line)
			}

			left, err := strconv.Atoi(order[0])
			if err != nil {
				log.Fatalf("could not convert left operand in rule: %s: %v", line, err)
			}

			right, err := strconv.Atoi(order[1])
			if err != nil {
				log.Fatalf("could not convert right operand in rule: %s: %v", line, err)
			}

			if rules[left] == nil {
				rules[left] = make(set)
			}
			rules[left][right] = true
		} else {
			parts := strings.Split(line, ",")
			if len(parts) < 2 {
				continue
			}

			var nums []int
			for _, p := range parts {
				n, err := strconv.Atoi(p)
				if err != nil {
					log.Fatalf("could not convert page number (%s) in line (%s): %v", p, line, err)
				}
				nums = append(nums, n)
			}
			pages = append(pages, nums)
		}
	}

	var sum int
	for _, order := range pages {
		correct := true
		prefix := make(set)
		for i := 0; i < len(order); i++ {
			if intersection(prefix, rules[order[i]]) {
				correct = false
				break
			}

			prefix[order[i]] = true
		}

		if correct {
			sum += order[(len(order)-1)/2]
		}
	}

	fmt.Println(sum)
}

func intersection(a set, b set) bool {
	if len(a) > len(b) {
		a, b = b, a
	}
	for k, _ := range a {
		if b[k] {
			return true
		}
	}
	return false
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	rules := make(map[int]set)
	var pages [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			order := strings.Split(line, "|")
			if len(order) != 2 {
				log.Fatalf("rule is not correct: %s", line)
			}

			left, err := strconv.Atoi(order[0])
			if err != nil {
				log.Fatalf("could not convert left operand in rule: %s: %v", line, err)
			}

			right, err := strconv.Atoi(order[1])
			if err != nil {
				log.Fatalf("could not convert right operand in rule: %s: %v", line, err)
			}

			if rules[left] == nil {
				rules[left] = make(set)
			}
			rules[left][right] = true
		} else {
			parts := strings.Split(line, ",")
			if len(parts) < 2 {
				continue
			}

			var nums []int
			for _, p := range parts {
				n, err := strconv.Atoi(p)
				if err != nil {
					log.Fatalf("could not convert page number (%s) in line (%s): %v", p, line, err)
				}
				nums = append(nums, n)
			}
			pages = append(pages, nums)
		}
	}

	sum := 0
	for _, order := range pages {
		wasFixed := false
		for i := 0; i < len(order); i++ {
			if correct(order[i], order[:i], rules) {
				continue
			}

			for j := i; j > 0; j-- {
				order[j], order[j-1] = order[j-1], order[j]
				if correct(order[j-1], order[:j-1], rules) {
					break
				}
			}
			wasFixed = true
		}
		if wasFixed {
			sum += order[(len(order)-1)/2]
		}
	}

	fmt.Println(sum)
}

func correct(n int, prefix []int, rules map[int]set) bool {
	for _, k := range prefix {
		if rules[n][k] {
			return false
		}
	}
	return true
}

func main() {
	part2()
}
