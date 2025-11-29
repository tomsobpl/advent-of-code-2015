package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day1"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day2"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day3"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day4"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day5"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day6"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day7"
)

var rootCmd = &cobra.Command{
	Use:   "advent-of-code-2015",
	Short: "Advent of code 2015 solutions in Go",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Lookup(cli.FlagInputDataPath).Value.String() == "" {
			cmd.PrintErrf("%s is required", cli.FlagInputDataPath)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP(cli.FlagInputDataPath, "i", "", "Path to the input data directory")
	rootCmd.AddCommand(day1.NewCommand("-3", "5"))
	rootCmd.AddCommand(day2.NewCommand("101", "48"))
	rootCmd.AddCommand(day3.NewCommand("2", "11"))
	rootCmd.AddCommand(day4.NewCommand("609043", "6742839"))
	rootCmd.AddCommand(day5.NewCommand("2", "2"))
	rootCmd.AddCommand(day6.NewCommand("998996", "1001996"))
	rootCmd.AddCommand(day7.NewCommand("NA", "NA"))
}
