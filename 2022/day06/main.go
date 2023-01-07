package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	stream := string(buf)
	ml := 14
	for i := range stream {
		if i-ml+1 < 0 {
			continue
		}
		visited := make(map[byte]bool, ml)
		for j := i - ml + 1; j <= i; j++ {
			visited[stream[j]] = true
		}
		if len(visited) == ml {
			fmt.Println(i + 1)
			return
		}
	}
}
