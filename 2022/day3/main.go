package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parse(l string) (string, string) {
	size := len(l)
	a := l[:size/2]
	b := l[size/2:]
	return a, b
}

func intersect(a string, b string) rune {
	hash := map[rune]struct{}{}
	for _, i := range a {
		hash[i] = struct{}{}
	}

	for _, j := range b {
		if _, ok := hash[j]; ok {
			return j
		}
	}
	return 0
}

func fullIntersect(a string, b string) string {
	hash := map[rune]struct{}{}
	for _, i := range a {
		hash[i] = struct{}{}
	}

	in := []rune{}
	for _, j := range b {
		if _, ok := hash[j]; ok {
			in = append(in, j)
		}
	}
	return string(in)
}

func priority(c rune) int {
	chr := int(c)
	if int('a') <= chr && chr <= int('z') {
		return chr - int('a') + 1
	}

	if int('A') <= chr && chr <= int('Z') {
		return chr - int('A') + 27
	}

	return 0
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		c := intersect(parse(scanner.Text()))
		total += priority(c)
	}

	fmt.Println(total)
}

func findBadge(group []string) int {
	temp := fullIntersect(group[0], group[1])
	c := intersect(temp, group[2])
	fmt.Println(string(c))
	return priority(c)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	r := []string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		r = append(r, scanner.Text())
	}
	fmt.Println(r)

	i := 0
	for i < len(r) {
		fmt.Println(r[i : i+3])
		total += findBadge(r[i : i+3])
		i += 3
	}

	fmt.Println(total)
}

func main() {
	part2()
}
