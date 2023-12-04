package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	Part1()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsNumeric(s string) bool {
	numeric := regexp.MustCompile(`\d`).MatchString(s)
	return numeric
}

func Part1() {
	matrix := BuildMatrix()
	printMatrix(matrix)
}

func BuildMatrix() [][]string {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	var matrix [][]string
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "")
		matrix = append(matrix, items)

	}

	return matrix
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Print(element + " ")
		}
		fmt.Println()
	}
}
