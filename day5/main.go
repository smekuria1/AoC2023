package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	Part2()
	fmt.Printf("P2 took %v\n", time.Since(start))
}

var re = regexp.MustCompile("[0-9]+")

type Seed struct {
	val int
}

type Ranges struct {
	dRangeStart int
	sRangeStart int
	rangeLen    int
}

var mappers [][]Ranges = make([][]Ranges, 7)
var minLocation int = math.MaxInt

//var mut sync.Mutex

func Part1() {
	lines := readlines()
	clean := strings.Split(strings.Join(lines, "\n"), "\n\n")
	var seeds []Seed
	seedString := clean[0]
	seedNumbers := re.FindAllString(seedString, -1)
	fmt.Printf("seedNumbers: %v\n", seedNumbers)

	for _, seedNum := range seedNumbers {
		seedNumber, err := strconv.Atoi(seedNum)
		check(err)

		seeds = append(seeds, Seed{
			seedNumber,
		})

	}

	for index, split := range clean[1:] {
		buildmap(split, index+1)
	}

	for _, seed := range seeds {
		updateminloc(seed.val, 0)
	}

	//fmt.Printf("mappers: %v\n", mappers)
	fmt.Printf("minLocation: %v\n", minLocation)
}

func Part2() {
	lines := readlines()
	clean := strings.Split(strings.Join(lines, "\n"), "\n\n")
	var seeds []Seed
	seedString := clean[0]
	seedNumbers := re.FindAllString(seedString, -1)
	fmt.Printf("seedNumbers: %v\n", seedNumbers)
	for _, seedNum := range seedNumbers {
		seedNumber, err := strconv.Atoi(seedNum)
		check(err)

		seeds = append(seeds, Seed{
			seedNumber,
		})

	}

	var seedsPairs [][]Seed
	for i := 0; i < len(seeds); i += 2 {
		seedsPairs = append(seedsPairs, []Seed{seeds[i], seeds[i+1]})
	}

	for index, split := range clean[1:] {
		buildmap(split, index+1)
	}

	for _, seedPair := range seedsPairs {
		fmt.Println("Calculating for seed pair", seedPair[0].val, seedPair[1].val)
		runForSeedPair(seedPair)
	}
	fmt.Printf("minLocation: %v\n", minLocation)
}

func runForSeedPair(seedPair []Seed) {
	for i := seedPair[0].val; i <= seedPair[0].val+seedPair[1].val-1; i++ {
		updateminloc(Seed{i}.val, 0)
	}

}

func updateminloc(source int, index int) {

	n := source
	iterover := mappers[index]
	for _, mapper := range iterover {
		if isBetween(source, mapper.sRangeStart, mapper.sRangeStart+mapper.rangeLen-1) {
			n = source + (mapper.dRangeStart - mapper.sRangeStart)
		}
	}

	if index < 6 {
		updateminloc(n, index+1)
	} else {
		if n < minLocation {

			minLocation = n

		}
	}
}

func isBetween(num, min, max int) bool {
	return num >= min && num <= max
}

func buildmap(clean string, index int) {
	elems := strings.Split(clean, "\n")[1:]
	builder(elems, &mappers[index-1])

}

func builder(lines []string, maptoUpdate *[]Ranges) {
	for _, line := range lines {
		elem := re.FindAllString(line, -1)

		dRangeStart, err := strconv.Atoi(elem[0])
		check(err)
		sRangeStart, err := strconv.Atoi(elem[1])
		check(err)
		rangeLen, err := strconv.Atoi(elem[2])
		check(err)

		*maptoUpdate = append(*maptoUpdate, Ranges{
			dRangeStart,
			sRangeStart,
			rangeLen,
		})
	}
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
