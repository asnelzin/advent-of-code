package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type hand int

const (
	rock hand = iota
	paper
	scissor
)

type outcome int

const (
	lose outcome = iota + 10
	draw
	win
)

var shapeScore = map[hand]int{
	rock:    1,
	paper:   2,
	scissor: 3,
}

var outcomeScore = map[outcome]int{
	lose: 0,
	draw: 3,
	win:  6,
}

func score(opponent hand, need outcome) int {
	score := outcomeScore[need]
	if need == draw {
		return score + shapeScore[opponent]
	}

	switch opponent {
	case rock:
		switch need {
		case lose:
			score += shapeScore[scissor]
		case win:
			score += shapeScore[paper]
		}
	case paper:
		switch need {
		case lose:
			score += shapeScore[rock]
		case win:
			score += shapeScore[scissor]
		}
	case scissor:
		switch need {
		case lose:
			score += shapeScore[paper]
		case win:
			score += shapeScore[rock]
		}
	}
	return score
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			first  rune
			second rune
		)
		fmt.Sscanf(scanner.Text(), "%c %c", &first, &second)

		var opponent hand
		switch first {
		case 'A':
			opponent = rock
		case 'B':
			opponent = paper
		case 'C':
			opponent = scissor
		}

		var need outcome
		switch second {
		case 'X':
			need = lose
		case 'Y':
			need = draw
		case 'Z':
			need = win
		}

		total += score(opponent, need)
	}

	fmt.Println(total)
}
