package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/* The newly-improved calibration document consists of lines of text; each line originally
contained a specific calibration value that the Elves now need to recover.
On each line, the calibration value can be found y combining the first digit and the last digit (
in that order) to form a single two-digit number.*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	Part2()
}

func Part1() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		out := ""
		for _, c := range line {
			if unicode.IsDigit(c) {
				out += string(c)
				break
			}
		}

		rev := reverseString(line)
		for _, c := range rev {
			if unicode.IsDigit(c) {
				out += string(c)
				break
			}
		}

		if len(out) < 2 {
			out += out
		}

		num, _ := strconv.Atoi(out)
		sum += num
		fmt.Println(out)
	}
	fmt.Println("The sum = ", sum)
	file.Close()
}

// function to reverse string
func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _, v := range str {
		result = string(v) + result
	}
	return
}

/*
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line.
*/

var numbermap = map[string]int{
	"one":   1,
	"eno":   1,
	"two":   2,
	"owt":   2,
	"three": 3,
	"eerht": 3,
	"four":  4,
	"ruof":  4,
	"five":  5,
	"evif":  5,
	"six":   6,
	"xis":   6,
	"seven": 7,
	"neves": 7,
	"eight": 8,
	"thgie": 8,
	"nine":  9,
	"enin":  9,
}

func Part2() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		out := ""
		line := scanner.Text()
		rev := reverseString(line)
		firstIndex, firstDigit := Part2FirstDig(line)
		secondIndex, secondDigit := Part2SecondDig(rev)
		for i, c := range line {
			if unicode.IsDigit(c) {
				if i < firstIndex {
					//fmt.Println("Found in unicode check", i, "firstDigit", string(line[i]))
					firstIndex = i
					firstDigit, _ = strconv.Atoi(string(line[i]))
				}
				break
			}
		}

		for i, c := range rev {
			if unicode.IsDigit(c) {
				if i < secondIndex {
					//fmt.Println("Found in unicode check", i, "secondDigit", string(rev[i]))
					secondIndex = i
					secondDigit, _ = strconv.Atoi(string(rev[i]))
				}
				break
			}
		}

		out = fmt.Sprintf("%v%v", firstDigit, secondDigit)
		outNum, _ := strconv.Atoi(out)
		sum += outNum
		fmt.Println("Final Digit", out)

	}

	fmt.Println("The sum = ", sum)
}

func Part2FirstDig(line string) (int, int) {
	keys := make([]string, len(numbermap))
	i := 0
	smallIndex := 1000
	for k := range numbermap {
		keys[i] = k
		i++
	}
	var dig string

	for _, k := range keys {
		if firstI := strings.Index(line, k); firstI >= 0 && firstI < smallIndex {
			smallIndex = firstI
			dig = k
			fmt.Println(line)
			//fmt.Println("Found in map check", smallIndex, "firstDigit", k)
		}
	}

	return smallIndex, numbermap[dig]

}

func Part2SecondDig(line string) (int, int) {
	keys := make([]string, len(numbermap))
	i := 0
	smallIndex := 1000
	for k := range numbermap {
		keys[i] = k
		i++
	}
	var dig string
	for _, k := range keys {
		if firstI := strings.Index(line, k); firstI >= 0 && firstI < smallIndex {
			smallIndex = firstI
			dig = k
			fmt.Println(line)
			//fmt.Println("Found in map check", smallIndex, "secondDigit")
		}
	}

	return smallIndex, numbermap[dig]

}
