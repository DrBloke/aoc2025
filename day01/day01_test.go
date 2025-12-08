package day01

import (
	"reflect"
	"strconv"
	"testing"
)

type TestCase struct {
	input string
	expected int
}

var testCases = []TestCase{
	{"L50", 1},
	{"R50", 1},
	{"R100", 1},
	{"R150", 2},
	{"R49", 0},
	{"R51", 1},
	{"R300", 3},
	{"R350", 4},
	{"L49", 0},
	{"L51", 1},
	{"L300", 3},
	{"L350", 4},
	{"L50 R100", 2},
	{"R50 L100", 2},
	{"L50 R300", 4},
	{"R50 L350", 4},
	{"L25 L8 L44 L6 R1 R26 R9", 2},
	{"L68 L30 R48 L5 R60 L55 L1 L99 R14 L82", 6},
	}

// TestParser
// for a valid return value.
func TestParse(t *testing.T) {
	lastIndex := len(testCases) - 1
	data, err := Parse(testCases[lastIndex].input)
	expected := []Operation{{Direction: "L", Clicks: 68}, {Direction: "L", Clicks: 30}, {Direction: "R", Clicks: 48}, {Direction: "L", Clicks: 5}, {Direction: "R", Clicks: 60}, {Direction: "L", Clicks: 55}, {Direction: "L", Clicks: 1}, {Direction: "L", Clicks: 99}, {Direction: "R", Clicks: 14}, {Direction: "L", Clicks: 82}}
	if !reflect.DeepEqual(data, expected) {
		t.Errorf(`Parse("%s") = %v, %v, want %v, nil`, testCases[lastIndex].input, data, err, expected)
	}
	
	if err != nil {
		t.Errorf(`Parse("%s") returned unexpected error: %v`, testCases[lastIndex].input, err)
	}
}

func TestSolveB(t *testing.T) {
	for _, test := range testCases {
		data, err := Parse(test.input)
		if err != nil {
			t.Errorf(`Parse("%s") returned unexpected error: %v`, test.input, err)
		}
		output := SolveB(data)
		expected := strconv.Itoa(test.expected)
		if output != expected {
			t.Errorf(`SolveB("%s") = %v, want %v`, test.input, output, expected)
		}
	}
}