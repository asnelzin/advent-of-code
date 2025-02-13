package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/asnelzin/advent-of-code/pkg/geo"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{
		R: 0,
		G: 100,
		B: 0,
		A: 255,
	}
)

type Velocity struct {
	X, Y int
}

type Robot struct {
	start geo.Point
	v     Velocity
}

func part1() {
	const (
		seconds = 100
		tall    = 103
		wide    = 101
		// seconds = 100
		// tall    = 7
		// wide    = 11
	)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var robots []Robot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var start geo.Point
		var v Velocity

		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			log.Fatalf("could not parse input line: %v", line)
		}

		if len(parts[0]) < 4 {
			log.Fatalf("could not parse starting point: %v", parts[0])
		}
		if len(parts[1]) < 4 {
			log.Fatalf("could not parse velocity: %v", parts[1])
		}

		startParts := strings.Split(parts[0][2:], ",")
		x, err := strconv.Atoi(startParts[0])
		if err != nil {
			log.Fatalf("could not parse start X (%v): %v", startParts[0], err)
		}
		y, err := strconv.Atoi(startParts[1])
		if err != nil {
			log.Fatalf("could not parse start Y (%v): %v", startParts[1], err)
		}

		start.X = x
		start.Y = y

		velParts := strings.Split(parts[1][2:], ",")
		x, err = strconv.Atoi(velParts[0])
		if err != nil {
			log.Fatalf("could not parse velocity X (%v): %v", velParts[0], err)
		}
		y, err = strconv.Atoi(velParts[1])
		if err != nil {
			log.Fatalf("could not parse velocity Y (%v): %v", velParts[1], err)
		}

		v.X = x
		v.Y = y

		robots = append(robots, Robot{start, v})
	}

	finish := make(map[geo.Point]int)

	for _, r := range robots {
		var cur geo.Point
		cur.X = r.start.X
		cur.Y = r.start.Y

		cur.X = (r.start.X + (r.v.X * seconds)) % wide
		if cur.X < 0 {
			cur.X = wide + cur.X
		}
		cur.Y = (r.start.Y + (r.v.Y * seconds)) % tall
		if cur.Y < 0 {
			cur.Y = tall + cur.Y
		}

		finish[cur] += 1
	}

	mx := (wide - 1) / 2
	my := (tall - 1) / 2

	var quadrants [4]int
	for y := 0; y < tall; y++ {
		for x := 0; x < wide; x++ {
			p := geo.Point{x, y}
			count, ok := finish[p]
			if !ok {
				continue
			}
			if x < mx && y < my {
				quadrants[0] += count
			}
			if x < mx && y > my {
				quadrants[1] += count
			}
			if x > mx && y < my {
				quadrants[2] += count
			}
			if x > mx && y > my {
				quadrants[3] += count
			}
		}
	}
	total := 1
	for _, q := range quadrants {
		total *= q
	}
	fmt.Println(total)
}

