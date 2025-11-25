package day6

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tomsobpl/advent-of-code-2015/internal"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli"
)

const (
	TurnOn  = "turn on"
	TurnOff = "turn off"
	Toggle  = "toggle"
)

func decodeInstructionLine(line string) (string, []int, []int) {
	re := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	matches := re.FindStringSubmatch(line)

	if matches == nil {
		return "", []int{}, []int{}
	}

	startX, _ := strconv.Atoi(matches[2])
	startY, _ := strconv.Atoi(matches[3])
	endX, _ := strconv.Atoi(matches[4])
	endY, _ := strconv.Atoi(matches[5])

	return matches[1], []int{startX, startY}, []int{endX, endY}
}

func applyToGrid(grid [][]int, start []int, end []int, action string) {
	for x := start[0]; x <= end[0]; x++ {
		for y := start[1]; y <= end[1]; y++ {
			switch action {
			case TurnOn:
				grid[x][y] = 1
			case TurnOff:
				grid[x][y] = 0
			case Toggle:
				grid[x][y] = toggle(grid[x][y])
			}
		}
	}
}

func applyToGridWithBrightness(grid [][]int, start []int, end []int, action string) {
	for x := start[0]; x <= end[0]; x++ {
		for y := start[1]; y <= end[1]; y++ {
			switch action {
			case TurnOn:
				grid[x][y] += 1
			case TurnOff:
				grid[x][y] = decrementOrZero(grid[x][y])
			case Toggle:
				grid[x][y] += 2
			}
		}
	}
}

func toggle(input int) int {
	if input == 0 {
		return 1
	}

	return 0
}

func reduce(input [][]int) int {
	_sum := 0

	for _, row := range input {
		for _, value := range row {
			_sum += value
		}
	}

	return _sum
}

func decrementOrZero(input int) int {
	if input == 0 {
		return 0
	}

	return input - 1
}

func partOneCompute(input string) string {
	grid := internal.NewGridOfIntegers(1000, 1000)

	for _, line := range internal.ConvertStringToArrayOfStrings(input) {
		command, start, end := decodeInstructionLine(line)
		applyToGrid(grid, start, end, command)
	}

	return fmt.Sprintf("%d", reduce(grid))
}

func partTwoCompute(input string) string {
	grid := internal.NewGridOfIntegers(1000, 1000)

	for _, line := range internal.ConvertStringToArrayOfStrings(input) {
		command, start, end := decodeInstructionLine(line)
		applyToGridWithBrightness(grid, start, end, command)
	}

	return fmt.Sprintf("%d", reduce(grid))
}

func NewCommand(partOneExpectedResult string, partTwoExpectedResult string) *cobra.Command {
	tsk := cli.AocTask{
		AocDay:                 6,
		PartOneExpectedResult:  partOneExpectedResult,
		PartOneComputeFunction: partOneCompute,
		PartTwoExpectedResult:  partTwoExpectedResult,
		PartTwoComputeFunction: partTwoCompute,
	}
	cmd := &cobra.Command{
		Use:   "day6",
		Short: "Solves day 6 of the advent of code 2015 (Probably a Fire Hazard)",
		Run: func(cmd *cobra.Command, args []string) {
			tsk.Solve(cmd.Flags().GetString(cli.FlagInputDataPath))
		},
	}

	return cmd
}
