package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/asnelzin/advent-of-code/pkg/geo"
)

var (
	buttonRowRe = regexp.MustCompile(`^Button ([AB]): X\+(\d+), Y\+(\d+)$`)
	priceRowRe  = regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)$`)
)

type Button struct {
	dx int
	dy int
}

type Machine struct {
	buttons [2]Button
	prize   geo.Point
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var buttons []Button
	var prizes []geo.Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "Button"):
			parsedButton := buttonRowRe.FindStringSubmatch(line)
			if len(parsedButton) < 4 {
				log.Fatalf("could not parse button row: %s", line)
			}
			dx, err := strconv.Atoi(parsedButton[2])
			if err != nil {
				log.Fatalf("could not parse button: %v", err)
			}
			dy, err := strconv.Atoi(parsedButton[3])
			if err != nil {
				log.Fatalf("could not parse button: %v", err)
			}

			b := Button{
				dx: dx,
				dy: dy,
			}
			buttons = append(buttons, b)

		case strings.HasPrefix(line, "Prize"):
			parsedPrize := priceRowRe.FindStringSubmatch(line)
			if len(parsedPrize) < 3 {
				log.Fatalf("could not parse prize row: %s", line)
			}
			x, err := strconv.Atoi(parsedPrize[1])
			if err != nil {
				log.Fatalf("could not parse prize: %v", err)
			}
			y, err := strconv.Atoi(parsedPrize[2])
			if err != nil {
				log.Fatalf("could not parse prize: %v", err)
			}

			prizes = append(prizes, geo.Point{X: x, Y: y})

		default:
			continue
		}
	}

	machines := make([]Machine, 0, len(prizes))
	for i := 1; i < len(buttons); i += 2 {
		m := Machine{}
		m.buttons = [2]Button{buttons[i-1], buttons[i]}
		m.prize = prizes[(i-1)/2]
		machines = append(machines, m)
	}
	fmt.Println(machines)

	totalCost := 0
	for _, m := range machines {
		found := false
		cheapest := 400
		buttonA := m.buttons[0]
		buttonB := m.buttons[1]

		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				pos := geo.Point{
					X: a*buttonA.dx + b*buttonB.dx,
					Y: a*buttonA.dy + b*buttonB.dy,
				}
				if pos == m.prize {
					fmt.Println(a, b)
					found = true
					cheapest = min(cheapest, a*3+b)
				}
			}
		}

		if found {
			totalCost += cheapest
			fmt.Println(cheapest)
		}
	}

	fmt.Println(totalCost)
}

func part2() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var buttons []Button
	var prizes []geo.Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "Button"):
			parsedButton := buttonRowRe.FindStringSubmatch(line)
			if len(parsedButton) < 4 {
				log.Fatalf("could not parse button row: %s", line)
			}
			dx, err := strconv.Atoi(parsedButton[2])
			if err != nil {
				log.Fatalf("could not parse button: %v", err)
			}
			dy, err := strconv.Atoi(parsedButton[3])
			if err != nil {
				log.Fatalf("could not parse button: %v", err)
			}

			b := Button{
				dx: dx,
				dy: dy,
			}
			buttons = append(buttons, b)

		case strings.HasPrefix(line, "Prize"):
			parsedPrize := priceRowRe.FindStringSubmatch(line)
			if len(parsedPrize) < 3 {
				log.Fatalf("could not parse prize row: %s", line)
			}
			x, err := strconv.Atoi(parsedPrize[1])
			if err != nil {
				log.Fatalf("could not parse prize: %v", err)
			}
			y, err := strconv.Atoi(parsedPrize[2])
			if err != nil {
				log.Fatalf("could not parse prize: %v", err)
			}

			prizes = append(prizes, geo.Point{X: x, Y: y})

		default:
			continue
		}
	}

	machines := make([]Machine, 0, len(prizes))
	for i := 1; i < len(buttons); i += 2 {
		m := Machine{}
		m.buttons = [2]Button{buttons[i-1], buttons[i]}
		m.prize = prizes[(i-1)/2]
		// part 2 tweak
		m.prize.X += 10000000000000
		m.prize.Y += 10000000000000
		machines = append(machines, m)
	}

	totalCost := 0
	for _, m := range machines {
		// solving a system of two linear equations
		buttonA := m.buttons[0]
		buttonB := m.buttons[1]
		py := m.prize.Y
		px := m.prize.X
		adx := buttonA.dx
		ady := buttonA.dy
		bdx := buttonB.dx
		bdy := buttonB.dy

		pressedB := (py*adx - px*ady) / (bdy*adx - bdx*ady)
		pressedA := (px - pressedB*bdx) / adx
		if px == pressedA*adx+pressedB*bdx && py == pressedA*ady+pressedB*bdy {
			totalCost += pressedA*3 + pressedB
		}
	}
	fmt.Println(totalCost)

}

func main() {
	// part1()
	part2()
}
