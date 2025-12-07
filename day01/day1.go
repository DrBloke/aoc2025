package day01

import (
	"strconv"

	"github.com/oleiade/gomme"
)

type Operation struct {
	Direction string
	Clicks    int8
}

func Input() string { return "L78 R30 L10 R20" }

func Parse(input string) ([]Operation, error) {
	step1 := gomme.Sequence(
		gomme.Alternative(gomme.Token[string]("L"), gomme.Token[string]("R")),
		gomme.Digit1[string](),
		gomme.Whitespace0[string](),
	)
	step2 := gomme.Map(step1, func(seq []string) (Operation, error) {
		clicks, err := strconv.Atoi(seq[1])
		if err != nil {
			return Operation{}, err
		}
		return Operation{
			Direction: seq[0],
			Clicks:    int8(clicks),
		}, nil
	})
	step3 := gomme.Many1(step2)

	result := step3(input)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Output, nil
}

func Solve(input []Operation) (string) {
	return ""
}