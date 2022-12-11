package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func sign(x int) int {
	if x == 0 {
		return 0
	}

	if x < 0 {
		return -1
	}
	return 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	x, y int
}

func (p Point) DistanceSquared(other Point) float64 {
	return math.Pow(float64(p.x-other.x), 2) + math.Pow(float64(p.y-other.y), 2)
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Direction string

const (
	right Direction = "R"
	left            = "L"
	up              = "U"
	down            = "D"
)

func adjust(head Point, tail Point) Point {
	d := head.DistanceSquared(tail)
	if d < 4 {
		return tail
	}

	dx, dy := head.x-tail.x, head.y-tail.y
	t := Point{tail.x, tail.y}

	if dx == 0 || dy == 0 {
		t.y += sign(dy)
		t.x += sign(dx)
		return t
	}

	if abs(dx) == abs(dy) {
		t.x += dx / 2
		t.y += dy / 2
		return t
	}

	if (abs(dx) == 1 && abs(dy) == 2) || (abs(dx) == 2 && abs(dy) == 1) {
		t.x += sign(dx)
		t.y += sign(dy)
	}

	return t
}

type Tail interface {
	Adjust(head Point)
	Last() Point
	String() string
}

type SingleTail struct {
	tail Point
}

func (t *SingleTail) Adjust(head Point) {
	t.tail = adjust(head, t.tail)
}

func (t *SingleTail) Last() Point {
	return t.tail
}

func (t *SingleTail) String() string {
	return fmt.Sprint(t.tail)
}

type MultipleTail struct {
	tails []Point
}

func NewMultipleTails(length int) *MultipleTail {
	knots := make([]Point, length)
	return &MultipleTail{knots}
}

func (t *MultipleTail) Adjust(head Point) {
	current := head
	for i := 0; i < len(t.tails); i++ {
		t.tails[i] = adjust(current, t.tails[i])
		current = t.tails[i]
	}
}

func (t *MultipleTail) Last() Point {
	return t.tails[len(t.tails)-1]
}

func (t *MultipleTail) String() string {
	b := strings.Builder{}
	for _, tail := range t.tails {
		b.WriteString(fmt.Sprintf("%s-", tail))
	}
	return b.String()[:b.Len()-1]
}

type Rope struct {
	head Point
	Tail
}

func (r Rope) String() string {
	return fmt.Sprintf("head - %s, tail - %s", r.head, r.Tail)
}

func (r *Rope) Move(d Direction) {
	switch d {
	case down:
		r.head.y -= 1
	case up:
		r.head.y += 1
	case left:
		r.head.x -= 1
	case right:
		r.head.x += 1
	}
	r.Adjust(r.head)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	rope := Rope{Point{}, &SingleTail{}}
	largeRope := Rope{Point{}, NewMultipleTails(9)}
	pos1 := map[Point]bool{{0, 0}: true}
	pos2 := map[Point]bool{{0, 0}: true}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if len(args) < 2 {
			log.Print("input line does not have Direction and/or distance")
			continue
		}
		cmd := args[0]
		distance, err := strconv.Atoi(args[1])
		if err != nil {
			log.Printf("could not parse distance `%s`: %v", args[1], err)
			continue
		}
		direction := Direction(cmd)
		for i := 0; i < distance; i++ {
			rope.Move(direction)
			largeRope.Move(direction)
			pos1[rope.Last()] = true
			pos2[largeRope.Last()] = true
		}
	}

	fmt.Printf("part1 := %d\n", len(pos1))
	fmt.Printf("part2 := %d\n", len(pos2))
}
