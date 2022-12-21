package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"sort"
	"strings"
)

type Interval struct {
	start, end int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func distance(a image.Point, b image.Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func interval(sensor image.Point, beacon image.Point, y int) Interval {
	d := distance(sensor, beacon)
	k := abs(y - sensor.Y)

	x1, x2 := sensor.X-(d-k), sensor.X+(d-k)
	return Interval{min(x1, x2), max(x1, x2)}
}

func inside(sensor image.Point, beacon image.Point, y int) bool {
	d := distance(sensor, beacon)
	dline := abs(sensor.Y - y)
	return dline < d
}

func merge(intervals []Interval) []Interval {
	var merged []Interval

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].start < intervals[j].start {
			return true
		} else if intervals[i].start == intervals[j].start {
			return intervals[i].end < intervals[j].end
		}
		return false
	})
	acc := intervals[0]
	for _, i := range intervals[1:] {
		if (i.start-acc.end == 1) || (i.start <= acc.end && acc.end <= i.end) {
			acc.end = i.end
		} else if acc.end < i.start {
			merged = append(merged, acc)
			acc = i
		}
	}
	merged = append(merged, acc)

	return merged
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}

	keyRow := 2_000_000
	//keyRow := 10

	onRow := map[image.Point]bool{}
	var intervals []Interval

	for _, line := range strings.Split(string(data), "\n") {
		var sensor, beacon image.Point
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)
		if err != nil {
			log.Fatalf("could not parse input line: %v", err)
		}

		if !inside(sensor, beacon, keyRow) {
			continue
		}

		i := interval(sensor, beacon, keyRow)
		intervals = append(intervals, i)

		if beacon.Y == keyRow {
			onRow[beacon] = true
		}
	}

	cant := 0

	for _, i := range merge(intervals) {
		cant += i.end - i.start + 1
	}
	fmt.Println(cant - len(onRow))
}
