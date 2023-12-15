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

func printWorld(world [][]byte) {
	for _, row := range world {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func manhattanDistance(p1, p2 Pair) int {
	// https://en.wikipedia.org/wiki/Taxicab_geometry
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
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
	occupiedRows := make(map[int]bool)
	occupiedCols := make(map[int]bool)
	for i := range world {
		for j := range world[i] {
			if world[i][j] == '#' {
				occupiedRows[i] = true
				occupiedCols[j] = true
			}
		}
	}

	var missingRows, missingCols []int
	for i := range world {
		if !occupiedRows[i] {
			missingRows = append(missingRows, i)
		}
	}
	for j := range world[0] {
		if !occupiedCols[j] {
			missingCols = append(missingCols, j)
		}
	}

	shift := 0
	for _, i := range missingRows {
		row := func() [][]byte {
			r := make([][]byte, 1)
			r[0] = make([]byte, len(world[0]))
			for j := range r[0] {
				r[0][j] = '.'
			}
			return r
		}()
		i = i + shift
		world = append(world[:i], append(row, world[i:]...)...)
		shift++
	}

	shift = 0
	for i := 0; i < len(world); i++ {
		for _, j := range missingCols {
			j = j + shift
			world[i] = append(world[i][:j], append([]byte{'.'}, world[i][j:]...)...)
			shift++
		}
		shift = 0
	}

	var galaxies []Pair
	for i := range world {
		for j := range world[i] {
			if world[i][j] == '#' {
				galaxies = append(galaxies, Pair{i, j})
			}
		}
	}

	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			d := manhattanDistance(galaxies[i], galaxies[j])
			total += d
		}
	}
	fmt.Println(total)
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
	occupiedRows := make(map[int]bool)
	occupiedCols := make(map[int]bool)
	var galaxies []Pair
	for i := range world {
		for j := range world[i] {
			if world[i][j] == '#' {
				occupiedRows[i] = true
				occupiedCols[j] = true
				galaxies = append(galaxies, Pair{j, i})
			}
		}
	}

	total := 0
	factor := 1000000

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			shiftX, shiftY := 0, 0
			lo := min(galaxies[i].x, galaxies[j].x)
			hi := max(galaxies[i].x, galaxies[j].x)
			for x := lo + 1; x < hi; x++ {
				if !occupiedCols[x] {
					shiftX++
				}
			}

			lo = min(galaxies[i].y, galaxies[j].y)
			hi = max(galaxies[i].y, galaxies[j].y)
			for y := lo + 1; y < hi; y++ {
				if !occupiedRows[y] {
					shiftY++
				}
			}

			shiftX *= factor - 1
			shiftY *= factor - 1

			d := manhattanDistance(galaxies[i], galaxies[j]) + shiftX + shiftY
			total += d
		}
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
