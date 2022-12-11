package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func draw(cycle int, x int) {
	pixel := (cycle - 1) % 40

	// x means that sprite is currently in (x-1, x, x+1) pixels
	if x-1 <= pixel && pixel <= x+1 {
		print("#")
	} else {
		print(".")
	}

	if cycle%40 == 0 {
		print("\n")
	}
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cycle := 0
	register := 1
	prev := register

	checkpoint := 20
	totalSignal := 0

	for scanner.Scan() {
		argv := strings.Split(scanner.Text(), " ")
		if len(argv) == 0 {
			continue
		}
		var (
			op  string
			arg int
		)
		op = argv[0]
		if len(argv) > 1 {
			arg, err = strconv.Atoi(argv[1])
			if err != nil {
				log.Printf("could not parse argument %s: %v", argv[1], err)
				continue
			}
		}

		prev = register
		switch op {
		case "noop":
			cycle += 1
		case "addx":
			cycle += 2
			register += arg
		}

		dx := cycle % checkpoint
		if dx == 0 || dx == 1 {
			totalSignal += prev * (cycle - dx)
			checkpoint += 40
		}
	}

	fmt.Println(totalSignal)

}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cycle := 1
	register := 1

	for scanner.Scan() {
		argv := strings.Split(scanner.Text(), " ")
		if len(argv) == 0 {
			continue
		}
		var (
			op  string
			arg int
		)
		op = argv[0]
		if len(argv) > 1 {
			arg, err = strconv.Atoi(argv[1])
			if err != nil {
				log.Printf("could not parse argument %s: %v", argv[1], err)
				continue
			}
		}

		switch op {
		case "noop":
			draw(cycle, register)
			cycle++
		case "addx":
			draw(cycle, register)
			cycle++
			draw(cycle, register)
			cycle++
			register += arg
		}

	}
}

func main() {
	part1()
	part2()
}
