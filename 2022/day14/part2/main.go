package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func cmp(a, b int) int {
	d := b - a
	if d < 0 {
		return -1
	} else if d > 0 {
		return 1
	}
	return 0
}

type point struct {
	x, y int
}

func (p *point) Add(other point) {
	p.x += other.x
	p.y += other.y
}

type segment []point

func printWorld(w [][]string) {
	for i := 0; i < len(w); i++ {
		for j := 0; j < len(w[i]); j++ {
			if w[i][j] == "" {
				fmt.Print(".")
				continue
			}
			fmt.Print(w[i][j])
		}
		fmt.Print("\n")
	}
}

func sand(rocks map[point]bool, rested map[point]bool, floorY int, start point) (int, point) {
	check := func(p point) bool {
		if p.y == floorY {
			return false
		}
		if _, ok := rocks[p]; ok {
			return false
		}
		if _, ok := rested[p]; ok {
			return false
		}
		return true
	}

	n := 0
	current := start

loop:
	for ; ; n++ {
		down := point{current.x, current.y + 1}
		downLeft := point{current.x - 1, current.y + 1}
		downRight := point{current.x + 1, current.y + 1}
		switch {
		case check(down):
			current = down
		case check(downLeft):
			current = downLeft
		case check(downRight):
			current = downRight
		default:
			// can't go anywhere, come to rest
			break loop
		}
	}

	return n, current
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}

	var traces []segment

	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := 0, 0
	for _, line := range strings.Split(string(data), "\n") {
		pp := strings.Split(line, " -> ")
		var trace segment
		for _, p := range pp {
			c := strings.Split(p, ",")
			x, err := strconv.Atoi(c[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(c[1])
			if err != nil {
				log.Fatal(err)
			}
			if x > maxX {
				maxX = x
			}
			if x < minX {
				minX = x
			}

			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}

			trace = append(trace, point{x, y})
		}

		traces = append(traces, trace)
	}

	fmt.Println(minX, maxX)
	fmt.Println(minY, maxY)

	rocks := map[point]bool{}
	for _, tr := range traces {
		start := tr[0]
		for _, p := range tr[1:] {
			current := start
			step := point{cmp(start.x, p.x), cmp(start.y, p.y)}
			for current != p {
				rocks[point{current.x, current.y}] = true
				current.Add(step)
			}
			rocks[point{p.x, p.y}] = true

			start = current
		}
	}

	grains := 1
	rested := map[point]bool{}
	for ; ; grains++ {
		n, rest := sand(rocks, rested, maxY+2, point{500, 0})
		rested[rest] = true
		if n == 0 {
			fmt.Println(rest)
			break
		}
	}

	fmt.Println(grains)
}
