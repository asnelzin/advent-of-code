package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	name      string
	parent    *Dir
	size      int
	totalsize int
	sub       []*Dir
}

func parseTree(filename string) *Dir {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	var pwd *Dir = nil
	var root *Dir = nil

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			c := strings.Split(line, " ")
			cmd := c[1]
			if cmd == "ls" {
				continue
			}
			if cmd != "cd" {
				// impossible with current input
				continue
			}
			arg := c[2]

			if arg == ".." {
				if pwd == nil {
					continue
				}

				if pwd.parent == nil {
					pwd.parent = &Dir{
						name:   "?",
						parent: nil,
					}
					root = pwd.parent
				}

				if pwd == root {
					root = pwd.parent
				}

				pwd.parent.totalsize += pwd.totalsize
				pwd = pwd.parent
			} else {
				newdir := &Dir{
					name:   arg,
					parent: pwd,
				}
				if root == nil {
					root = newdir
				}
				if pwd != nil {
					pwd.sub = append(pwd.sub, newdir)
				}
				pwd = newdir
			}

		} else {
			// ls output
			if pwd == nil {
				continue
			}
			item := strings.Split(line, " ")
			t := item[0]
			if t == "dir" {
				continue
			}
			filesize, err := strconv.Atoi(t)
			if err != nil {
				log.Fatalf("could not convert filesize to int: %v", err)
			}

			pwd.size += filesize
			pwd.totalsize += filesize
		}
	}

	for pwd != root {
		pwd.parent.totalsize += pwd.totalsize
		pwd = pwd.parent
	}

	return root
}

func printDir(d *Dir, level int) {
	for k := 0; k < level*2; k++ {
		fmt.Print(" ")
	}
	fmt.Printf("- %s (size %d | total %d)\n", d.name, d.size, d.totalsize)
}

func printTree(root *Dir, level int) {
	if level == 0 {
		printDir(root, level)
	}
	for _, subdir := range root.sub {
		printDir(subdir, level+1)
		printTree(subdir, level+1)
	}
}

func findTotalUpTo(root *Dir, upto int, sum *int) {
	for _, sub := range root.sub {
		if sub.totalsize <= upto {
			*sum += sub.totalsize
		}
		findTotalUpTo(sub, upto, sum)
	}
}

func findSmallestUpTo(root *Dir, minsize int, result *int) {
	for _, sub := range root.sub {
		size := sub.totalsize
		if size >= minsize && (*result == 0 || *result > size) {
			*result = size
		}
		findSmallestUpTo(sub, minsize, result)
	}
}

func main() {
	root := parseTree("input.txt")
	printTree(root, 0)

	var sum int
	findTotalUpTo(root, 100000, &sum)
	fmt.Println(sum)

	minsize := 30000000 - (70000000 - root.totalsize)
	var dsize int
	findSmallestUpTo(root, minsize, &dsize)
	fmt.Println(dsize)
}
