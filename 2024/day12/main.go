package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/asnelzin/advent-of-code/pkg/geo"
)

var directions = []geo.Direction{geo.Left, geo.Right, geo.Down, geo.Up}

func countSides(field [][]byte, p geo.Point) int {
	sides := 0
	for _, d := range directions {
		n := p.Add(d)
		if !geo.InBounds2D(field, n) {
			sides++
			continue
		}

		if geo.At(field, p) != geo.At(field, n) {
			sides++
			continue
		}
	}
	return sides
}

type k struct {
	direction  geo.Direction
	coordinate int
}

func openSides(field [][]byte, p geo.Point) []geo.Direction {
	var sides []geo.Direction
	for _, d := range directions {
		n := p.Add(d)
		if !geo.InBounds2D(field, n) {
			sides = append(sides, d)
			continue
		}

		if geo.At(field, p) != geo.At(field, n) {
			sides = append(sides, d)
			continue
		}
	}
	return sides
}

func printSides(sides map[k][]int) {
	for key, v := range sides {
		fmt.Printf("Side(%s, %d): %v; ", key.direction, key.coordinate, v)
	}
	fmt.Println()
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

	visited := make(map[geo.Point]bool)
	price := 0
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			at := geo.Point{x, y}
			if _, ok := visited[at]; ok {
				continue
			}

			var perimeter, area int

			var top geo.Point
			stack := []geo.Point{at}
			for len(stack) > 0 {
				top, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if _, ok := visited[top]; ok {
					continue
				}

				visited[top] = true
				area += 1
				perimeter += countSides(field, top)

				for _, d := range directions {
					n := top.Add(d)
					if !geo.InBounds2D(field, n) {
						continue
					}

					if geo.At(field, n) != geo.At(field, top) {
						continue
					}

					stack = append(stack, n)
				}
			}

			fmt.Println(string(geo.At(field, at)), area, perimeter)
			price += area * perimeter
		}
	}
	fmt.Println(price)
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

	visited := make(map[geo.Point]bool)
	price := 0
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			at := geo.Point{x, y}
			if _, ok := visited[at]; ok {
				continue
			}

			var area int

			sides := make(map[k][]int)

			var top geo.Point
			stack := []geo.Point{at}
			for len(stack) > 0 {
				top, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if _, ok := visited[top]; ok {
					continue
				}

				visited[top] = true
				area += 1
				open := openSides(field, top)
				for _, d := range open {
					var key k
					key.direction = d
					if d == geo.Up || d == geo.Down {
						key.coordinate = top.Y
						sides[key] = append(sides[key], top.X)
					} else {
						key.coordinate = top.X
						sides[key] = append(sides[key], top.Y)
					}
				}

				for _, d := range directions {
					n := top.Add(d)
					if !geo.InBounds2D(field, n) {
						continue
					}

					if geo.At(field, n) != geo.At(field, top) {
						continue
					}

					stack = append(stack, n)
				}
			}

			totalSides := 0
			for key, _ := range sides {
				slices.Sort(sides[key])

				for i := 0; i < len(sides[key])-1; i++ {
					if sides[key][i+1]-sides[key][i] == 1 {
						continue
					}
					totalSides += 1
				}
				totalSides += 1
			}

			printSides(sides)
			fmt.Println(string(geo.At(field, at)), area, totalSides)

			price += area * totalSides
		}
	}
	fmt.Println(price)
}

func main() {
	// part1()
	part2()
}
