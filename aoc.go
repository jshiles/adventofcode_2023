package main

import (
	"fmt"
	"os"

	"github.com/jshiles/adventofcode_2023/aoc/day08"
	"github.com/jshiles/adventofcode_2023/aoc/day09"
	"github.com/jshiles/adventofcode_2023/aoc/day10"
	"github.com/jshiles/adventofcode_2023/aoc/day11"
	"github.com/jshiles/adventofcode_2023/aoc/day12"
	"github.com/jshiles/adventofcode_2023/aoc/day13"
	"github.com/jshiles/adventofcode_2023/aoc/day14"
	"github.com/jshiles/adventofcode_2023/aoc/day15"
	"github.com/jshiles/adventofcode_2023/aoc/day16"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <day> [filename]")
		return
	}

	day := args[1]
	filename := args[2]

	switch day {
	case "8", "08":
		day08.Run(filename)
	case "9", "09":
		day09.Run(filename)
	case "10":
		day10.Run(filename)
	case "11":
		day11.Run(filename)
	case "12":
		day12.Run(filename)
	case "13":
		day13.Run(filename)
	case "14":
		day14.Run(filename)
	case "15":
		day15.Run(filename)
	case "16":
		day16.Run(filename)
	default:
		fmt.Printf("Day %s not implemented.\n", day)
	}
}
