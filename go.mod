module github.com/DrBloke/aoc2025

go 1.25.4

replace example.com/day01 => ./day01

require (
	example.com/day00 v0.0.0-00010101000000-000000000000
	example.com/day01 v0.0.0-00010101000000-000000000000
)

require github.com/oleiade/gomme v0.0.0-20231216113819-c8967c191356 // indirect

replace example.com/day00 => ./day00
