package geo

import "fmt"

type Direction struct {
	X, Y int
}

var (
	Left  = Direction{-1, 0}
	Right = Direction{1, 0}
	Down  = Direction{0, 1}
	Up    = Direction{0, -1}
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return "Unknown"
	}
}

type Point struct {
	X, Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%d, %d)", p.X, p.Y)
}

func (p *Point) Add(d Direction) Point {
	return Point{p.X + d.X, p.Y + d.Y}
}

func (p *Point) Sub(d Direction) Point {
	return Point{p.X - d.X, p.Y - d.Y}
}

func InBounds2D[T any](field [][]T, p Point) bool {
	if p.Y < 0 || p.Y >= len(field) {
		return false
	}
	if p.X < 0 || p.X >= len(field[p.Y]) {
		return false
	}
	return true
}

func At[T any](field [][]T, p Point) T {
	return field[p.Y][p.X]
}

func PrintFieldBytes(field [][]byte) {
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			fmt.Printf("%s", string(field[y][x]))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}
