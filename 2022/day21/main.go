package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation func(int, int) int

type Monkey struct {
	name string
	val  int

	left  string
	right string

	opval string
	op    Operation

	parent *Monkey
}

var (
	add Operation = func(l int, r int) int { return l + r }
	sub Operation = func(l int, r int) int { return l - r }
	mul Operation = func(l int, r int) int { return l * r }
	div Operation = func(l int, r int) int { return l / r }
)

func proveOnlyOnePathToHumn(monkeys map[string]*Monkey) {
	var currentPath []string
	var allPaths [][]string
	visited := map[string]bool{}
	var paths func(start string, end string)
	paths = func(u string, v string) {
		if visited[u] == true {
			return
		}
		visited[u] = true
		currentPath = append(currentPath, u)
		if u == v {
			var p []string
			for _, n := range currentPath {
				p = append(p, n)
			}
			allPaths = append(allPaths, p)
			visited[u] = false
			currentPath = currentPath[:len(currentPath)-1]
			return
		}

		m := monkeys[u]
		if m.left != "" {
			paths(monkeys[u].left, v)
		}
		if m.right != "" {
			paths(monkeys[u].right, v)
		}

		currentPath = currentPath[:len(currentPath)-1]
		visited[u] = false
	}

	paths("root", "humn")

	if len(allPaths) > 1 {
		panic("There is more than one path from `root` to `humn`: need a new solution")
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}

	monkeys := map[string]*Monkey{}
	for _, line := range strings.Split(string(data), "\n") {
		m := Monkey{}
		parts := strings.Split(line, " ")

		if len(parts) == 2 { // Ex: `root: 20`
			m.name = parts[0][:len(parts[0])-1]
			m.val, err = strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalf("could not parse number `%s`: %v", parts[1], err)
			}
		} else if len(parts) == 4 { // Ex: `root: alice + bob`
			m.name = parts[0][:len(parts[0])-1]
			m.left = parts[1]
			op := parts[2]
			m.right = parts[3]
			m.opval = op

			switch op {
			case "+":
				m.op = add
			case "-":
				m.op = sub
			case "*":
				m.op = mul
			case "/":
				m.op = div
			}
		} else {
			log.Fatalf("not enough parts in line: `%s`", line)
		}
		monkeys[m.name] = &m
	}

	// fill parent
	for _, m := range monkeys {
		if m.left != "" {
			monkeys[m.left].parent = m

		}
		if m.right != "" {
			monkeys[m.right].parent = m
		}
	}
	monkeys["root"].parent = nil

	var dfs func(*Monkey) int
	dfs = func(m *Monkey) int {
		if m.val != 0 {
			return m.val
		}

		return m.op(dfs(monkeys[m.left]), dfs(monkeys[m.right]))
	}
	fmt.Printf("Part 1 = %d\n", dfs(monkeys["root"]))

	proveOnlyOnePathToHumn(monkeys)

	var path []string
	cur := monkeys["humn"]
	for cur != nil {
		path = append(path, cur.name)
		cur = cur.parent
	}

	var humn int
	for i := len(path) - 2; i >= 0; i-- {
		slow := monkeys[path[i+1]]
		fast := monkeys[path[i]]

		isLeft := slow.left == fast.name

		other := 0
		if isLeft {
			other = dfs(monkeys[slow.right])
		} else {
			other = dfs(monkeys[slow.left])
		}

		// set initial
		if slow.name == "root" {
			humn = other
			continue
		}

		switch slow.opval {
		case "+":
			humn = sub(humn, other)
		case "*":
			humn = div(humn, other)
		case "-":
			if isLeft {
				humn = add(humn, other)
			} else {
				humn = sub(other, humn)
			}
		case "/":
			if isLeft {
				humn = mul(humn, other)
			} else {
				humn = div(other, humn)
			}
		}
	}
	fmt.Printf("Part 2 = %d\n", humn)

	// check answer
	monkeys["root"].op = sub
	monkeys["humn"].val = humn
	if dfs(monkeys["root"]) != 0 {
		panic("Answer is not correct")
	}
}
