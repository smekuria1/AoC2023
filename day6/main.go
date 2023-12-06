package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	Part2()
}

func Part2() {
	linetimes := strings.Split(readlinesP2()[0], ":")
	linedistances := strings.Split(readlinesP2()[1], ":")
	fmt.Println(linetimes)
	fmt.Println(linedistances)
	var times []int
	var distances []int
	time, _ := strconv.Atoi(linetimes[1])
	times = append(times, time)
	distance, _ := strconv.Atoi(linedistances[1])
	distances = append(distances, distance)

	var validtimes []int
	for i, time := range times {
		record := distances[i]
		distance := 0
		button := 0
		count := 0
		for button = 1; button <= time; button++ {
			distance = button * (time - button)
			if distance > record {
				count += 1
			}

			if button+1 > time {
				validtimes = append(validtimes, count)
			}
		}

	}
	result := 1
	for _, vtime := range validtimes {
		result *= vtime
	}
	fmt.Printf("result: %v\n", result)
}

func Par1() {
	linetimes := strings.Split(readlines()[0], ":")
	linedistances := strings.Split(readlines()[1], ":")

	var times []int
	var distances []int
	for _, c := range strings.Split(linetimes[1], " ") {
		if c == "" {
			continue
		}
		c = SpaceMap(c)
		y, _ := strconv.Atoi(c)
		times = append(times, y)
	}

	for _, c := range strings.Split(linedistances[1], " ") {
		if c == "" {
			continue
		}
		c = SpaceMap(c)
		y, _ := strconv.Atoi(c)
		distances = append(distances, y)
	}
	var validtimes []int
	for i, time := range times {
		record := distances[i]
		distance := 0
		button := 0
		count := 0
		for button = 1; button <= time; button++ {
			distance = button * (time - button)
			if distance > record {
				count += 1
			}

			if button+1 > time {
				validtimes = append(validtimes, count)
			}
		}

	}
	result := 1
	for _, vtime := range validtimes {
		result *= vtime
	}
	fmt.Printf("result: %v\n", result)
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
func readlinesP2() []string {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, SpaceMap(scanner.Text()))
	}

	file.Close()

	return lines
}
