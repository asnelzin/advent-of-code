package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type CardType int

const (
	highCard CardType = iota
	onePair
	twoPairs
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func order(c byte) int {
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
	switch c {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'J':
		return -1 // 9
	case 'T':
		return 8
	default:
		return int(c-'0') - 2
	}
}

type Hand struct {
	cards []byte
	bid   int
}

func (h *Hand) containsJoker() bool {
	return strings.Contains(string(h.cards), "J")
}

func (h *Hand) String() string {
	return string(h.cards)
}

func (h *Hand) TypePart1() CardType {
	count := make(map[byte]int)
	for _, c := range h.cards {
		count[c]++
	}

	maxRepeat := 0
	for _, v := range count {
		if v > maxRepeat {
			maxRepeat = v
		}
	}

	switch maxRepeat {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		if len(count) == 2 {
			return fullHouse
		}
		return threeOfAKind
	case 2:
		if len(count) == 3 {
			return twoPairs
		}
		return onePair
	}

	return highCard
}

func (h *Hand) Type() CardType {
	if !h.containsJoker() {
		return h.TypePart1()
	}

	// replace joker with all possible cards
	maxType := highCard
	for _, c := range []byte("AKQT98765432") {
		newHand := Hand{}
		newHand.cards = []byte(strings.ReplaceAll(string(h.cards), "J", string(c)))
		maxType = max(maxType, newHand.TypePart1())
	}
	return maxType
}

func (h *Hand) Less(other Hand) bool {
	if h.Type() != other.Type() {
		return h.Type() < other.Type()
	}

	for i := 0; i < len(h.cards); i++ {
		if order(h.cards[i]) != order(other.cards[i]) {
			return order(h.cards[i]) < order(other.cards[i])
		}
	}
	return false
}

// Hands attaches the methods of Interface to []Hand, sorting in increasing order.
type Hands []Hand

func (x Hands) Len() int           { return len(x) }
func (x Hands) Less(i, j int) bool { return x[i].Less(x[j]) }
func (x Hands) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

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

	lines := strings.Split(string(buf), "\n")
	hands := make(Hands, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatalf("invalid line: %s", line)
		}

		hand := Hand{}
		hand.cards = []byte(parts[0])
		_, err := fmt.Sscanf(parts[1], "%d", &hand.bid)
		if err != nil {
			log.Fatalf("invalid line: %s", line)
		}
		hands = append(hands, hand)
	}

	sort.Sort(hands)

	total := 0
	for rank, hand := range hands {
		total += hand.bid * (rank + 1)
	}
	fmt.Println(total)
}
