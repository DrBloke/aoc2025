package day02

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input string
	expected string
}

var testCases = []TestCase{
		{"hello world", ""},
		{"hello world", ""},
		}

// TestParser
// for a valid return value.
func TestParse(t *testing.T) {
	lastIndex := len(testCases) - 1
	data, err := Parse(testCases[lastIndex].input)
	expected := []Data{{word: "hello"}, {word: "world"}}
	if !reflect.DeepEqual(data, expected) {
		t.Errorf(`Parse("%s") = %v, %v, want %v, nil`, testCases[lastIndex].input, data, err, expected)
	}
	
	if err != nil {
		t.Errorf(`Parse("%s") returned unexpected error: %v`, testCases[lastIndex].input, err)
	}
}


