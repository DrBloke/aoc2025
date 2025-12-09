package main

import (
	"log"
	"os"

	"example.com/day00"
	"example.com/day01"
	"example.com/day02"
	"example.com/day03"
	"example.com/day04"
	"example.com/day05"
	"example.com/day06"
	"example.com/day07"
	"example.com/day08"
	"example.com/day09"
	"example.com/day10"
	"example.com/day11"
	"example.com/day12"
	"example.com/day13"
	"example.com/day14"
)

// DayOutput is an interface that all day outputs must implement
type DayOutput interface {
	GetSolveA() string
	GetSolveB() string
	GetParsedData() string // TODO: if this was a string type then we could get rid of the any type
}

func main() {
	day := "00"
	if len(os.Args) > 1 {
		day = os.Args[1]
	}
	
	var output DayOutput

	switch day {
	case "00":
		output = day00.GetOutput()
	case "01":
		output = day01.GetOutput()
	case "02":
		output = day02.GetOutput()
	case "03":
		output = day03.GetOutput()
	case "04":
		output = day04.GetOutput()
	case "05":
		output = day05.GetOutput()
	case "06":
		output = day06.GetOutput()
	case "07":
		output = day07.GetOutput()
	case "08":
		output = day08.GetOutput()
	case "09":
		output = day09.GetOutput()
	case "10":
		output = day10.GetOutput()
	case "11":
		output = day11.GetOutput()
	case "12":
		output = day12.GetOutput()
	case "13":
		output = day13.GetOutput()
	case "14":
		output = day14.GetOutput()
	default:
		log.Fatal("unknown day")
	}

	log.Println("Parsed data:", output.GetParsedData())
	log.Println("Output:", output.GetSolveA(), output.GetSolveB())
}