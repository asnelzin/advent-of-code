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

type op func(int, int) int

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	type Eq struct {
		target int
		nums   []int
	}

	var equations []Eq
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			log.Fatalf("could not parse line %s", line)
		}

		value := parts[0]
		target, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("could not parse number (%s): %v", value, err)
		}

		var nums []int
		for _, operand := range strings.Split(strings.TrimSpace(parts[1]), " ") {
			n, err := strconv.Atoi(operand)
			if err != nil {
				log.Fatalf("could not parse operand (%s): %v", operand, err)
			}
			nums = append(nums, n)
		}
		equations = append(equations, Eq{target, nums})
	}

	ops := map[int]op{
		0: func(a, b int) int { return a + b },
		1: func(a, b int) int { return a * b },
	}
	sum := 0
	for _, eq := range equations {
		t := eq.target
		nums := eq.nums
		permutations := int(math.Pow(2, float64(len(nums)-1)))
		for i := 0; i < permutations; i++ {
			acc := nums[0]
			for j, operand := range nums[1:] {
				acc = ops[(i>>j)&1](acc, operand)
			}
			if t == acc {
				sum += t
				fmt.Printf("%d: ", t)
				printOperations(i, nums)
				break
			}
		}
	}
	fmt.Println("PART 1 TOTAL:", sum)
}

func printOperations(seq int, operands []int) {
	symbol := map[int]string{
		0: "+",
		1: "*",
	}

	fmt.Print(operands[0])
	for j, operand := range operands[1:] {
		sym := symbol[(seq>>j)&1]
		fmt.Printf(" %s %d", sym, operand)
	}
	fmt.Println()
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	type Eq struct {
		target int
		nums   []int
	}

	var equations []Eq
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			log.Fatalf("could not parse line %s", line)
		}

		value := parts[0]
		target, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("could not parse number (%s): %v", value, err)
		}

		var nums []int
		for _, operand := range strings.Split(strings.TrimSpace(parts[1]), " ") {
			n, err := strconv.Atoi(operand)
			if err != nil {
				log.Fatalf("could not parse operand (%s): %v", operand, err)
			}
			nums = append(nums, n)
		}
		equations = append(equations, Eq{target, nums})
	}

	sum := 0
	for _, eq := range equations {
		t := eq.target
		nums := eq.nums

		stack := []string{"+", "*", "|"}
		for len(stack) > 0 {
			var top string
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]

			if len(top) == len(nums)-1 {
				acc := nums[0]
				for i, n := range nums[1:] {
					op := top[i]
					switch op {
					case '+':
						acc += n
					case '*':
						acc *= n
					case '|':
						combined := fmt.Sprintf("%d%d", acc, n)
						conv, err := strconv.Atoi(combined)
						if err != nil {
							log.Fatalf("could not convert in || operation: %v", err)
						}
						acc = conv
					}
				}

				if acc == t {
					sum += t
					break
				}
			} else {
				for _, op := range []string{"+", "*", "|"} {
					stack = append(stack, top+op)
				}
			}
		}

	}
	fmt.Println("PART 2 TOTAL:", sum)

}

func main() {
	// part1()
	part2()
}
