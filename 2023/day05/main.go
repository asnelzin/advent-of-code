package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Range struct {
	min int
	max int
}

func (r Range) InRange(n int) bool {
	return r.min <= n && n <= r.max
}

func (r Range) String() string {
	return fmt.Sprintf("[%d-%d]", r.min, r.max)
}

type MapRecord struct {
	dstStart int
	src      Range
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

	parts := strings.Split(string(buf), "\n\n")
	var seeds []int
	seedstring := strings.TrimPrefix(parts[0], "seeds: ")
	for _, s := range strings.Split(seedstring, " ") {
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("could not parse seed: %v", s)
		}
		seeds = append(seeds, n)
	}

	current := seeds[:]
	for _, m := range parts[1:] {
		parts := strings.Split(m, "\n")
		var records []MapRecord
		var (
			minSrc int = math.MaxInt64
			maxSrc int = -1
		)
		for _, p := range parts[1:] {
			nums := strings.Split(p, " ")
			if len(nums) < 3 {
				log.Fatalf("could not parse map record: %v", p)
			}

			dstStart, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatalf("could not parse map record: %v", p)
			}
			srcmin, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatalf("could not parse map record: %v", p)
			}
			depth, err := strconv.Atoi(nums[2])
			if err != nil {
				log.Fatalf("could not parse map record: %v", p)
			}

			srcmax := srcmin + depth - 1
			records = append(records, MapRecord{
				dstStart: dstStart,
				src:      Range{srcmin, srcmin + depth - 1},
			})

			if srcmin < minSrc {
				minSrc = srcmin
			}
			if srcmax > maxSrc {
				maxSrc = srcmax
			}
		}

		for i, src := range current {
			if src < minSrc || src > maxSrc {
				continue
			}

			for _, r := range records {
				if r.src.InRange(src) {
					current[i] = r.dstStart + (src - r.src.min)
					break
				}
			}
		}
	}

	minLoc := current[0]
	for _, n := range current {
		if n < minLoc {
			minLoc = n
		}
	}
	fmt.Println(minLoc)
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

	parts := strings.Split(string(buf), "\n\n")
	var seedsrange []int
	seedstring := strings.TrimPrefix(parts[0], "seeds: ")
	for _, s := range strings.Split(seedstring, " ") {
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("could not parse seed: %v", s)
		}
		seedsrange = append(seedsrange, n)
	}

	var seeds []int
	for i := 1; i < len(seedsrange); i += 2 {
		seedmin := seedsrange[i-1]
		for j := seedmin; j <= seedmin+seedsrange[i]-1; j++ {
			seeds = append(seeds, j)
		}
	}

	current := seeds[:]
	for i, m := range parts[1:] {
		start := time.Now()
		fmt.Printf("staring map %d\n", i)
		parts := strings.Split(m, "\n")
		var records []MapRecord
		var (
			minSrc int = math.MaxInt64
			maxSrc int = -1
		)
		for _, p := range parts[1:] {
			nums := strings.Split(p, " ")
			if len(nums) < 3 {
				log.Fatalf("could not parse map record: %v", p)
			}

			dstStart, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatalf("could not parse map record: %v", p)
			}
			srcmin, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatalf("could not parse map record: %v", p)
			}
			depth, err := strconv.Atoi(nums[2])
			if err != nil {
				log.Fatalf("could not parse map record: %v", p)
			}

			srcmax := srcmin + depth - 1
			records = append(records, MapRecord{
				dstStart: dstStart,
				src:      Range{srcmin, srcmin + depth - 1},
			})

			if srcmin < minSrc {
				minSrc = srcmin
			}
			if srcmax > maxSrc {
				maxSrc = srcmax
			}
		}

		for i, src := range current {
			if src < minSrc || src > maxSrc {
				continue
			}

			for _, r := range records {
				if r.src.InRange(src) {
					current[i] = r.dstStart + (src - r.src.min)
					break
				}
			}
		}
		fmt.Printf("map %d took %v\n", i, time.Since(start))
	}

	minLoc := current[0]
	for _, n := range current {
		if n < minLoc {
			minLoc = n
		}
	}
	fmt.Println(minLoc)
}

func main() {
	part2()
}
