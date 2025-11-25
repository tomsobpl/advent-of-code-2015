package day5

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/tomsobpl/advent-of-code-2015/internal"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli"
)

func calculateStringValue(s string) int {
	if len(s) < 3 {
		return 0
	}

	// A nice string is one with all of the following properties:
	// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
	if regexp.MustCompile("(ab|cd|pq|xy)").MatchString(s) {
		return 0
	}

	// It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
	vowelCount := len(regexp.MustCompile("[aeiou]").FindAllString(s, -1))

	// It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
	hasDoubleLetter := false
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			hasDoubleLetter = true
			break
		}
	}

	if vowelCount >= 3 && hasDoubleLetter {
		return 1
	}

	return 0
}

func calculateCorrectedStringValue(s string) int {
	if len(s) < 4 {
		return 0
	}

	// Now, a nice string is one with all of the following properties:
	// It contains a pair of any two letters that appears at least twice in the string without overlapping,
	// like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
	hasRepeatedLetterPair := false

	runes := []rune(s)

	for i := 0; i < len(runes)-1; i++ {
		if len(regexp.MustCompile(string(runes[i])+string(runes[i+1])).FindAllString(s, -1)) >= 2 {
			hasRepeatedLetterPair = true
			break
		}
	}

	if !hasRepeatedLetterPair {
		return 0
	}

	// It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe),
	// or even aaa.
	hasRepeatedLetterAroundOther := false
	for i := 0; i < len(runes)-2; i++ {
		if runes[i] == runes[i+2] {
			hasRepeatedLetterAroundOther = true
			break
		}
	}

	if hasRepeatedLetterAroundOther {
		return 1
	}

	return 0
}

func partOneCompute(input string) string {
	validated := 0

	for _, s := range internal.ConvertStringToArrayOfStrings(input) {
		validated += calculateStringValue(s)
	}

	return fmt.Sprintf("%d", validated)
}

func partTwoCompute(input string) string {
	validated := 0

	for _, s := range internal.ConvertStringToArrayOfStrings(input) {
		validated += calculateCorrectedStringValue(s)
	}

	return fmt.Sprintf("%d", validated)
}

func NewCommand(partOneExpectedResult string, partTwoExpectedResult string) *cobra.Command {
	tsk := cli.AocTask{
		AocDay:                 5,
		PartOneExpectedResult:  partOneExpectedResult,
		PartOneComputeFunction: partOneCompute,
		PartTwoExpectedResult:  partTwoExpectedResult,
		PartTwoComputeFunction: partTwoCompute,
	}
	cmd := &cobra.Command{
		Use:   "day5",
		Short: "Solves day 5 of the advent of code 2015 (Doesn't He Have Intern-Elves For This?)",
		Run: func(cmd *cobra.Command, args []string) {
			tsk.Solve(cmd.Flags().GetString(cli.FlagInputDataPath))
		},
	}

	return cmd
}
