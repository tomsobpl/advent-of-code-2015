package internal

import (
	"strconv"
	"strings"
)

func ConvertStringToArrayOfRunes(input string) []rune {
	return []rune(input)
}

func ConvertStringToArrayOfStrings(input string) []string {
	return strings.Split(input, "\n")
}

func ConvertArrayOfStringsToArrayOfIntegers(input []string) []int {
	result := make([]int, len(input))

	for i, v := range input {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

func NewGridOfIntegers(rows int, cols int) [][]int {
	grid := make([][]int, rows)

	for i := 0; i < rows; i++ {
		grid[i] = make([]int, cols)
	}

	return grid
}
