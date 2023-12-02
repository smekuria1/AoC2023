package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?
*/
func main() {
	Part2()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

var colormap = map[string]int{
	"blue":  14,
	"red":   12,
	"green": 13,
}

func Part1() {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		gameline := strings.Split(scanner.Text(), ":")
		subsets := strings.Split(gameline[1], ";")
		skipFlag := false
		for _, c := range subsets {
			if skipFlag {
				break
			}
			cubes := strings.Split(c, ",")
			cubebreak := false
			for _, cube := range cubes {

				numofcube, _ := strconv.Atoi(strings.Trim(string(cube[1:3]), " "))
				colorofCube := strings.TrimLeft((cube[3:]), " ")

				switch colorofCube {
				case "blue":
					if colormap["blue"] < numofcube {
						cubebreak = true
					}
				case "red":
					if colormap["red"] < numofcube {
						cubebreak = true
					}
				case "green":
					if colormap["green"] < numofcube {
						cubebreak = true
					}
				}
				if cubebreak {
					skipFlag = true
					break
				}

			}

		}

		if !skipFlag {
			gameNum, _ := strconv.Atoi(strings.Split(gameline[0], " ")[1])
			result += gameNum

		}

	}

	fmt.Println("The sum = ", result)
}

func Part2() {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	result := 0

	for scanner.Scan() {
		gameline := strings.Split(scanner.Text(), ":")
		subsets := strings.Split(gameline[1], ";")
		powermap := make(map[string]int)
		for _, c := range subsets {

			cubes := strings.Split(c, ",")
			for _, cube := range cubes {
				numofcube, _ := strconv.Atoi(strings.Trim(string(cube[1:3]), " "))
				colorofCube := strings.TrimLeft((cube[3:]), " ")
				val, ok := powermap[colorofCube]
				if ok {
					if val < numofcube {
						powermap[colorofCube] = numofcube
					}
				} else {
					powermap[colorofCube] = numofcube
				}

			}
		}
		products := 1
		for _, v := range powermap {
			products *= v
		}
		result += products
	}

	fmt.Println("The power sum = ", result)
}
