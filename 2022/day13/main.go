package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func cmp(l, r any) int {
	vl, lok := l.([]any)
	vr, rok := r.([]any)
	if !lok && !rok {
		dl, dr := l.(float64), r.(float64)
		if dl < dr {
			return -1
		} else if dl > dr {
			return 1
		}
		return 0
	}

	if !lok {
		vl = []any{l}
	}

	if !rok {
		vr = []any{r}
	}

	for i := 0; i < len(vl); i++ {
		if i >= len(vr) {
			return 1
		}

		c := cmp(vl[i], vr[i])
		if c != 0 {
			return c
		}
	}
	if len(vl) < len(vr) {
		return -1
	}
	return 0
}

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}
	pairs := strings.Split(string(data), "\n\n")

	var sum int
	var packets []any

	for i, ps := range pairs {
		pair := strings.Split(ps, "\n")
		if len(pair) < 2 {
			log.Fatalf("could not parse pair: not enough packets: only have %d", len(pair))
		}

		fmt.Println(pair[0])
		fmt.Println(pair[1])
		var l, r any
		err := json.Unmarshal([]byte(pair[0]), &l)
		if err != nil {
			log.Fatalf("could not parse left packet %v: %v", pair[0], err)
		}
		err = json.Unmarshal([]byte(pair[1]), &r)
		if err != nil {
			log.Fatalf("could not parse right packet %v: %v", pair[1], err)
		}

		if cmp(l, r) < 0 {
			sum += i + 1
		}
		packets = append(packets, l, r)
	}

	fmt.Println(sum)

	packets = append(packets, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) < 0
	})

	decoderKey := 1
	for i, p := range packets {
		s := fmt.Sprint(p)
		if s == "[[2]]" || s == "[[6]]" {
			decoderKey *= i + 1
		}
	}
	fmt.Println(decoderKey)
}
