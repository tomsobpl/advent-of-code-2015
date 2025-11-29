package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tomsobpl/advent-of-code-2015/internal"
	"github.com/tomsobpl/advent-of-code-2015/internal/cli"
)

type GateInstruction interface {
	Execute() uint16
}

type DirectAssignmentInstruction struct {
	Instruction string
	Signal      string
}

func (i DirectAssignmentInstruction) Compute(signal uint16) uint16 {
	return signal
}

func (i DirectAssignmentInstruction) Execute() uint16 {
	return i.Compute(decodeInputSignal(i.Signal))
}

type BitwiseAndInstruction struct {
	Instruction string
	SignalLeft  string
	SignalRight string
}

func (i BitwiseAndInstruction) Compute(signalLeft uint16, signalRight uint16) uint16 {
	return signalLeft & signalRight
}

func (i BitwiseAndInstruction) Execute() uint16 {
	return i.Compute(decodeInputSignal(i.SignalLeft), decodeInputSignal(i.SignalRight))
}

type BitwiseNotInstruction struct {
	Instruction string
	Signal      string
}

func (i BitwiseNotInstruction) Compute(signal uint16) uint16 {
	return ^signal
}

func (i BitwiseNotInstruction) Execute() uint16 {
	return i.Compute(decodeInputSignal(i.Signal))
}

type BitwiseOrInstruction struct {
	Instruction string
	SignalLeft  string
	SignalRight string
}

func (i BitwiseOrInstruction) Compute(signalLeft uint16, signalRight uint16) uint16 {
	return signalLeft | signalRight
}

func (i BitwiseOrInstruction) Execute() uint16 {
	return i.Compute(decodeInputSignal(i.SignalLeft), decodeInputSignal(i.SignalRight))
}

type BitshiftLeftInstruction struct {
	Instruction string
	SignalLeft  string
	SignalRight string
}

func (i BitshiftLeftInstruction) Compute(signalLeft uint16, signalRight uint16) uint16 {
	return signalLeft << signalRight
}

func (i BitshiftLeftInstruction) Execute() uint16 {
	return i.Compute(decodeInputSignal(i.SignalLeft), decodeInputSignal(i.SignalRight))
}

type BitshiftRightInstruction struct {
	Instruction string
	SignalLeft  string
	SignalRight string
}

func (i BitshiftRightInstruction) Compute(signalLeft uint16, signalRight uint16) uint16 {
	return signalLeft >> signalRight
}

func (i BitshiftRightInstruction) Execute() uint16 {
	return i.Compute(decodeInputSignal(i.SignalLeft), decodeInputSignal(i.SignalRight))
}

func decodeInputSignal(signal string) uint16 {
	if value, exists := CircuitCache[signal]; exists {
		return value
	}

	if i, err := strconv.Atoi(signal); err == nil {
		return uint16(i)
	}

	CircuitCache[signal] = Circuit[signal].Execute()
	return CircuitCache[signal]
}

var (
	Circuit                 = make(map[string]GateInstruction)
	CircuitCache            = make(map[string]uint16)
	BinaryGateSignalPattern = regexp.MustCompile(`\w+ (AND|OR|LSHIFT|RSHIFT) \w+`)
	NotSignalPattern        = regexp.MustCompile(`NOT \w+`)
)

func decodeInputLine(line string) (string, string) {
	parts := strings.Split(line, "->")
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func createGateInstruction(instruction string, rawInstruction string) GateInstruction {
	// Check if it's negation (NOT x -> y)
	if NotSignalPattern.MatchString(instruction) {
		return BitwiseNotInstruction{
			Instruction: rawInstruction,
			Signal:      strings.Fields(instruction)[1],
		}
	}

	// Check if it's binary gate (x AND y -> z; x OR y -> z; x LSHIFT y -> z; x RSHIFT y -> z)
	if BinaryGateSignalPattern.MatchString(instruction) {
		parts := strings.Fields(instruction)

		switch parts[1] {
		case "AND":
			return BitwiseAndInstruction{
				Instruction: rawInstruction,
				SignalLeft:  parts[0],
				SignalRight: parts[2],
			}
		case "OR":
			return BitwiseOrInstruction{
				Instruction: rawInstruction,
				SignalLeft:  parts[0],
				SignalRight: parts[2],
			}
		case "LSHIFT":
			return BitshiftLeftInstruction{
				Instruction: rawInstruction,
				SignalLeft:  parts[0],
				SignalRight: parts[2],
			}
		case "RSHIFT":
			return BitshiftRightInstruction{
				Instruction: rawInstruction,
				SignalLeft:  parts[0],
				SignalRight: parts[2],
			}
		}
	}

	return DirectAssignmentInstruction{
		Instruction: rawInstruction,
		Signal:      instruction,
	}
}

func partOneCompute(input string) string {
	Circuit = make(map[string]GateInstruction)
	CircuitCache = make(map[string]uint16)

	for _, line := range internal.ConvertStringToArrayOfStrings(input) {
		instruction, wire := decodeInputLine(line)
		Circuit[wire] = createGateInstruction(instruction, line)
	}

	if _, exists := Circuit["a"]; exists {
		return fmt.Sprint(Circuit["a"].Execute())
	}

	return fmt.Sprint("NA")
}

func partTwoCompute(input string) string {
	partOneResult := partOneCompute(input)

	Circuit = make(map[string]GateInstruction)
	CircuitCache = make(map[string]uint16)

	for _, line := range internal.ConvertStringToArrayOfStrings(input) {
		instruction, wire := decodeInputLine(line)
		Circuit[wire] = createGateInstruction(instruction, line)
	}

	Circuit["b"] = DirectAssignmentInstruction{
		Instruction: fmt.Sprintf("%s -> b", partOneResult),
		Signal:      partOneResult,
	}

	if _, exists := Circuit["a"]; exists {
		return fmt.Sprint(Circuit["a"].Execute())
	}

	return fmt.Sprint("NA")
}

func NewCommand(partOneExpectedResult string, partTwoExpectedResult string) *cobra.Command {
	tsk := cli.AocTask{
		AocDay:                 7,
		PartOneExpectedResult:  partOneExpectedResult,
		PartOneComputeFunction: partOneCompute,
		PartTwoExpectedResult:  partTwoExpectedResult,
		PartTwoComputeFunction: partTwoCompute,
	}
	cmd := &cobra.Command{
		Use:   "day7",
		Short: "Solves day 7 of the advent of code 2015 (Some Assembly Required)",
		Run: func(cmd *cobra.Command, args []string) {
			tsk.Solve(cmd.Flags().GetString(cli.FlagInputDataPath))
		},
	}

	return cmd
}