func part2() {
	const (
		seconds = 100
		tall    = 103
		wide    = 101
		// seconds = 100
		// tall    = 7
		// wide    = 11
	)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var robots []Robot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var start geo.Point
		var v Velocity

		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			log.Fatalf("could not parse input line: %v", line)
		}

		if len(parts[0]) < 4 {
			log.Fatalf("could not parse starting point: %v", parts[0])
		}
		if len(parts[1]) < 4 {
			log.Fatalf("could not parse velocity: %v", parts[1])
		}

		startParts := strings.Split(parts[0][2:], ",")
		x, err := strconv.Atoi(startParts[0])
		if err != nil {
			log.Fatalf("could not parse start X (%v): %v", startParts[0], err)
		}
		y, err := strconv.Atoi(startParts[1])
		if err != nil {
			log.Fatalf("could not parse start Y (%v): %v", startParts[1], err)
		}

		start.X = x
		start.Y = y

		velParts := strings.Split(parts[1][2:], ",")
		x, err = strconv.Atoi(velParts[0])
		if err != nil {
			log.Fatalf("could not parse velocity X (%v): %v", velParts[0], err)
		}
		y, err = strconv.Atoi(velParts[1])
		if err != nil {
			log.Fatalf("could not parse velocity Y (%v): %v", velParts[1], err)
		}

		v.X = x
		v.Y = y

		robots = append(robots, Robot{start, v})
	}

	// Step 1: find a "strange" image: with weird robots' positions

	fillFrom := 0
	padding := 5 // px
	totalHeight := (tall + padding) * 100
	img := generateImage(totalHeight, wide)
	for sec := 1; sec <= seconds; sec++ {
		positions := make(map[geo.Point]int)
		for _, r := range robots {
			var cur geo.Point
			cur.X = r.start.X
			cur.Y = r.start.Y

			cur.X = (r.start.X + (r.v.X * sec)) % wide
			if cur.X < 0 {
				cur.X = wide + cur.X
			}
			cur.Y = (r.start.Y + (r.v.Y * sec)) % tall
			if cur.Y < 0 {
				cur.Y = tall + cur.Y
			}

			positions[cur] += 1
		}

		if sec > 1 {
			fillPadding(img, padding, fillFrom)
			fillFrom += padding
		}
		appendToImage(img, positions, fillFrom)
		fillFrom += tall
	}
	fillPadding(img, padding, fillFrom)
	saveImage(img, "map.png")

	// weird vertical lines appeared at 99th second
	// this position will repeat every 101 second from this point (101 is a wide)

	// Step 2: Each weird position should repeat every 101 seconds
	strangeSec := 99

	fillFrom = 0
	totalHeight = (tall + padding) * 100
	img2 := generateImage(totalHeight, wide)
	for i := 0; i < 100; i++ {
		sec := strangeSec + (wide * i)
		positions := make(map[geo.Point]int)
		for _, r := range robots {
			var cur geo.Point
			cur.X = r.start.X
			cur.Y = r.start.Y

			cur.X = (r.start.X + (r.v.X * sec)) % wide
			if cur.X < 0 {
				cur.X = wide + cur.X
			}
			cur.Y = (r.start.Y + (r.v.Y * sec)) % tall
			if cur.Y < 0 {
				cur.Y = tall + cur.Y
			}

			positions[cur] += 1
		}

		if i > 0 {
			fillPadding(img2, padding, fillFrom)
			fillFrom += padding
		}
		appendToImage(img2, positions, fillFrom)

		found := true
		for _, c := range positions {
			if c != 1 {
				found = false
				break
			}
		}
		if found {
			fmt.Println("Answer: ", sec) // 8280
		}

		fillFrom += tall
	}
	fillPadding(img2, padding, fillFrom)
	saveImage(img2, "repeated.png")
}

func generateImage(h, w int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < img.Rect.Max.Y; y++ {
		for x := 0; x < img.Rect.Max.X; x++ {
			img.Set(x, y, white)
		}
	}
	return img
}

func appendToImage(img *image.RGBA, points map[geo.Point]int, minY int) {
	for p, count := range points {
		clr := adjustSaturationRGBA(green, float64(count+1))
		if count == 1 {
			clr = black
		}
		img.Set(p.X, p.Y+minY, clr)
	}
}

func fillPadding(img *image.RGBA, size int, startY int) {
	for y := startY; y < startY+size; y++ {
		for x := 0; x < img.Rect.Max.X; x++ {
			img.Set(x, y, red)
		}
	}
}

func adjustSaturationRGBA(c color.RGBA, factor float64) color.RGBA {
	// Calculate the average intensity of the color
	avg := float64(c.R+c.G+c.B) / 3.0

	// Adjust each RGB component
	adjust := func(value uint8) uint8 {
		adjusted := int(math.Round(avg + (float64(value)-avg)*factor))
		if adjusted < 0 {
			return 0
		} else if adjusted > 255 {
			return 255
		}
		return uint8(adjusted)
	}

	// Apply adjustment
	return color.RGBA{
		R: adjust(c.R),
		G: adjust(c.G),
		B: adjust(c.B),
		A: c.A, // Keep alpha unchanged
	}
}

func saveImage(img image.Image, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("could not open map file: %v", err)
	}

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("could not write to opened map file: %v", err)
	}
}

func main() {
	// part1()
	part2()
}
