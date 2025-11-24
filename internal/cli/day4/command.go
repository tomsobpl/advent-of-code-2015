package day4

import (
	"crypto/md5"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli"
)

func partOneCompute(input string) string {
	for i := 1; ; i++ {
		data := []byte(fmt.Sprintf("%s%d", input, i))
		hash := md5.Sum(data)
		if hash[0] == 0 && hash[1] == 0 && hash[2] < 0x10 {
			return fmt.Sprint(i)
		}
	}
}

func partTwoCompute(input string) string {
	for i := 1; ; i++ {
		data := []byte(fmt.Sprintf("%s%d", input, i))
		hash := md5.Sum(data)
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			return fmt.Sprint(i)
		}
	}
}

func NewCommand(partOneExpectedResult string, partTwoExpectedResult string) *cobra.Command {
	tsk := cli.AocTask{
		AocDay:                 4,
		PartOneExpectedResult:  partOneExpectedResult,
		PartOneComputeFunction: partOneCompute,
		PartTwoExpectedResult:  partTwoExpectedResult,
		PartTwoComputeFunction: partTwoCompute,
	}
	cmd := &cobra.Command{
		Use:   "day4",
		Short: "Solves day 4 of the advent of code 2015 (The Ideal Stocking Stuffer)",
		Run: func(cmd *cobra.Command, args []string) {
			tsk.Solve(cmd.Flags().GetString(cli.FlagInputDataPath))
		},
	}

	return cmd
}
