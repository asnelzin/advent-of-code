package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	redLimit   = 12
	greenLimit = 13
	blueLimit  = 14
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		idstring := strings.TrimPrefix(parts[0], "Game ")
		gameID, err := strconv.Atoi(idstring)
		if err != nil {
			log.Fatalf("could not parse id: %v", err)
		}

		takeouts := strings.Split(parts[1], ";")
		isOver := false
		for _, takeout := range takeouts {
			for _, cubes := range strings.Split(takeout, ",") {
				cubes := strings.TrimSpace(cubes)
				p := strings.Split(cubes, " ")
				if len(p) > 2 {
					log.Fatalf("could not parse cubes record: `%v`", cubes)
				}

				nums := p[0]
				color := p[1]
				num, err := strconv.Atoi(nums)
				if err != nil {
					log.Fatalf("could not parse nums: %v", err)
				}

				switch color {
				case "red":
					if num > redLimit {
						isOver = true
					}
				case "green":
					if num > greenLimit {
						isOver = true
					}
				case "blue":
					if num > blueLimit {
						isOver = true
					}
				}
			}
		}
		if !isOver {
			sum += gameID
		}
	}
	fmt.Println(sum)
}

type Triplet struct {
	Red   int
	Green int
	Blue  int
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		takeouts := strings.Split(parts[1], ";")
		triplet := Triplet{}
		for _, takeout := range takeouts {
			for _, cubes := range strings.Split(takeout, ",") {
				cubes := strings.TrimSpace(cubes)
				p := strings.Split(cubes, " ")
				if len(p) > 2 {
					log.Fatalf("could not parse cubes record: `%v`", cubes)
				}

				nums := p[0]
				color := p[1]
				num, err := strconv.Atoi(nums)
				if err != nil {
					log.Fatalf("could not parse nums: %v", err)
				}

				switch color {
				case "red":
					triplet.Red = max(num, triplet.Red)
				case "green":
					triplet.Green = max(num, triplet.Green)
				case "blue":
					triplet.Blue = max(num, triplet.Blue)
				}
			}
		}

		sum += triplet.Red * triplet.Green * triplet.Blue
	}
	fmt.Println(sum)
}

func main() {
	part2()
}
