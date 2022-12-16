package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type point struct {
	x, y int
}

func findShortestPath(world [][]rune, start point, end point) int {
	steps := make([][]int, len(world))
	for i := 0; i < len(world); i++ {
		steps[i] = make([]int, len(world[0]))
	}

	reachable := func(p point) []point {
		r := make([]point, 0, 4)
		left := point{p.x - 1, p.y}
		right := point{p.x + 1, p.y}
		up := point{p.x, p.y - 1}
		down := point{p.x, p.y + 1}

		for _, dir := range []point{left, right, up, down} {
			// check constraints
			if dir.x < 0 || dir.x >= len(world[0]) {
				continue
			}

			if dir.y < 0 || dir.y >= len(world) {
				continue
			}

			// not yet visited
			if steps[dir.y][dir.x] != 0 {
				continue
			}

			// check if height is at most one higher than current
			if world[dir.y][dir.x]-world[p.y][p.x] > 1 {
				continue
			}

			r = append(r, dir)
		}

		return r
	}
	q := []point{start}
	steps[start.y][start.x] = 1

	for len(q) > 0 {
		var v point
		v, q = q[0], q[1:]
		for _, w := range reachable(v) {
			steps[w.y][w.x] = steps[v.y][v.x] + 1
			q = append(q, w)
		}
	}

	return steps[end.y][end.x] - 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		world          [][]rune
		start          point
		end            point
		possibleStarts []point
	)

	for scanner.Scan() {
		world = append(world, []rune{})
		for i, c := range scanner.Text() {
			if c == 'S' {
				start.x = i
				start.y = len(world) - 1
				c = 'a'
			}
			if c == 'E' {
				end.x = i
				end.y = len(world) - 1
				c = 'z'
			}

			if c == 'a' {
				possibleStarts = append(possibleStarts, point{i, len(world) - 1})
			}

			world[len(world)-1] = append(world[len(world)-1], c)
		}
	}

	// part 1
	fmt.Println(findShortestPath(world, start, end))

	// part 2
	var lengths []int
	for _, st := range possibleStarts {
		l := findShortestPath(world, st, end)
		if l == -1 {
			continue
		}
		lengths = append(lengths, l)
	}

	sort.Ints(lengths)
	fmt.Println(lengths[0])
}
