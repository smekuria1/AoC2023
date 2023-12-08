package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {

	Part2()
}

type Node struct {
	left, right string
}

func Part1() {
	lines := readlines()
	moves := lines[0]
	nodes := lines[1:]

	fmt.Printf("moves: %v\n", moves)

	graph := buildgraph(nodes)
	fmt.Printf("graph: %v\n", graph)

	curr := "AAA"
	steps := 0

	for curr != "ZZZ" {
		move := moves[steps%len(moves)]
		if move == 'L' {
			curr = graph[curr].left
		} else {
			curr = graph[curr].right
		}

		steps += 1
	}
	fmt.Printf("steps: %v\n", steps)
}

func Part2() {
	lines := readlines()
	moves := lines[0]
	nodes := lines[1:]

	fmt.Printf("moves: %v\n", moves)

	graph := buildgraph(nodes)
	fmt.Printf("graph: %v\n", graph)

	stepsC := []int{}
	for key := range graph {
		if key[2] == 'A' {
			curr := key
			steps := 0
			for curr[2] != 'Z' {
				move := moves[steps%len(moves)]
				if move == 'L' {
					curr = graph[curr].left
				} else {
					curr = graph[curr].right
				}

				steps += 1

			}
			stepsC = append(stepsC, steps)
		}

	}

	fmt.Printf("stepsC: %v\n", stepsC)
	lcm := LCM(stepsC[0], stepsC[1], stepsC...)
	fmt.Printf("lcm: %v\n", lcm)
}

func buildgraph(nodes []string) map[string]Node {
	list := make(map[string]Node, len(nodes))
	for _, node := range nodes {
		splitNode := splitInput(node)

		list[splitNode[0]] = Node{splitNode[1], splitNode[2]}
	}

	return list

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitInput(input string) []string {
	// Define a regular expression to capture the values inside parentheses
	re := regexp.MustCompile(`^(\w+)=\((\w+),(\w+)\)$`)

	// Find submatches in the input string
	matches := re.FindStringSubmatch(input)

	// Extract the values from the submatches
	result := matches[1:]

	return result
}
func readlines() []string {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		lines = append(lines, SpaceMap(scanner.Text()))
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
