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

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var forest [][]int
	for scanner.Scan() {
		var row []int
		for _, tree := range scanner.Text() {
			row = append(row, int(tree)-int('0'))
		}
		forest = append(forest, row)
	}

	visible := map[point]bool{}

	// from left and top
	for i := 0; i < len(forest); i++ {
		maxFromLeft := -1
		maxFromTop := -1
		for j := 0; j < len(forest); j++ {
			if forest[i][j] > maxFromLeft {
				maxFromLeft = forest[i][j]
				visible[point{i, j}] = true
			}
			if forest[j][i] > maxFromTop {
				maxFromTop = forest[j][i]
				visible[point{j, i}] = true
			}
		}
	}

	// from right and bottom
	for i := 0; i < len(forest); i++ {
		maxFromRight := -1
		maxFromBottom := -1
		for j := len(forest) - 1; j >= 0; j-- {
			if forest[i][j] > maxFromRight {
				maxFromRight = forest[i][j]
				visible[point{i, j}] = true
			}
			if forest[j][i] > maxFromBottom {
				maxFromBottom = forest[j][i]
				visible[point{j, i}] = true
			}
		}
	}
	fmt.Println(len(visible))

	var scores []int
	for t := range visible {
		scores = append(scores, score(forest, t.x, t.y))
	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)-1])
}

func score(forest [][]int, i int, j int) int {
	dest := []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	score := 1
	for _, d := range dest {
		distance := 0
		ii := i
		jj := j
		for {
			ii += d.x
			jj += d.y
			if ii < 0 || jj < 0 || ii >= len(forest) || jj >= len(forest) {
				break
			}

			distance++
			if forest[ii][jj] >= forest[i][j] {
				break
			}
		}

		score *= distance
	}
	return score
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var forest [][]int
	for scanner.Scan() {
		var row []int
		for _, tree := range scanner.Text() {
			row = append(row, int(tree)-int('0'))
		}
		forest = append(forest, row)
	}

	var scores []int

	// do not bother calculating scores for outer trees
	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest)-1; j++ {
			scores = append(scores, score(forest, i, j))
		}
	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)-1])
}

func main() {
	part1()
	part2()
}
