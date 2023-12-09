package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func findNext(line string) int {
	var nums []int
	for _, part := range strings.Split(line, " ") {
		n, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("could not parse line: %v", err)
		}
		nums = append(nums, n)
	}

	var triangle [][]int
	triangle = append(triangle, nums[:])
	at := 0
	for {
		var row []int
		current := triangle[at]
		for i := 1; i < len(current); i++ {
			row = append(row, current[i]-current[i-1])
		}
		//fmt.Println(len(row))
		triangle = append(triangle, row)
		allzeros := true
		for i := 0; i < len(row); i++ {
			if row[i] != 0 {
				allzeros = false
				break
			}
		}
		if allzeros {
			break
		}
		at++
	}

	// add a zero to the end of the last row
	triangle[len(triangle)-1] = append(triangle[len(triangle)-1], 0)

	for i := len(triangle) - 2; i >= 0; i-- {
		tail := len(triangle[i]) - 1
		triangle[i] = append(triangle[i], triangle[i][tail]+triangle[i+1][tail])
	}

	return triangle[0][len(triangle[0])-1]
}

func findPrev(line string) int {
	var nums []int
	for _, part := range strings.Split(line, " ") {
		n, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("could not parse line: %v", err)
		}
		nums = append(nums, n)
	}

	var triangle [][]int
	triangle = append(triangle, nums[:])
	at := 0
	for {
		var row []int
		current := triangle[at]
		for i := 1; i < len(current); i++ {
			row = append(row, current[i]-current[i-1])
		}
		//fmt.Println(len(row))
		triangle = append(triangle, row)
		allzeros := true
		for i := 0; i < len(row); i++ {
			if row[i] != 0 {
				allzeros = false
				break
			}
		}
		if allzeros {
			break
		}
		at++
	}

	// add a zero to the front of the last row
	triangle[len(triangle)-1] = append([]int{0}, triangle[len(triangle)-1]...)

	for i := len(triangle) - 2; i >= 0; i-- {
		triangle[i] = append([]int{triangle[i][0] - triangle[i+1][0]}, triangle[i]...)
	}

	return triangle[0][0]
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

	total := 0
	for _, line := range strings.Split(string(buf), "\n") {
		if line == "" {
			continue
		}
		n := findNext(line)
		total += n
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

	total := 0
	for _, line := range strings.Split(string(buf), "\n") {
		if line == "" {
			continue
		}
		n := findPrev(line)
		total += n
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
