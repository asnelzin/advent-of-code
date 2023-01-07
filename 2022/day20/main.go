package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
	prev *Node
}

func printLinkedList(head *Node, length int) {
	n := head
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", n.val)
		n = n.next
	}
	fmt.Println()
}

// swap swaps two nodes in a linked list. First node moves to the right of the second node.
func swap(a *Node, b *Node) {
	// z <-> a <-> b <-> c
	// z <-> b <-> a <-> c

	a.next, b.next = b.next, a
	a.prev, b.prev = b, a.prev

	a.next.prev = a
	b.prev.next = b
}

func mix(order []*Node) {
	for _, current := range order {
		to := current
		if current.val > 0 {
			for i := current.val % (len(order) - 1); i > 0; i-- {
				to = current.next
				swap(current, to)
			}
		} else {
			for i := current.val % (len(order) - 1); i < 0; i++ {
				to = current.prev
				swap(to, current)
			}
		}
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}

	var zero *Node
	var order []*Node

	var prev *Node
	for _, line := range strings.Split(string(data), "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("could not parse number `%s`: %v", line, err)
		}

		node := &Node{n, nil, prev}
		if zero == nil && n == 0 {
			zero = node
		}
		if prev != nil {
			prev.next = node
		}
		prev = node

		order = append(order, node)
	}
	order[0].prev = order[len(order)-1]
	order[len(order)-1].next = order[0]

	part1(order, zero)
	//part2(order, zero)

}

func part1(order []*Node, zero *Node) {
	mix(order)

	xindex, yindex, zindex := 1000, 2000, 3000
	sum := 0
	for _, idx := range []int{xindex, yindex, zindex} {
		st := zero
		for i := 0; i < idx; i++ {
			st = st.next
		}
		sum += st.val
	}
	fmt.Println(sum)
}

func part2(order []*Node, zero *Node) {
	for _, n := range order {
		n.val *= 811589153
	}

	for i := 0; i < 10; i++ {
		mix(order)
	}

	xindex, yindex, zindex := 1000, 2000, 3000
	sum := 0
	for _, idx := range []int{xindex, yindex, zindex} {
		st := zero
		for i := 0; i < idx; i++ {
			st = st.next
		}
		sum += st.val
	}
	fmt.Println(sum)
}
