package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func add(x int) func(int) int {
	return func(level int) int {
		return level + x
	}
}

func multiply(x int) func(int) int {
	return func(level int) int {
		return level * x
	}
}

func pow2(level int) int {
	return level * level
}

func divisible(x int) func(int) bool {
	return func(level int) bool {
		return level%x == 0
	}
}

type Monkey struct {
	ID        int
	items     []int
	operation func(level int) int
	test      func(level int) bool
	next      map[bool]int
}

func NewMonkey(id int) *Monkey {
	return &Monkey{ID: id, next: make(map[bool]int, 2)}
}

func round(all []*Monkey, activity []int, reducer func(int) int) {
	for _, m := range all {
		var old int
		for len(m.items) > 0 {
			old, m.items = m.items[0], m.items[1:]
			level := m.operation(old)

			level = reducer(level)

			nextID := m.next[m.test(level)]

			next := all[nextID]
			next.items = append(next.items, level)
			activity[m.ID]++
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		all     []*Monkey
		current *Monkey
		modulo  = 1
	)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			all = append(all, current)
			continue
		}

		if strings.HasPrefix(line, "Monkey") {
			id, err := strconv.Atoi(line[7:8])
			if err != nil {
				log.Printf("could not parse monkey's ID %s: %v", line[7:8], err)
				continue
			}
			current = NewMonkey(id)
		}

		if strings.HasPrefix(line, "  Starting") {
			for _, it := range strings.Split(line[18:], ", ") {
				itemID, err := strconv.Atoi(it)
				if err != nil {
					log.Printf("could not parse monkey's item %s: %v", it, err)
					continue
				}
				current.items = append(current.items, itemID)
			}
		}

		if strings.HasPrefix(line, "  Operation") {
			fn := strings.Split(line[19:], " ")
			if len(fn) < 3 {
				log.Printf("could not parse opeation: not enough data: %v", fn)
				continue
			}

			op := fn[1]
			right := fn[2]

			if right == "old" {
				current.operation = pow2
				continue
			}
			x, err := strconv.Atoi(right)
			if err != nil {
				log.Printf("could not parse operation right operand %s: %v", right, err)
				continue
			}
			switch op {
			case "*":
				current.operation = multiply(x)
			case "+":
				current.operation = add(x)
			}
		}

		if strings.HasPrefix(line, "  Test") {
			n := line[21:]
			x, err := strconv.Atoi(n)
			if err != nil {
				log.Printf("could not parse divisible test number %s: %v", n, err)
				continue
			}
			current.test = divisible(x)
			modulo *= x
		}

		if strings.HasPrefix(line, "    If true") {
			n := line[29:]
			id, err := strconv.Atoi(n)
			if err != nil {
				log.Printf("could not parse next monkey ID `%s`: %v", n, err)
				continue
			}
			current.next[true] = id
		}

		if strings.HasPrefix(line, "    If false") {
			n := line[30:]
			id, err := strconv.Atoi(n)
			if err != nil {
				log.Printf("could not parse next monkey ID `%s`: %v", n, err)
				continue
			}
			current.next[false] = id
		}
	}
	// add last one
	all = append(all, current)

	// part 1
	var activity = make([]int, len(all))
	for i := 0; i < 20; i++ {
		round(all, activity, func(l int) int {
			return int(math.Floor(float64(l) / 3.0))
		})
	}
	sort.Slice(activity, func(i, j int) bool {
		return activity[i] > activity[j]
	})
	fmt.Println(activity[0] * activity[1])

	// part 2
	activity = make([]int, len(all))
	for i := 0; i < 10_000; i++ {
		round(all, activity, func(l int) int {
			return l % modulo
		})
	}
	sort.Slice(activity, func(i, j int) bool {
		return activity[i] > activity[j]
	})
	fmt.Println(activity[0] * activity[1])
}
