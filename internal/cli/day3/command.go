package day3

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomsobpl/advent-of-code-2015/internal"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli"
)

const (
	North = '^'
	East  = '>'
	South = 'v'
	West  = '<'
)

func markPositionOnGrid(grid map[string]int, position map[string]int) {
	_tmp := fmt.Sprintf("%d,%d", position["x"], position["y"])

	_, exists := grid[_tmp]
	if !exists {
		grid[_tmp] = 1
	} else {
		grid[_tmp]++
	}

	grid[_tmp] = 1
}

func nextPosition(position map[string]int, direction rune) map[string]int {
	switch direction {
	case North:
		return map[string]int{"x": position["x"], "y": position["y"] + 1}
	case East:
		return map[string]int{"x": position["x"] + 1, "y": position["y"]}
	case South:
		return map[string]int{"x": position["x"], "y": position["y"] - 1}
	case West:
		return map[string]int{"x": position["x"] - 1, "y": position["y"]}
	default:
		panic(fmt.Sprintf("Unknown direction: %c", direction))
	}
}

func advancePositionOnGrid(grid map[string]int, position map[string]int, direction rune) map[string]int {
	_tmp := nextPosition(position, direction)
	markPositionOnGrid(grid, _tmp)
	return _tmp
}

func partOneCompute(input string) string {
	steps := internal.ConvertStringToArrayOfRunes(input)

	grid := map[string]int{"0,0": 1}
	curr := map[string]int{"x": 0, "y": 0}

	for _, direction := range steps {
		curr = advancePositionOnGrid(grid, curr, direction)
	}

	return fmt.Sprint(len(grid))
}

func partTwoCompute(input string) string {
	steps := internal.ConvertStringToArrayOfRunes(input)

	grid := map[string]int{"0,0": 2}
	curr := map[string]int{"x": 0, "y": 0}
	robo := map[string]int{"x": 0, "y": 0}

	for i := 0; i < len(steps)-1; i += 2 {
		curr = advancePositionOnGrid(grid, curr, steps[i])
		robo = advancePositionOnGrid(grid, robo, steps[i+1])
	}

	if len(steps)%2 != 0 {
		lastIndex := len(steps) - 1
		curr = advancePositionOnGrid(grid, curr, steps[lastIndex])
	}

	return fmt.Sprint(len(grid))
}

func NewCommand(partOneExpectedResult string, partTwoExpectedResult string) *cobra.Command {
	tsk := cli.AocTask{
		AocDay:                 3,
		PartOneExpectedResult:  partOneExpectedResult,
		PartOneComputeFunction: partOneCompute,
		PartTwoExpectedResult:  partTwoExpectedResult,
		PartTwoComputeFunction: partTwoCompute,
	}
	cmd := &cobra.Command{
		Use:   "day3",
		Short: "Solves day 3 of the advent of code 2015 (Perfectly Spherical Houses in a Vacuum)",
		Run: func(cmd *cobra.Command, args []string) {
			tsk.Solve(cmd.Flags().GetString(cli.FlagInputDataPath))
		},
	}

	return cmd
}
