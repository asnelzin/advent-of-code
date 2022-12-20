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

func sand(world [][]string, start point) point {
	checkAbyss := func(p point) bool {
		if p.y < 0 || p.y >= len(world) {
			return true
		}
		if p.x < 0 || p.x >= len(world[p.y]) {
			return true
		}
		return false
	}

	check := func(p point) bool {
		if world[p.y][p.x] == "#" || world[p.y][p.x] == "o" {
			return false
		}
		return true
	}

	current := start

loop:
	for {
		down := point{current.x, current.y + 1}
		downLeft := point{current.x - 1, current.y + 1}
		downRight := point{current.x + 1, current.y + 1}
		for _, d := range []point{down, downLeft, downRight} {
			if checkAbyss(d) {
				return point{-1, -1}
			}
		}

		switch {
		case check(down):
			current = down
		case check(downLeft):
			current = downLeft
		case check(downRight):
			current = downRight
		default:
			break loop
		}
	}

	return current
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}

	var rocks []segment

	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := 0, 0
	for _, line := range strings.Split(string(data), "\n") {
		pp := strings.Split(line, " -> ")
		var rock segment
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

			rock = append(rock, point{x, y})
		}

		rocks = append(rocks, rock)
	}

	fmt.Println(minX, maxX)
	fmt.Println(minY, maxY)

	world := make([][]string, maxY+1)
	for i := 0; i < maxY+1; i++ {
		world[i] = make([]string, maxX-minX+1)
	}

	for _, r := range rocks {
		start := r[0]
		for _, p := range r[1:] {
			current := start
			step := point{cmp(start.x, p.x), cmp(start.y, p.y)}
			for current != p {
				world[current.y][current.x-minX] = "#"
				current.Add(step)
			}
			world[p.y][p.x-minX] = "#"

			start = current
		}
	}
	world[0][500-minX] = "+"

	printWorld(world)

	grains := 0
	for ; ; grains++ {
		rest := sand(world, point{500 - minX, 0})
		if rest.x == -1 && rest.y == -1 {
			break
		}
		world[rest.y][rest.x] = "o"
	}

	fmt.Println(grains)
	printWorld(world)

}
