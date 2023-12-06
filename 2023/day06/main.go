package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	parts := strings.Split(string(buf), "\n")
	if len(parts) < 2 {
		log.Fatalf("invalid input")
	}

	parts[0] = strings.TrimPrefix(parts[0], "Time: ")
	parts[1] = strings.TrimPrefix(parts[1], "Distance: ")

	var times []int
	for _, v := range strings.Split(parts[0], " ") {
		if v == "" {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("could not parse time: %v", v)
		}
		times = append(times, n)
	}
	var distances []int
	for _, v := range strings.Split(parts[1], " ") {
		if v == "" {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("could not parse distance: %v", v)
		}
		distances = append(distances, n)
	}

	mul := 1
	for i, t := range times {
		roots := func(T float64, d float64) (float64, float64) {
			return 0.5 * (T - math.Sqrt(T*T-4*d)), 0.5 * (T + math.Sqrt(T*T-4*d))
		}

		rl, rh := roots(float64(t), float64(distances[i]))
		if rl == math.Trunc(rl) {
			rl += 1.0
		}
		if rh == math.Trunc(rh) {
			rh -= 1.0
		}

		lo, hi := int(math.Ceil(rl)), int(math.Floor(rh))
		mul *= hi - lo + 1
	}

	fmt.Println(mul)
}
