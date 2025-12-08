package main

import (
	"log"

	"example.com/day01"
)

func main() {
	log.SetPrefix("day01: ")
	log.SetFlags(0)
	day := "01" // TODO: get from command line args
	
	var data []day01.Operation
	var output []string
	var err error
	
	switch day {
	case "01":
		data, err = day01.Parse(day01.Input())
	default:
		log.Fatal("unknown day")
	}
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Parsed data:", data)
	
	switch day {
	case "01":
		output = []string{day01.SolveA(data), day01.SolveB(data)}
	default:
		log.Fatal("unknown day")
	}

	log.Println("Output:", output[0], output[1])
}