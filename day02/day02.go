package day02

import (
	"fmt"
	"log"
	"os"

	"github.com/oleiade/gomme"
)

type Data struct {
	word string
}

func Input() string { 	
	data, err := os.ReadFile("day02/day02.txt")
	if err != nil {
		log.Fatal("Could not read input file")
	}
	return string(data)
 }

type Output struct {
	ParsedData []Data
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
func (o Output) GetParsedData() string {
	return fmt.Sprintf("%v", o.ParsedData)
}

func Parse(input string) ([]Data, error) {
	step1 := gomme.Sequence(
		gomme.Alphanumeric1[string](),
		gomme.Whitespace0[string](),
	)
	step2 := gomme.Map(step1, func(seq []string) (Data, error) {
		return Data{
			word: seq[0],
		}, nil
	})
	step3 := gomme.Many1(step2)

	result := step3(input)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Output, nil
}

func SolveA(input []Data) string {
	return "TODO"
}

func SolveB(input []Data) string {
	return "TODO"
}