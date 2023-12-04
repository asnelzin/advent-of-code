package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var points int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			log.Fatalf("could not parse line: %v", line)
		}

		parts = strings.Split(parts[1], "|")
		if len(parts) < 2 {
			log.Fatalf("could not parse card numbers: %v", line)
		}

		amount := 0
		winning := make(map[int]bool)
		for _, s := range strings.Split(parts[0], " ") {
			if s == "" {
				continue
			}

			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("could not parse card's winning number: %v", s)
			}
			winning[n] = true
		}

		for _, s := range strings.Split(parts[1], " ") {
			if s == "" {
				continue
			}
			candidate, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("could not parse card's candidate number: %v", s)
			}
			if _, ok := winning[candidate]; ok {
				amount++
			}
		}
		if amount > 0 {
			points += int(math.Pow(2, float64(amount-1)))
		}
	}
	fmt.Println(points)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	id := 1
	cards := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			log.Fatalf("could not parse line: %v", line)
		}

		parts = strings.Split(parts[1], "|")
		if len(parts) < 2 {
			log.Fatalf("could not parse card numbers: %v", line)
		}

		cards[id] += 1

		winning := make(map[int]bool)
		for _, s := range strings.Split(parts[0], " ") {
			if s == "" {
				continue
			}

			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("could not parse card's winning number: %v", s)
			}
			winning[n] = true
		}

		points := 0
		for _, s := range strings.Split(parts[1], " ") {
			if s == "" {
				continue
			}
			candidate, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("could not parse card's candidate number: %v", s)
			}
			if _, ok := winning[candidate]; ok {
				points++
			}
		}

		for i := 0; i < cards[id]; i++ {
			for j := 1; j <= points; j++ {
				cards[id+j] += 1
			}
		}
		id++
	}

	sum := 0
	for _, v := range cards {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	part2()
}
