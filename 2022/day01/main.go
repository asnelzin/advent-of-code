package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	h := &IntHeap{}
	heap.Init(h)

	current := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Println(current)
			heap.Push(h, current)
			current = 0
			continue
		}
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("could not convert string to int: %v", err)
		}
		current += n
	}
	heap.Push(h, current)

	fmt.Println()

	max := 0
	for i := 0; i < 3; i++ {
		n := heap.Pop(h).(int)
		fmt.Println(n)
		max += n
	}
	fmt.Println()
	fmt.Println(max)
}
