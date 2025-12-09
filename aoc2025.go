package main

import (
	"log"

	"example.com/day00"
	"example.com/day01"
)

// DayOutput is an interface that all day outputs must implement
type DayOutput interface {
	GetSolveA() string
	GetSolveB() string
	GetParsedData() interface{}
}

func main() {
	day := "00" // TODO: get from command line args
	
	var output DayOutput
	
	switch day {
	case "00":
		output = day00.GetOutput()
	case "01":
		output = day01.GetOutput()
	default:
		log.Fatal("unknown day")
	}

	log.Println("Parsed data:", output.GetParsedData())
	log.Println("Output:", output.GetSolveA(), output.GetSolveB())
}