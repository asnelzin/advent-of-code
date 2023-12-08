package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	parts := strings.Split(string(buf), "\n\n")
	if len(parts) < 2 {
		log.Fatalf("could not parse file: %v", err)
	}

	path := strings.Split(parts[0], "")

	network := make(map[string][2]string)
	for _, line := range strings.Split(parts[1], "\n") {
		if line == "" {
			continue
		}
		var node, left, right string
		_, err := fmt.Sscanf(line, "%3s = (%3s, %3s)", &node, &left, &right)
		if err != nil {
			log.Fatalf("could not parse line: %v", err)
		}
		network[node] = [2]string{left, right}
	}

	at := "AAA"
	pointer := 0
	steps := 0
	for at != "ZZZ" {
		direction := path[pointer]
		switch direction {
		case "L":
			at = network[at][0]
		case "R":
			at = network[at][1]
		}
		steps++

		pointer++
		if pointer >= len(path) {
			pointer = 0
		}
	}
	fmt.Println(steps)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	parts := strings.Split(string(buf), "\n\n")
	if len(parts) < 2 {
		log.Fatalf("could not parse file: %v", err)
	}

	path := strings.Split(parts[0], "")

	network := make(map[string][2]string)
	at := make([]string, 0)
	for _, line := range strings.Split(parts[1], "\n") {
		if line == "" {
			continue
		}
		var node, left, right string
		_, err := fmt.Sscanf(line, "%3s = (%3s, %3s)", &node, &left, &right)
		if err != nil {
			log.Fatalf("could not parse line: %v", err)
		}
		network[node] = [2]string{left, right}
		if strings.HasSuffix(node, "A") {
			at = append(at, node)
		}
	}

	total := 1
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, start := range at {
		wg.Add(1)
		go func(at string) {
			defer wg.Done()
			p := 0
			steps := 0

			for !strings.HasSuffix(at, "Z") {
				direction := path[p]
				switch direction {
				case "L":
					at = network[at][0]
				case "R":
					at = network[at][1]
				}
				steps++
				p++
				if p >= len(path) {
					p = 0
				}
			}

			mu.Lock()
			total = lcm(total, steps)
			mu.Unlock()
		}(start)
	}
	wg.Wait()
	fmt.Println(total)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	part1()
	part2()
}
