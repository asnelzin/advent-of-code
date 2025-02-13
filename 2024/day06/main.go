package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/asnelzin/advent-of-code/pkg/geo"
)

var directions = []geo.Direction{geo.Up, geo.Right, geo.Down, geo.Left}

func inBounds(field [][]byte, p geo.Point) bool {
	if p.Y < 0 || p.Y >= len(field) {
		return false
	}
	if p.X < 0 || p.X >= len(field[p.Y]) {
		return false
	}
	return true
}

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

	p := geo.Point{}
	// find start
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] != '^' {
				continue
			}
			p.X = x
			p.Y = y
		}
	}

	total := 0
	turns := 0
	d := geo.Up
	for inBounds(field, p) {
		if field[p.Y][p.X] == '#' {
			p = p.Sub(d)
			turns++
			d = directions[turns%len(directions)]
			continue
		}

		if field[p.Y][p.X] != 'X' {
			field[p.Y][p.X] = 'X'
			total++
		}
		p = p.Add(d)
	}

	fmt.Println(total)
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

	start := geo.Point{}
	// find start
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] != '^' {
				continue
			}
			start.X = x
			start.Y = y
		}
	}

	turns := 0
	d := geo.Up
	p := start

	boulders := make(map[geo.Point]bool)
	for {
		next := p.Add(d)
		if !inBounds(field, next) {
			break
		}

		for field[next.Y][next.X] == '#' {
			turns++
			d = directions[turns%len(directions)]
			next = p.Add(d)
		}
		p = p.Add(d)

		// try place obstacle at next position
		if hasCycle(field, start, p) {
			boulders[p] = true
		}
	}

	fmt.Println(len(boulders))
}

func hasCycle(field [][]byte, start geo.Point, b geo.Point) bool {
	if !inBounds(field, b) || field[b.Y][b.X] != '.' {
		return false
	}

	cf := copyField(field)
	cf[b.Y][b.X] = 'O'

	turns := 0
	p := start
	d := geo.Up

	type pair struct {
		p geo.Point
		d geo.Direction
	}

	visited := make(map[pair]int)
	for inBounds(cf, p) {
		if visited[pair{p, d}] >= 1 {
			return true
		}
		visited[pair{p, d}] += 1

		next := p.Add(d)
		if !inBounds(cf, next) {
			break
		}
		for cf[next.Y][next.X] == '#' || cf[next.Y][next.X] == 'O' {
			turns++
			d = directions[turns%len(directions)]
			next = p.Add(d)
		}
		p = p.Add(d)
	}
	return false
}

func copyField(field [][]byte) [][]byte {
	c := make([][]byte, len(field))
	for i := range field {
		c[i] = make([]byte, len(field[i]))
		copy(c[i], field[i])
	}
	return c
}

func printField(field [][]byte) {
	for _, line := range field {
		for _, ch := range line {
			fmt.Printf("%c", ch)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func main() {
	part2()
}
