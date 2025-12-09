package main

import (
	"log"

	"example.com/day01"
)

func main() {
	log.SetPrefix("day01: ")
	log.SetFlags(0)
	day := "01" // TODO: get from command line args
	
	var output day01.Output
	
	switch day {
	case "01":
		output = day01.GetOutput()
	default:
		log.Fatal("unknown day")
	}

	log.Println("Parsed data:", output.ParsedData)
	log.Println("Output:", output.SolveA, output.SolveB)
}