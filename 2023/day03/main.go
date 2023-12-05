package main

import (
	"fmt"
	"regexp"

	"github.com/jswiss/adventofcode/common/input"
)

func GetSymbolIndices(input string) []int {
	pattern := `[^a-zA-Z0-9.]`
	indices := []int{}
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringIndex(input, -1)

	for _, arr := range matches {
		indices = append(indices, arr[0])
	}
	return indices

}

func main() {
	fileContent, err := input.ReadFileLines("input_test.txt")

	if err != nil {
		panic("cannot read input file")
	}

	// loop through and get index of every symbol on each line
	symbolIndices := []int{}
	for _, val := range fileContent {
		indices := GetSymbolIndices(val)
		fmt.Println("Indices: ", indices)
		symbolIndices = append(symbolIndices, indices...)
	}
	fmt.Println(symbolIndices)
	// loop through input and get the index and value of numbers on each line
	// get valid numbers per line
	// get valid numbers compared with symbol index on previous line  (diagonal)
	// get valid numbers compared with symbol index on next line  (diagonal)

}
