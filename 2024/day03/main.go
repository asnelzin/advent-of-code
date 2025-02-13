package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var commandsRe = regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)

func main() {
	part1()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	sum := 0
	mulEnabled := true
	for _, command := range commandsRe.FindAllString(string(b), -1) {
		prefix := command[:3]
		switch prefix {
		case "mul":
			if !mulEnabled {
				continue
			}
			var l, r int
			n, err := fmt.Sscanf(command, "mul(%d,%d)", &l, &r)
			if err != nil || n != 2 {
				log.Fatalf("could not scan mul command: %v", err)
			}
			sum += l * r
		case "do(":
			mulEnabled = true
		case "don":
			mulEnabled = false
		default:
			log.Fatalf("unrecognized command: %q", command)
		}
	}
	fmt.Println(sum)
}

func part2() {
}
