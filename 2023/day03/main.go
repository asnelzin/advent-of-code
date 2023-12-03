package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var scheme [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		scheme = append(scheme, []byte(line))
	}

	var (
		sum    int
		numlen int
		isPart bool
	)
	for i := 0; i < len(scheme); i++ {
		for j := 0; j < len(scheme[i]); j++ {
			for j < len(scheme[i]) && isDigit(scheme[i][j]) {
				numlen++
				// check above and below (2)
				if (i > 0 && scheme[i-1][j] != '.') || (i < len(scheme)-1 && scheme[i+1][j] != '.') {
					isPart = true
				}
				j++
			}
			if numlen == 0 {
				continue
			}

			// check right (3)
			if j < len(scheme[i]) {
				if scheme[i][j] != '.' || (i > 0 && scheme[i-1][j] != '.') || (i < len(scheme)-1 && scheme[i+1][j] != '.') {
					isPart = true
				}
			}
			// check left (3)
			left := j - numlen - 1
			if left >= 0 {
				if scheme[i][left] != '.' || (i > 0 && scheme[i-1][left] != '.') || (i < len(scheme)-1 && scheme[i+1][left] != '.') {
					isPart = true
				}
			}

			// copy number
			var num int
			if isPart {
				for k := 0; k < numlen; k++ {
					num += int(scheme[i][j-numlen+k]-'0') * int(math.Pow10(numlen-k-1))
				}
				sum += num
			}

			numlen = 0
			isPart = false
		}
	}
	fmt.Println(sum)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var scheme [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		scheme = append(scheme, []byte(line))
	}

	type pair struct {
		x, y int
	}

	var (
		numlen int
		gear   pair
	)
	gears := make(map[pair]map[int]bool) // map of gears and their adjacent numbers
	for i := 0; i < len(scheme); i++ {
		for j := 0; j < len(scheme[i]); j++ {
			for j < len(scheme[i]) && isDigit(scheme[i][j]) {
				numlen++
				// check above and below (2)
				if i > 0 && scheme[i-1][j] == '*' {
					gear = pair{i - 1, j}
				}
				if i < len(scheme)-1 && scheme[i+1][j] == '*' {
					gear = pair{i + 1, j}
				}
				j++
			}
			if numlen == 0 {
				continue
			}

			// check right (3)
			if j < len(scheme[i]) {
				if scheme[i][j] == '*' {
					gear = pair{i, j}
				}
				if i > 0 && scheme[i-1][j] == '*' {
					gear = pair{i - 1, j}
				}
				if i < len(scheme)-1 && scheme[i+1][j] == '*' {
					gear = pair{i + 1, j}
				}
			}

			// check left (3)
			left := j - numlen - 1
			if left >= 0 {
				if scheme[i][left] == '*' {
					gear = pair{i, left}
				}
				if i > 0 && scheme[i-1][left] == '*' {
					gear = pair{i - 1, left}
				}
				if i < len(scheme)-1 && scheme[i+1][left] == '*' {
					gear = pair{i + 1, left}
				}
			}

			// copy number
			var num int
			if gear != (pair{}) {
				for k := 0; k < numlen; k++ {
					num += int(scheme[i][j-numlen+k]-'0') * int(math.Pow10(numlen-k-1))
				}
				if gears[gear] == nil {
					gears[gear] = make(map[int]bool)
				}
				gears[gear][num] = true
			}

			numlen = 0
			gear = pair{}
		}
	}

	sum := 0
	for _, nums := range gears {
		if len(nums) < 2 {
			continue
		}

		ration := 1
		for num := range nums {
			ration *= num
		}
		sum += ration
	}
	fmt.Println(sum)
}

func main() {
	part2()
}
