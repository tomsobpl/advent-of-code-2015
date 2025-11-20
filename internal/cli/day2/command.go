package day2

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tomsobpl/advent-of-code-2015/internal"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli"
)

func getEdges(dimensions string) (int, int, int) {
	edges := internal.ConvertArrayOfStringsToArrayOfIntegers(strings.Split(dimensions, "x"))
	return edges[0], edges[1], edges[2]
}

func partOneCompute(input string) string {
	packages := internal.ConvertStringToArrayOfStrings(input)
	required := 0

	for _, dimensions := range packages {
		l, w, h := getEdges(dimensions)
		required += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}

	return fmt.Sprintf("%d", required)
}

func partTwoCompute(input string) string {
	packages := internal.ConvertStringToArrayOfStrings(input)
	required := 0

	for _, dimensions := range packages {
		l, w, h := getEdges(dimensions)
		required += l*w*h + min(2*(l+w), 2*(w+h), 2*(h+l))
	}

	return fmt.Sprintf("%d", required)
}

func NewCommand(partOneExpectedResult string, partTwoExpectedResult string) *cobra.Command {
	tsk := cli.AocTask{
		AocDay:                 2,
		PartOneExpectedResult:  partOneExpectedResult,
		PartOneComputeFunction: partOneCompute,
		PartTwoExpectedResult:  partTwoExpectedResult,
		PartTwoComputeFunction: partTwoCompute,
	}
	cmd := &cobra.Command{
		Use:   "day2",
		Short: "Solves day 2 of the advent of code 2015 (I Was Told There Would Be No Math)",
		Run: func(cmd *cobra.Command, args []string) {
			tsk.Solve(cmd.Flags().GetString(cli.FlagInputDataPath))
		},
	}

	return cmd
}
