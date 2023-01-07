package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var cargoRe = regexp.MustCompile(`\s?(\[[A-Z]\])\s?|\s?(   )\s?`)

type Stack struct {
	data []string
}

func (s *Stack) Pop() string {
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last
}

func (s *Stack) PopMany(amount int) []string {
	size := len(s.data)
	tail := s.data[size-amount : size]
	s.data = s.data[:size-amount]
	return tail
}

func (s *Stack) Push(r string) {
	s.data = append(s.data, r)
}

func (s *Stack) PushMany(items []string) {
	for _, i := range items {
		s.data = append(s.data, i)
	}
}

func parseData(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	var split int
	for idx, s := range data {
		if s == "" {
			split = idx
			break
		}
	}
	table := data[:split]
	commands := data[split+1:]
	return table, commands
}

func prepareStacks(table []string) []Stack {
	header := table[len(table)-1]
	stacksN := int(header[len(header)-1] - '0')
	stacks := make([]Stack, stacksN)

	for i := len(table) - 2; i >= 0; i-- {
		cargo := cargoRe.FindAllStringSubmatch(table[i], -1)
		for n, cc := range cargo {
			c := cc[1]
			if c == "" {
				continue
			}
			stacks[n].Push(c[1 : len(c)-1])
		}
	}
	return stacks
}

func part1() {
	table, commands := parseData("input.txt")
	stacks := prepareStacks(table)

	for _, command := range commands {
		var (
			amount int
			from   int
			to     int
		)
		fmt.Sscanf(command, "move %d from %d to %d", &amount, &from, &to)
		from -= 1
		to -= 1

		for i := 0; i < amount; i++ {
			stacks[to].Push(stacks[from].Pop())
		}
	}

	top := strings.Builder{}
	for _, s := range stacks {
		top.WriteString(s.Pop())
	}
	fmt.Println(top.String())
}

func part2() {
	table, commands := parseData("input.txt")
	stacks := prepareStacks(table)

	for _, command := range commands {
		var (
			amount int
			from   int
			to     int
		)
		fmt.Sscanf(command, "move %d from %d to %d", &amount, &from, &to)
		from -= 1
		to -= 1
		stacks[to].PushMany(stacks[from].PopMany(amount))
	}

	top := strings.Builder{}
	for _, s := range stacks {
		top.WriteString(s.Pop())
	}
	fmt.Println(top.String())
}

func main() {
	part1()
	part2()
}
