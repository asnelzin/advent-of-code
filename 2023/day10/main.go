package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type Pair struct {
	x, y int
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	world := make([][]byte, 0)
	for _, line := range strings.Split(string(buf), "\n") {
		if line == "" {
			continue
		}
		world = append(world, []byte(line))
	}

	start := Pair{0, 0}
	for y, row := range world {
		for x, col := range row {
			if col == 'S' {
				start = Pair{x, y}
			}
		}
	}

	type record struct {
		pos   Pair
		steps int
	}

	shapes := map[byte][2]Pair{
		'|': {{0, 1}, {0, -1}},  // up, down
		'-': {{1, 0}, {-1, 0}},  // left, right
		'L': {{0, -1}, {1, 0}},  // up, right
		'J': {{0, -1}, {-1, 0}}, // up, left
		'7': {{0, 1}, {-1, 0}},  // down, left
		'F': {{0, 1}, {1, 0}},   // down, right
	}

	queue := make([]record, 0)
	visited := make(map[Pair]bool)
	visited[start] = true

	connected := func(pipe Pair, from Pair) bool {
		connections, ok := shapes[world[pipe.y][pipe.x]]
		if !ok {
			return false
		}
		for _, c := range connections {
			if c.x == -from.x && c.y == -from.y {
				return true
			}
		}
		return false
	}

	for _, dir := range []Pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		next := Pair{start.x + dir.x, start.y + dir.y}
		if next.x < 0 || next.x >= len(world[0]) || next.y < 0 || next.y >= len(world) {
			continue
		}
		if connected(next, dir) {
			queue = append(queue, record{next, 1})
		}
	}

	neighbours := func(pos Pair) []Pair {
		neighbours := make([]Pair, 0)
		for _, dir := range shapes[world[pos.y][pos.x]] {
			next := Pair{pos.x + dir.x, pos.y + dir.y}
			if next.x < 0 || next.x >= len(world[0]) || next.y < 0 || next.y >= len(world) {
				continue
			}
			if connected(next, dir) {
				neighbours = append(neighbours, next)
			}
		}
		return neighbours
	}

	maxSteps := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current.pos] {
			continue
		}
		if current.steps > maxSteps {
			maxSteps = current.steps
		}
		visited[current.pos] = true
		for _, neighbour := range neighbours(current.pos) {
			queue = append(queue, record{neighbour, current.steps + 1})
		}
	}
	fmt.Println(maxSteps)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	world := make([][]byte, 0)
	for _, line := range strings.Split(string(buf), "\n") {
		if line == "" {
			continue
		}
		world = append(world, []byte(line))
	}

	start := Pair{0, 0}
	for y, row := range world {
		for x, col := range row {
			if col == 'S' {
				start = Pair{x, y}
			}
		}
	}

	shapes := map[byte][2]Pair{
		'|': {{0, 1}, {0, -1}},  // up, down
		'-': {{1, 0}, {-1, 0}},  // left, right
		'L': {{0, -1}, {1, 0}},  // up, right
		'J': {{0, -1}, {-1, 0}}, // up, left
		'7': {{0, 1}, {-1, 0}},  // down, left
		'F': {{0, 1}, {1, 0}},   // down, right
	}

	connected := func(pipe Pair, from Pair) bool {
		connections, ok := shapes[world[pipe.y][pipe.x]]
		if !ok {
			return false
		}
		for _, c := range connections {
			if c.x == -from.x && c.y == -from.y {
				return true
			}
		}
		return false
	}

	startDirections := make([]Pair, 0)
	for _, dir := range []Pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		next := Pair{start.x + dir.x, start.y + dir.y}
		if next.x < 0 || next.x >= len(world[0]) || next.y < 0 || next.y >= len(world) {
			continue
		}
		if connected(next, dir) {
			startDirections = append(startDirections, dir)
		}
	}
	if len(startDirections) > 2 {
		log.Fatalf("start has more than 2 directions: %v", startDirections)
	}

	// replace start with proper shape
	for shape, dirs := range shapes {
		if (startDirections[0] == dirs[0] && startDirections[1] == dirs[1]) || (startDirections[0] == dirs[1] && startDirections[1] == dirs[0]) {
			world[start.y][start.x] = shape
		}
	}

	neighbours := func(pos Pair) []Pair {
		neighbours := make([]Pair, 0)
		for _, dir := range shapes[world[pos.y][pos.x]] {
			next := Pair{pos.x + dir.x, pos.y + dir.y}
			if next.x < 0 || next.x >= len(world[0]) || next.y < 0 || next.y >= len(world) {
				continue
			}
			if connected(next, dir) {
				neighbours = append(neighbours, next)
			}
		}
		return neighbours
	}

	path := []Pair{start}
	current := start
	for {
		n := neighbours(current)
		for _, neighbour := range n {
			if len(path) > 1 && path[len(path)-2] == neighbour {
				continue
			}
			current = neighbour
			break
		}
		if current == start {
			break
		}
		path = append(path, current)
	}

	zip := func(a, b []Pair) [][2]Pair {
		r := make([][2]Pair, len(a))
		for i, e := range a {
			r[i] = [2]Pair{e, b[i]}
		}
		return r
	}

	// Shoelace formula
	s := 0
	for _, combined := range zip(path[:len(path)-1], path[1:]) {
		s += combined[0].x*combined[1].y - combined[1].x*combined[0].y
	}
	s += path[len(path)-1].x*path[0].y - path[0].x*path[len(path)-1].y

	area := math.Abs(float64(s)) / 2.0

	// Pick's theorem
	fmt.Println(area - float64(len(path)/2.0) + 1)
}

func main() {
	part1()
	part2()
}
