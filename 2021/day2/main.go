package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	horizontal, depth := 0, 0

	for scanner.Scan() {
		var (
			command string
			arg     int
		)
		fmt.Sscanf(scanner.Text(), "%s %d", &command, &arg)

		switch command {
		case "forward":
			horizontal += arg
		case "down":
			depth += arg
		case "up":
			depth -= arg
		}
	}

	fmt.Println(horizontal * depth)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	horizontal, depth, aim := 0, 0, 0

	for scanner.Scan() {
		var (
			command string
			arg     int
		)
		fmt.Sscanf(scanner.Text(), "%s %d", &command, &arg)

		switch command {
		case "forward":
			horizontal += arg
			depth += aim * arg
		case "down":
			aim += arg
		case "up":
			aim -= arg
		}
	}

	fmt.Println(horizontal * depth)
}

func main() {
	part1()
	part2()
}
