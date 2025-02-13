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

	var field [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []byte(line))
	}

	total := 0
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] != 'X' {
				continue
			}

			for _, d := range directions {
				if tryDirection(field, d, Point{x, y}) {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

type Direction struct {
	x, y int
}
type Point struct {
	x, y int
}

func (p Point) Add(d Direction) Point {
	return Point{p.x + d.x, p.y + d.y}
}

var (
	left      = Direction{-1, 0}
	right     = Direction{1, 0}
	up        = Direction{0, 1}
	down      = Direction{0, -1}
	upright   = Direction{1, 1}
	downright = Direction{1, -1}
	downleft  = Direction{-1, -1}
	upleft    = Direction{-1, 1}
)

var directions = []Direction{left, right, up, down, upright, downright, downleft, upleft}

func tryDirection(field [][]byte, d Direction, from Point) bool {
	word := []byte{'X', 'M', 'A', 'S'}
	for _, letter := range word {
		if !inBounds(field, from) {
			return false
		}
		if field[from.y][from.x] != letter {
			return false
		}
		from = from.Add(d)
	}

	return true
}

func inBounds(field [][]byte, p Point) bool {
	if p.y < 0 || p.y >= len(field) {
		return false
	}
	if p.x < 0 || p.x >= len(field[p.y]) {
		return false
	}
	return true
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var field [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []byte(line))
	}

	total := 0
	for y := 1; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] != 'A' {
				continue
			}

			if tryPoint(field, Point{x, y}) {
				total++
			}
		}
	}
	fmt.Println(total)
}

func tryPoint(field [][]byte, p Point) bool {
	// S   S
	//   A
	// M   M
	if at(field, p.Add(upleft)) == 'S' && at(field, p.Add(upright)) == 'S' && at(field, p.Add(downleft)) == 'M' && at(field, p.Add(downright)) == 'M' {
		return true
	}

	// M   M
	//   A
	// S   S
	if at(field, p.Add(upleft)) == 'M' && at(field, p.Add(upright)) == 'M' && at(field, p.Add(downleft)) == 'S' && at(field, p.Add(downright)) == 'S' {
		return true
	}

	// M   S
	//   A
	// M   S
	if at(field, p.Add(upleft)) == 'M' && at(field, p.Add(upright)) == 'S' && at(field, p.Add(downleft)) == 'M' && at(field, p.Add(downright)) == 'S' {
		return true
	}

	// S   M
	//   A
	// S   M
	if at(field, p.Add(upleft)) == 'S' && at(field, p.Add(upright)) == 'M' && at(field, p.Add(downleft)) == 'S' && at(field, p.Add(downright)) == 'M' {
		return true
	}

	return false
}

func at(field [][]byte, p Point) byte {
	if !inBounds(field, p) {
		return 'X'
	}

	return field[p.y][p.x]
}

func main() {
	part2()
}
