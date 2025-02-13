package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/asnelzin/advent-of-code/pkg/geo"
)

func main() {
	part2()
}

func inBounds(field [][]byte, p geo.Point) bool {
	return p.Y >= 0 && p.Y < len(field) && p.X >= 0 && p.X < len(field[p.Y])
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var field [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []byte(line))
	}

	antennas := make(map[byte][]geo.Point)
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] == '.' {
				continue
			}
			name := field[y][x]
			if antennas[name] == nil {
				antennas[name] = make([]geo.Point, 0)
			}
			antennas[name] = append(antennas[name], geo.Point{x, y})
		}
	}

	fmt.Println(antennas)

	antinodes := make(map[geo.Point]bool)
	for name, points := range antennas {
		fmt.Println(name)

		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {

				k := float64(points[j].Y-points[i].Y) / float64(points[j].X-points[i].X)
				b := float64(points[i].Y) - (k * float64(points[i].X))

				distance := points[i].X - points[j].X

				x1 := float64(points[i].X + distance)
				y1 := x1*k + b
				p1 := geo.Point{int(math.Round(x1)), int(math.Round(y1))}
				if inBounds(field, p1) {
					antinodes[p1] = true
				}

				x2 := float64(points[j].X - distance)
				y2 := x2*k + b

				p2 := geo.Point{int(math.Round(x2)), int(math.Round(y2))}
				if inBounds(field, p2) {
					antinodes[p2] = true
				}
			}
		}
	}
	fmt.Println(len(antinodes))

	for k, _ := range antinodes {
		field[k.Y][k.X] = '#'
	}
	printField(field)
}

func printField(field [][]byte) {
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			fmt.Print(string(field[y][x]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var field [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []byte(line))
	}

	antennas := make(map[byte][]geo.Point)
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] == '.' {
				continue
			}
			name := field[y][x]
			if antennas[name] == nil {
				antennas[name] = make([]geo.Point, 0)
			}
			antennas[name] = append(antennas[name], geo.Point{x, y})
		}
	}

	fmt.Println(antennas)

	antinodes := make(map[geo.Point]bool)
	for name, points := range antennas {
		fmt.Println(name)

		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {

				k := float64(points[j].Y-points[i].Y) / float64(points[j].X-points[i].X)
				b := float64(points[i].Y) - (k * float64(points[i].X))

				distance := points[i].X - points[j].X

				p1 := points[i]
				for inBounds(field, p1) {
					fmt.Println(p1)
					antinodes[p1] = true
					x1 := float64(p1.X + distance)
					y1 := x1*k + b
					p1 = geo.Point{X: int(math.Round(x1)), Y: int(math.Round(y1))}
				}

				p2 := points[j]
				for inBounds(field, p2) {
					antinodes[p2] = true
					x2 := float64(p2.X - distance)
					y2 := x2*k + b
					p2 = geo.Point{X: int(math.Round(x2)), Y: int(math.Round(y2))}
				}
			}
		}
	}
	fmt.Println(len(antinodes))

	for k, _ := range antinodes {
		field[k.Y][k.X] = '#'
	}
	printField(field)
}
