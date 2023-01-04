package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}

	rates := map[string]int{}
	tunnels := map[string][]string{}
	indexes := map[string]int{}

	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Split(line, ";")
		if len(parts) < 2 {
			log.Fatalf("could not parse input line `%s`: got %v", line, parts)
		}

		var name string
		var rate int
		_, err := fmt.Sscanf(parts[0], "Valve %s has flow rate=%d", &name, &rate)
		if err != nil {
			log.Fatalf("could not parse name and rate: %v", err)
		}

		offset := 23
		if strings.HasPrefix(parts[1], " tunnels") {
			offset = 24
		}
		tunnelsStr := parts[1][offset:]
		ts := strings.Split(tunnelsStr, ", ")

		rates[name] = rate
		tunnels[name] = ts
	}

	i := 0
	for v, rate := range rates {
		if v == "AA" || rate == 0 {
			continue
		}
		indexes[v] = i
		i++
	}

	distance := map[string]map[string]int{}
	type vd struct {
		v string
		d int
	}
	for valve := range tunnels {
		if valve != "AA" && rates[valve] == 0 {
			// no point to keep distances from 0-rate valves
			continue
		}

		q := []vd{{valve, 0}}
		visited := map[string]bool{valve: true}
		distance[valve] = map[string]int{}

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]

			for _, neighbor := range tunnels[cur.v] {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true
				if rates[neighbor] > 0 {
					distance[valve][neighbor] = cur.d + 1
				}
				q = append(q, vd{neighbor, cur.d + 1})
			}
		}
	}

	type state struct {
		at   string
		time int
		open int
	}
	cache := map[state]int{}

	var dfs func(state) int
	dfs = func(st state) int {
		if v, ok := cache[st]; ok {
			return v
		}

		max := 0

		for v, d := range distance[st.at] {
			//check if already open
			vn := 1 << indexes[v]
			if st.open&vn != 0 {
				continue
			}

			rem := st.time - d - 1
			if rem <= 0 {
				continue
			}
			sub := dfs(state{v, rem, st.open | vn}) + rates[v]*rem
			if sub > max {
				max = sub
			}
		}

		cache[st] = max
		return max
	}

	fmt.Printf("Part 1: %d\n", dfs(state{"AA", 30, 0}))

	var max int

	allOpen := (1 << len(indexes)) - 1
	for i := 0; i < (allOpen+1)/2; i++ {
		c := dfs(state{"AA", 26, i}) + dfs(state{"AA", 26, i ^ allOpen})

		if c > max {
			max = c
		}
	}
	fmt.Printf("Part 2: %d\n", max)
}
