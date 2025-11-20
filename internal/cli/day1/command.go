package day1

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli"
)

func partOneCompute(input string) string {
	currentFloor := 0

	for _, char := range input {
		if char == '(' {
			currentFloor++
		} else if char == ')' {
			currentFloor--
		}
	}

	return fmt.Sprintf("%d", currentFloor)
}

func partTwoCompute(input string) string {
	currentFloor := 0

	for i, char := range input {
		if char == '(' {
			currentFloor++
		} else if char == ')' {
			currentFloor--
		}

		if currentFloor == -1 {
			return fmt.Sprintf("%d", i+1)
		}
	}

	return fmt.Sprintf("%d", 0)
}

func NewCommand(partOneExpectedResult string, partTwoExpectedResult string) *cobra.Command {
	tsk := cli.AocTask{
		AocDay:                 1,
		PartOneExpectedResult:  partOneExpectedResult,
		PartOneComputeFunction: partOneCompute,
		PartTwoExpectedResult:  partTwoExpectedResult,
		PartTwoComputeFunction: partTwoCompute,
	}
	cmd := &cobra.Command{
		Use:   "day1",
		Short: "Solves day 1 of the advent of code 2015 (Not Quite Lisp)",
		Run: func(cmd *cobra.Command, args []string) {
			tsk.Solve(cmd.Flags().GetString(cli.FlagInputDataPath))
		},
	}

	return cmd
}
