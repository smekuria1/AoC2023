package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	Part1()
}

type Hand struct {
	hand string
	bid  int
	rank int
}

func Part1() {
	lines := readlines()
	var hands []Hand
	category := make(map[string][]Hand, 5)
	for _, c := range lines {
		h := strings.Split(c, " ")[0]
		b := strings.Split(c, " ")[1]
		bidInt, _ := strconv.Atoi(b)
		hand := Hand{
			hand: h,
			bid:  bidInt,
		}
		hands = append(hands, hand)
	}
	for _, hand := range hands {
		if isFiveOf(hand) {
			addtoMap(category, "fiveof", hand)
		} else if isFourof(hand) {
			addtoMap(category, "fourof", hand)
		} else if isFullhouse(hand) {
			addtoMap(category, "fullhouse", hand)
		} else if isThreeof(hand) {
			addtoMap(category, "threeof", hand)
		} else if isTwoPairs(hand) {
			addtoMap(category, "twopair", hand)
		} else if isOnePair(hand) {
			addtoMap(category, "onepair", hand)
		} else {
			addtoMap(category, "high", hand)
		}
	}

	sortedcards := make([]Hand, 0, len(hands))

	handTypes := []string{"high", "onepair", "twopair", "threeof", "fullhouse", "fourof", "fiveof"}
	result := 0
	for _, types := range handTypes {
		for _, v := range category[types] {
			v.rank = len(sortedcards) + 1
			result += v.rank * v.bid
			sortedcards = append(sortedcards, v)
		}
	}

	fmt.Printf("len(sortedcards): %v\n", len(sortedcards))
	fmt.Printf("sortedcards: %v\n", sortedcards)
	fmt.Printf("result: %v\n", result)

}

func digitCmp(a, b Hand) int {
	for i := 0; i < len(a.hand); i++ {
		strenA := getCardStrength(a.hand[i])
		strenB := getCardStrength(b.hand[i])

		switch {
		case strenA > strenB:
			return 1
		case strenA < strenB:
			return -1
		}

	}

	return 0
}
func addtoMap(category map[string][]Hand, catName string, hand Hand) {
	category[catName] = append(category[catName], hand)
	hands := category[catName]

	slices.SortFunc(hands, digitCmp)
	category[catName] = hands

}

func isFiveOf(a Hand) bool {
	firstchar := string(a.hand[0])
	for _, c := range a.hand {
		if string(c) != firstchar {
			return false
		}

	}
	return true
}

func isFourof(a Hand) bool {
	count := make(map[byte]int)
	for _, c := range a.hand {
		count[byte(c)]++
		if count[byte(c)] == 4 {
			return true
		}
	}

	return false
}

func isFullhouse(a Hand) bool {
	str := a.hand
	s := []rune(str)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return (s[0] == s[1] && s[2] == s[3] && s[3] == s[4]) ||
		(s[0] == s[1] && s[1] == s[2] && s[3] == s[4])
}

func isThreeof(hand Hand) bool {
	count := make(map[byte]int)

	for _, c := range hand.hand {
		count[byte(c)]++
		if count[byte(c)] == 3 {
			return true
		}
	}

	return false
}

func isTwoPairs(hand Hand) bool {
	count := make(map[byte]int)
	pairCount := 0

	for _, c := range hand.hand {
		if count[byte(c)] == 1 {
			pairCount++
		}
		count[byte(c)]++
	}

	return pairCount == 2
}

func isOnePair(hand Hand) bool {
	count := make(map[byte]int)

	for _, c := range hand.hand {
		count[byte(c)]++
		if count[byte(c)] == 2 {
			return true
		}
	}

	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readlines() []string {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return lines
}

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func getCardStrength(digit byte) int {
	switch digit {
	case 'A':
		return 13
	case 'K':
		return 12
	case 'Q':
		return 11
	case 'J':
		return 10
	case 'T':
		return 9
	case '9':
		return 8
	case '8':
		return 7
	case '7':
		return 6
	case '6':
		return 5
	case '5':
		return 4
	case '4':
		return 3
	case '3':
		return 2
	case '2':
		return 1
	default:
		return 0 // Default to 0 for invalid characters
	}
}
