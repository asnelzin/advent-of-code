package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var stones []int
	for scanner.Scan() {
		line := scanner.Text()
		for _, s := range strings.Split(line, " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("could not convert %s to number: %v", s, err)
			}

			stones = append(stones, n)
		}
	}

	fmt.Println(stones)

	for i := 0; i < 25; i++ {
		var updated []int

		for _, stone := range stones {
			// rule 1
			if stone == 0 {
				updated = append(updated, 1)
				continue
			}
			// rule 2
			asString := strconv.Itoa(stone)
			if len(asString)%2 == 0 {
				left, err := strconv.Atoi(asString[:len(asString)/2])
				if err != nil {
					log.Fatalf("could not convert to number: %v", err)
				}
				right, err := strconv.Atoi(asString[len(asString)/2:])
				if err != nil {
					log.Fatalf("could not convert to number: %v", err)
				}
				updated = append(updated, []int{left, right}...)
				continue
			}
			// rule 3
			updated = append(updated, stone*2024)
		}

		stones = updated
		// fmt.Println(stones)
	}
	fmt.Println(len(stones))
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var stones []int
	for scanner.Scan() {
		line := scanner.Text()
		for _, s := range strings.Split(line, " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("could not convert %s to number: %v", s, err)
			}

			stones = append(stones, n)
		}
	}

	fmt.Println(stones)

	cache := make(map[pair]int)
	total := 0
	for _, stone := range stones {
		total += expand(stone, 75, cache)
	}

	fmt.Println(total)
}

type pair struct {
	n, b int
}

func expand(n int, b int, cache map[pair]int) int {
	if b == 0 {
		return 1
	}
	if v, ok := cache[pair{n, b}]; ok {
		return v
	}

	var r int
	// rule 1
	if n == 0 {
		r = expand(1, b-1, cache)
	} else if len(strconv.Itoa(n))%2 == 0 {
		asString := strconv.Itoa(n)
		left, err := strconv.Atoi(asString[:len(asString)/2])
		if err != nil {
			log.Fatalf("could not convert to number: %v", err)
		}
		right, err := strconv.Atoi(asString[len(asString)/2:])
		if err != nil {
			log.Fatalf("could not convert to number: %v", err)
		}
		r = expand(left, b-1, cache) + expand(right, b-1, cache)
	} else {
		r = expand(n*2024, b-1, cache)
	}

	cache[pair{n, b}] = r
	return r
}

func main() {
	// part1()
	part2()
}
