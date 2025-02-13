package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

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

	totalLength := 0
	diskMap := make([]int, 0, len(buf))
	for _, b := range buf {
		n := int(b - byte('0'))
		totalLength += n

		diskMap = append(diskMap, n)
	}

	decompressed := make([]int, totalLength)
	offset := 0
	for i, size := range diskMap {
		index := i / 2
		isFreeSpace := i%2 == 1

		filler := index
		if isFreeSpace {
			filler = -1
		}

		for j := 0; j < size; j++ {
			decompressed[offset+j] = filler
		}
		offset += size
	}

	left := 0
	right := totalLength - 1
	for {
		for decompressed[left] != -1 && left < len(decompressed)-1 {
			left++
		}
		for decompressed[right] == -1 && right > 0 {
			right--
		}
		if left >= right {
			break
		}
		decompressed[left], decompressed[right] = decompressed[right], decompressed[left]
	}

	checksum := 0
	for i, n := range decompressed {
		if n == -1 {
			continue
		}
		checksum += i * n
	}
	fmt.Println(checksum)
}

func printD(d []int) {
	for _, n := range d {
		if n == -1 {
			fmt.Print(".")
			continue
		}
		fmt.Print(n)
	}
	fmt.Println()
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

	totalLength := 0
	diskMap := make([]int, 0, len(buf))
	for _, b := range buf {
		n := int(b - byte('0'))
		totalLength += n

		diskMap = append(diskMap, n)
	}

	decompressed := make([]int, totalLength)
	// free := make(map[int][]int) // map of free space available with size->[]offset
	var free []F
	freeLess := func(i, j int) bool {
		return free[i].offset < free[j].offset
	}
	files := make(map[int]F) // id -> size, offset

	// for fileID from max to smallest
	// to move file, find leftmost free space of size X (filesize) (should know each files' size)
	// move whole file in here in decompressed view
	// if filesize perfectly fits, remove offset from free map
	// if filesize is less, remove offset from map, newoffset += filezie, newsize = oldsize - filesize

	offset := 0
	for i, size := range diskMap {
		index := i / 2
		isFreeSpace := i%2 == 1

		filler := index
		if isFreeSpace {
			filler = -1
			if size != 0 {
				free = append(free, F{size, offset})
				// if free[size] == nil {
				// 	free[size] = make([]int, 0)
				// }
				// free[size] = append(free[size], offset)
			}
		} else {
			files[index] = F{size, offset}
		}

		for j := 0; j < size; j++ {
			decompressed[offset+j] = filler
		}
		offset += size
	}

	sort.Slice(free, freeLess)

	for i := len(files) - 1; i >= 0; i-- {
		move(files[i], decompressed, free)
		// printD(decompressed)
	}

	checksum := 0
	for i, n := range decompressed {
		if n == -1 {
			continue
		}
		checksum += i * n
	}
	fmt.Println(checksum)
}

type F struct {
	size   int
	offset int
}

func move(meta F, decompressed []int, free []F) {
	// find leftmost available space to fit the file
	spaceOffset := -1
	spaceSize := -1
	// for i := meta.size; i <= len(decompressed); i++ {
	// 	_, ok := free[i]
	// 	if !ok || len(free[i]) == 0 {
	// 		continue
	// 	}
	// 	if free[i][0] > meta.offset {
	// 		continue
	// 	}
	// 	spaceOffset = free[i][0]
	// 	spaceSize = i
	// 	free[i] = free[i][1:]
	// 	break
	// }

	var toDelete int
	for i, gap := range free {
		if gap.size < meta.size {
			continue
		}
		if gap.offset > meta.offset {
			return
		}

		toDelete = i
		spaceOffset = gap.offset
		spaceSize = gap.size
		break
	}

	if spaceOffset == -1 {
		return
	}

	free = append(free[:toDelete], free[toDelete+1:]...)

	// if spaceOffset == -1 {
	// 	return
	// }

	// move file in decompressed form
	for i := 0; i < meta.size; i++ {
		l, r := meta.offset+i, spaceOffset+i
		decompressed[l], decompressed[r] = decompressed[r], decompressed[l]
	}

	// put back free space
	if spaceSize-meta.size <= 0 {
		return
	}

	// spaceOffsetAfterMove := spaceOffset + meta.size
	// if free[spaceSizeAfterMove] == nil {
	// 	free[spaceSizeAfterMove] = make([]int, 0)
	// }

	free = append(free, F{spaceSize - meta.size, spaceOffset + meta.size})
	sort.Slice(free, func(i, j int) bool {
		return free[i].offset < free[j].offset
	})

	// free[spaceSizeAfterMove] = append(free[spaceSizeAfterMove], spaceOffsetAfterMove)
	// slices.Sort(free[spaceSizeAfterMove])
	// newPlace := 0
	// for i := 0; i < len(free[spaceSizeAfterMove]); i++ {
	// 	if free[spaceSizeAfterMove][i] > spaceOffsetAfterMove {
	// 		break
	// 	}
	// 	newPlace++
	// }
	// if newPlace == 0 || newPlace == len(free[spaceSizeAfterMove]) {
	// 	free[spaceSizeAfterMove] = append(free[spaceSizeAfterMove], spaceOffsetAfterMove)
	// 	return
	// }
	//
	// insertAt(free[spaceSizeAfterMove], newPlace, spaceOffsetAfterMove)
}

func insertAt(s []int, index int, value int) []int {
	if index < 0 || index > len(s) {
		panic("index out of range")
	}
	s = append(s[:index], append([]int{value}, s[index:]...)...)
	return s
}

func main() {
	// part1()
	part2()
}
