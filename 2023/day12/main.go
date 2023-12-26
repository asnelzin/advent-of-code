package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var cache map[string]int

func fits(record string, start int, end int) bool {
	// returns true if we can put # from start to end
	if start-1 < 0 || end+1 >= len(record) {
		return false
	}

	if record[start-1] == '#' || record[end+1] == '#' {
		return false
	}

	if strings.Contains(record[:start], "#") {
		return false
	}

	for i := start; i <= end; i++ {
		if record[i] == '.' {
			return false
		}
	}

	return true
}

func dfs(record string, groups []int) int {
	if v, ok := cache[key(record, groups)]; ok {
		return v
	}

	if len(groups) == 0 {
		if strings.Contains(record, "#") {
			return 0
		}
		return 1
	}
	size := groups[0]
	groups = groups[1:]

	count := 0
	for end := 0; end < len(record); end++ {
		start := end - (size - 1)
		if fits(record, start, end) {
			arrangements := dfs(record[end+1:], groups)
			cache[key(record[end+1:], groups)] = arrangements
			count += arrangements
		}
	}
	return count
}

func key(s string, g []int) string {
	return fmt.Sprintf("%s-%v", s, g)
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
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			log.Fatalf("could not parse line: %v", err)
		}

		record := parts[0]
		record = fmt.Sprintf(".%s.", record)
		var groups []int
		for _, size := range strings.Split(parts[1], ",") {
			var n int
			_, err := fmt.Sscanf(size, "%d", &n)
			if err != nil {
				log.Fatalf("could not parse line: %v", err)
			}
			groups = append(groups, n)
		}

		total += dfs(record, groups)
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
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			log.Fatalf("could not parse line: %v", err)
		}

		record := parts[0]

		// unfold the record
		record = strings.Repeat(fmt.Sprintf("%s?", record), 5)
		record = fmt.Sprintf(".%s.", record[:len(record)-1])

		// unfold groups
		g := parts[1]
		g = strings.Repeat(fmt.Sprintf("%s,", g), 5)
		g = g[:len(g)-1]

		var groups []int
		for _, size := range strings.Split(g, ",") {
			var n int
			_, err := fmt.Sscanf(size, "%d", &n)
			if err != nil {
				log.Fatalf("could not parse line: %v", err)
			}
			groups = append(groups, n)
		}

		total += dfs(record, groups)
	}
	fmt.Println(total)
}

func main() {
	cache = make(map[string]int)

	//part1()
	part2()
}
