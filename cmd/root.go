/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day1"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day2"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli/day3"
)

// rootCmd represents the base command when called without any subcommands
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
}
