package main

import (
	"fmt"

	"github.com/jswiss/adventofcode/common/input"
)

func main() {
	fileContent, err := input.ReadFileLines("input_test.txt")

	if err != nil {
		panic("cannot read input file")
	}

	// loop through and get index of every symbol on each line
	// loop through input and get the index and value of numbers on each line
	// get valid numbers per line
	// get valid numbers compared with symbol index on previous line  (diagonal)
	// get valid numbers compared with symbol index on next line  (diagonal)

	fmt.Println(fileContent)
}
