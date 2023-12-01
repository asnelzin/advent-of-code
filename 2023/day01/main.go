package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func replaceDigitsStrings(s string) string {
	modified := s
	for ds, d := range map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	} {
		modified = strings.ReplaceAll(modified, ds, d)
	}
	return modified
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = replaceDigitsStrings(line) // part2
		for i := 0; i < len(line); i++ {
			ord := int(line[i] - '0')
			if ord < 0 || ord > 9 {
				continue
			}
			sum += ord * 10
			break
		}

		for i := len(line) - 1; i >= 0; i-- {
			ord := int(line[i] - '0')
			if ord < 0 || ord > 9 {
				continue
			}
			sum += ord
			break
		}
	}
	fmt.Println(sum)
}
