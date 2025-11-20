package cli

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const FlagInputDataPath string = "input-data-path"

type compute func(input string) string

type AocTask struct {
	AocDay                 int
	AocInputDataPath       string
	PartOneExpectedResult  string
	PartOneComputeFunction compute
	PartTwoExpectedResult  string  `optional:"true"`
	PartTwoComputeFunction compute `optional:"true"`
}

func (t AocTask) Solve(inputDataPath string, _ error) {
	t.AocInputDataPath = inputDataPath
	t.checkAndCompute(t.PartOneExpectedResult, t.PartOneComputeFunction, 1)
	t.checkAndCompute(t.PartTwoExpectedResult, t.PartTwoComputeFunction, 2)
}

func (t AocTask) checkAndCompute(expectedResult string, computeFunction compute, part int) {
	if expectedResult == "" || computeFunction == nil {
		return
	}

	start := time.Now()
	actualResult := computeFunction(t.getRawTestInput("test", part))
	testDuration := time.Since(start)

	if actualResult == expectedResult {
		fmt.Printf("Part %d result matches expected result (took %v). Solving live data...\n", part, testDuration)
		start := time.Now()
		result := computeFunction(t.getRawTestInput("live", part))
		liveDuration := time.Since(start)
		fmt.Printf("Result: %s (took %v)\n", result, liveDuration)
	} else {
		fmt.Printf("Part %d result '%s' does not match expected result '%s'\n", part, actualResult, expectedResult)
		os.Exit(1)
	}
}

func (t AocTask) getRawTestInput(kind string, part int) string {
	filePath := fmt.Sprintf("%s/day%d_part%d_%s.txt", t.AocInputDataPath, t.AocDay, part, kind)
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(content))
}
