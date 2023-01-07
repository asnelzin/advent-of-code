package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	data := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	numOfBits := len(data[0])
	sums := make([]int, numOfBits)
	for _, s := range data {
		for i, bit := range s {
			// int('0') == 48
			sums[i] += int(bit) - 48
		}
	}

	totalNumbers := len(data)
	gamma := strings.Builder{}
	epsilon := strings.Builder{}
	for _, sum := range sums {
		if sum > totalNumbers/2 {
			gamma.WriteRune('1')
			epsilon.WriteRune('0')
		} else {
			gamma.WriteRune('0')
			epsilon.WriteRune('1')
		}
	}

	gammaN, _ := strconv.ParseInt(gamma.String(), 2, 32)
	epsilonN, _ := strconv.ParseInt(epsilon.String(), 2, 32)
	fmt.Println(gammaN * epsilonN)
}

func part2() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	}
}

func main() {

	part1()
	// part2()
}
