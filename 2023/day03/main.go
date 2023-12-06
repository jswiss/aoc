package main

import (
	"fmt"
	"regexp"

	"github.com/jswiss/adventofcode/common/input"
)

type Symbol struct {
	line    int
	symbols []int
}

var symbols []Symbol

func GetSymbolIndices(idx int, input string) Symbol {
	pattern := `[^a-zA-Z0-9.]`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringIndex(input, -1)

	s := Symbol{line: 0, symbols: []int{}}
	s.line = idx
	for _, arr := range matches {
		s.symbols = append(s.symbols, arr[0])
	}
	return s

}

func main() {
	fileContent, err := input.ReadFileLines("input_test.txt")

	if err != nil {
		panic("cannot read input file")
	}

	// loop through and get index of every symbol on each line
	for i, val := range fileContent {
		idx := i
		indices := GetSymbolIndices(idx, val)
		symbols = append(symbols, indices)
	}
	fmt.Println(symbols)
	// loop through input and get the index and value of numbers on each line
	// get valid numbers per line
	// get valid numbers compared with symbol index on previous line  (diagonal)
	// get valid numbers compared with symbol index on next line  (diagonal)

}
