package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jswiss/adventofcode/common/input"
	"github.com/jswiss/adventofcode/common/utils"
)

const Red = 12
const Green = 13
const Blue = 14

func FindDigitBeforeSubstring(inputString, substring string) (int, error) {
	// Construct the regular expression pattern
	pattern := `(\d+)\s*` + substring

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find the match in the input string
	match := re.FindStringSubmatch(inputString)

	// Check if a match is found
	if len(match) >= 2 {
		// Return the captured digit
		digit, err := strconv.Atoi(match[1])

		if err != nil {
			return 0, err
		}

		// Return the integer value
		return digit, nil
	}

	// Return an error if no match is found
	return 0, fmt.Errorf("no match found for substring %s", substring)
}

func CheckPossibility(comparisons []int, max int) bool {
	for i := 0; i < len(comparisons); i++ {
		if comparisons[i] > max {
			return false
		}
	}
	return true
}

func GetMax(vals []int) int {
	max := vals[0]
	for i := range vals {
		if vals[i] > max {
			max = vals[i]
		}
	}
	return max
}

func main() {
	var scores []int
	fileContent, err := input.ReadFileLines("input.txt")

	if err != nil {
		panic("cannot read input file")
	}

	var sumPower int
	for _, line := range fileContent {
		var isRed, isBlue, isGreen bool
		var reds, greens, blues []int

		splitOnColon := strings.Split(line, ":")
		games := strings.Split(splitOnColon[1], ";")
		gameNum := strings.Split(splitOnColon[0], " ")
		gameInt, _ := strconv.Atoi(gameNum[1])

		for _, val := range games {
			r, _ := FindDigitBeforeSubstring(val, "red")
			if r != 0 {
				reds = append(reds, r)
			}
			g, _ := FindDigitBeforeSubstring(val, "green")
			if g != 0 {
				greens = append(greens, g)
			}
			b, _ := FindDigitBeforeSubstring(val, "blue")
			if b != 0 {
				blues = append(blues, b)
			}
			isRed = CheckPossibility(reds, Red)
			isBlue = CheckPossibility(blues, Blue)
			isGreen = CheckPossibility(greens, Green)

		}
		if isRed && isBlue && isGreen {
			scores = append(scores, gameInt)
		}
		minRed := GetMax(reds)
		minBlue := GetMax(blues)
		minGreen := GetMax(greens)
		sumPower += (minRed * minBlue * minGreen)
	}

	u := utils.Ints(scores)
	sum := utils.SliceSum(u)
	fmt.Println("Answer 1 is: ", sum)
	fmt.Println("Answer 2 is: ", sumPower)

}
