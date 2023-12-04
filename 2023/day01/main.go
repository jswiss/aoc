package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jswiss/adventofcode/common/input"
	"github.com/jswiss/adventofcode/common/utils"
)

func getNums(line string) (int, error) {
	re := regexp.MustCompile(`(\d)`)
	matches := re.FindAllStringSubmatch(line, -1)

	if len(matches) < 1 {
		return 0, fmt.Errorf("no single-digit integers found in the input string")
	}

	firstSingleDigit := matches[0][1]

	lastIndex := len(matches) - 1
	lastSingleDigit := matches[lastIndex][1]

	strInt := firstSingleDigit + lastSingleDigit

	res, err := strconv.Atoi(strInt)
	if err != nil {
		fmt.Errorf("there is an error bruv")
	}
	return res, nil

}

func ReplaceSubstrings(line string) string {
	strNums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	// handle edge

	for key, val := range strNums {
		idx := strings.Index(line, key)

		if idx != -1 {
			before := line[:idx+1]
			after := line[idx+1:]
			line = before + strconv.Itoa(val) + after
		}
	}
	return line
}

func first(content []string) int {
	var nums []int
	// loop over each line in the file
	for _, val := range content {
		num, err := getNums(val)
		if err != nil {
			fmt.Errorf("problem getting nums")
		}
		nums = append(nums, num)
	}
	return utils.SliceSum(nums)
}

func second(content []string) int {
	var nums []int
	// loop over each line in the file
	for _, val := range content {
		res := ReplaceSubstrings(val)
		num, err := getNums(res)
		if err != nil {
			fmt.Errorf("problem getting nums")
		}
		nums = append(nums, num)
	}
	fmt.Println(nums)
	return utils.SliceSum(nums)
}

func main() {
	fileContent, err := input.ReadFileLines("input.txt")

	if err != nil {
		panic("cannot read input file")
	}

	res := first(fileContent)

	fmt.Println("first answer: ", res)

	res2 := second(fileContent)
	fmt.Println("second answer:", res2)

}
