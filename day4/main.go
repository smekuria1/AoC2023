package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Part2()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1() {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	winTracker := 0

	for scanner.Scan() {
		winmap := make(map[string]bool)

		card := strings.Split(scanner.Text(), ":")
		nums := strings.Split(card[1], "|")
		winNums := strings.Split(nums[0], " ")
		cardnums := strings.Split(nums[1], " ")

		for _, num := range winNums {
			if num == "" {
				continue
			}
			winmap[num] = true

		}
		fmt.Printf("winmap: %v\n", winmap)
		// fmt.Printf("winNums: %v\n", winNums)
		// fmt.Printf("cardnums: %v\n", cardnums)
		cardwintracker := 0
		for _, n := range cardnums {
			if val, ok := winmap[n]; ok && val {
				if cardwintracker == 0 {
					cardwintracker += 1
				} else {
					cardwintracker *= 2
				}
				fmt.Printf("found: %v\n", n)
				//fmt.Printf("cardwintracker: %v\n", cardwintracker)
			}
		}

		winTracker += cardwintracker
		fmt.Printf("winTracker: %v\n", winTracker)
	}

	fmt.Println("The sum", winTracker)
	file.Close()
}

func Part2() {
	lines, _ := readline()
	cardcounter := make(map[int]int)
	var p1 int
	for k, card := range lines {
		matches := cardwincalculator(card)
		p1 += matches
		cardname := k + 1
		cardcounter[cardname]++
		for i := 1; i <= matches; i++ {
			cardcounter[cardname+i] += cardcounter[cardname]
		}
	}

	var result int
	for _, v := range cardcounter {
		result += v
	}
	// fmt.Printf("cardcounter: %v\n", cardcounter)
	fmt.Printf("result: %v\n", result)
	// fmt.Printf("p1: %v\n", p1)

}

func cardwincalculator(card string) int {
	line := strings.Split(card, ":")
	nums := strings.Split(line[1], "|")

	winNums := strings.Split(nums[0], " ")
	cardnums := strings.Split(nums[1], " ")

	winmap := make(map[string]bool)
	for _, num := range winNums {
		if num == "" {
			continue
		}
		winmap[num] = true

	}
	cardwintracker := 0
	for _, n := range cardnums {
		if val, ok := winmap[n]; ok && val {
			cardwintracker++
			//fmt.Printf("found: %v\n", n)
			//fmt.Printf("cardwintracker: %v\n", cardwintracker)
		}
	}
	//fmt.Printf("winmap: %v\n", winmap)

	return cardwintracker
}

func readline() ([]string, error) {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}
