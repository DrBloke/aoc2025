package day01

import (
	"log"
	"os"
	"strconv"

	"github.com/oleiade/gomme"
)

type Operation struct {
	Direction string
	Clicks    int}

func Input() string { 	
	data, err := os.ReadFile("day01/day01.txt")
	if err != nil {
		log.Fatal("Could not read input file")
	}
	return string(data)
 }

type Output struct {
	ParsedData []Operation
	SolveA string
	SolveB string
}

func GetOutput() Output {
	input := Input()
	data, err := Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	return Output{
		ParsedData: data,
		SolveA: SolveA(data),
		SolveB: SolveB(data),
	}
}

// GetSolveA implements the DayOutput interface
func (o Output) GetSolveA() string {
	return o.SolveA
}

// GetSolveB implements the DayOutput interface
func (o Output) GetSolveB() string {
	return o.SolveB
}

// GetParsedData implements the DayOutput interface
func (o Output) GetParsedData() interface{} {
	return o.ParsedData
}

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
			Clicks:    int(clicks),
		}, nil
	})
	step3 := gomme.Many1(step2)

	result := step3(input)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Output, nil
}

func SolveA(input []Operation) string {
	position := 50
	zeros := 0
	for _, op := range input {
		if op.Direction == "L" {
			position = position - op.Clicks
		} else {
			position = position + op.Clicks
		}
		
		// Normalize position to 0-99 range
		position = ((position % 100) + 100) % 100
		
		if position == 0 {
			zeros++
		}
	}
	return strconv.Itoa(zeros)
}

func SolveB(input []Operation) string {
	position := 50
	zeros := 0
	for _, op := range input {
		if op.Direction == "L" {
			if position - op.Clicks < 0 {
				if position != 0 {
					zeros++
				}
				remainingClicks := op.Clicks - position
				zeros += remainingClicks / 100
				position = - remainingClicks % 100
				if position <0 {
					position = 100 + position
				}
			} else {
				position = (position - op.Clicks) % 100
			if position == 0 {
				zeros++
			}
			}
		} else {
			if position + op.Clicks > 100 {
				zeros++
				remainingClicks := op.Clicks - (100 - position)
				zeros += remainingClicks / 100
				position = remainingClicks % 100
			} else {
				position = (position + op.Clicks) % 100
			if position == 0 {
				zeros++
			}
			}
		}
	}
	return strconv.Itoa(zeros)
}