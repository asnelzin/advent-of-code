package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction struct {
	x, y int
}
type Point struct {
	x, y int
}

func (p *Point) Add(d Direction) Point {
	return Point{p.x + d.x, p.y + d.y}
}

func (p *Point) Sub(d Direction) Point {
	return Point{p.x - d.x, p.y - d.y}
}

var (
	left  = Direction{-1, 0}
	right = Direction{1, 0}
	down  = Direction{0, 1}
	up    = Direction{0, -1}
)
var directions = []Direction{left, right, down, up}

func inBounds[T any](field [][]T, p Point) bool {
	if p.y < 0 || p.y >= len(field) {
		return false
	}
	if p.x < 0 || p.x >= len(field[p.y]) {
		return false
	}
	return true
}

func printStack(field [][]int, stack []Point) {
	for _, elem := range stack {
		fmt.Printf("(%d, %d)=%d; ", elem.y, elem.x, field[elem.y][elem.x])
	}
	fmt.Println()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var field [][]int
	var starts []Point
	scanner := bufio.NewScanner(file)

	var y int
	for scanner.Scan() {
		line := scanner.Bytes()

		row := make([]int, 0, len(line))
		for x, b := range line {
			h := int(b - '0')
			if h == 0 {
				starts = append(starts, Point{x, y})
			}
			row = append(row, int(b-'0'))
		}
		field = append(field, row)
		y++
	}

	var total int
	for _, trailhead := range starts {
		var score int
		var stack []Point
		stack = append(stack, trailhead)

		visited := make(map[Point]bool)

		for len(stack) > 0 {
			var top Point
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]

			if field[top.y][top.x] == 9 && !visited[top] {
				score++
				visited[top] = true
				continue
			}

			// add all possible neighbors
			for _, d := range directions {
				next := top.Add(d)
				if !inBounds(field, next) {
					continue
				}
				if field[next.y][next.x]-field[top.y][top.x] != 1 {
					continue
				}
				stack = append(stack, next)
			}
		}
		// fmt.Println(trailhead, score)
		total += score
	}
	fmt.Println(total)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var field [][]int
	var starts []Point
	scanner := bufio.NewScanner(file)

	var y int
	for scanner.Scan() {
		line := scanner.Bytes()

		row := make([]int, 0, len(line))
		for x, b := range line {
			h := int(b - '0')
			if b == '.' {
				h = -1
			}
			if h == 0 {
				starts = append(starts, Point{x, y})
			}
			row = append(row, int(b-'0'))
		}
		field = append(field, row)
		y++
	}

	var total int
	fmt.Println(starts)
	for _, trailhead := range starts {
		var rating int
		var stack []Point
		stack = append(stack, trailhead)

		for len(stack) > 0 {
			var top Point
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]

			if field[top.y][top.x] == 9 {
				rating++
				continue
			}

			// add all possible neighbors
			for _, d := range directions {
				next := top.Add(d)
				if !inBounds(field, next) {
					continue
				}
				if field[next.y][next.x]-field[top.y][top.x] != 1 {
					continue
				}
				stack = append(stack, next)
			}
		}
		fmt.Println(trailhead, rating)
		total += rating
	}
	fmt.Println(total)
}

func main() {
	// part1()
	part2()
}
