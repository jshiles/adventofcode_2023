package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Rock struct {
	Type string
	X, Y int
}

func (r *Rock) MoveNorth(rocks Rocks) bool {
	if r.Type == "O" && r.X > 0 && !rocks.Contains(r.X-1, r.Y) {
		r.X -= 1
		return true
	}
	return false
}

// func (r *Rock) MoveSouth(rocks Rocks, maxX int) bool {
// 	if r.Type == "O" && r.X < maxX && !rocks.Contains(r.X+1, r.Y) {
// 		r.X += 1
// 		return true
// 	}
// 	return false
// }

type Rocks []Rock

func (r Rocks) Contains(x, y int) bool {
	for _, i := range r {
		if i.X == x && i.Y == y {
			return true
		}
	}
	return false
}

func (r Rocks) Len() int { return len(r) }

func (r Rocks) Less(i, j int) bool {
	if r[i].X != r[j].X {
		return r[i].X < r[j].X
	}
	return r[i].Y < r[j].Y
}

func (r Rocks) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func (r Rocks) PrettyPrint(maxX, maxY int) {
	sort.Sort(r)
	rcnt := 0
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if rcnt < len(r) && r[rcnt].X == i && r[rcnt].Y == j {
				fmt.Printf("%s", r[rcnt].Type)
				rcnt++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (r Rocks) LoadNorth(maxX int) int {
	load := 0
	for _, i := range r {
		if i.Type == "O" {
			load += maxX - i.X
		}
	}
	return load
}

func (r Rocks) LoadSouth() int {
	load := 0
	for _, i := range r {
		if i.Type == "O" {
			load += i.X + 1
		}
	}
	return load
}

func read_file(filename string) (Rocks, int, int) {
	file, _ := os.Open(filename)
	var rocks Rocks
	maxX, maxY := 0, 0

	scanner := bufio.NewScanner(file)
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		for col, char := range line {
			if char == 'O' || char == '#' {
				rocks = append(rocks, Rock{string(char), row, col})
			}
		}
		if len(line) > maxY {
			maxY = len(line)
		}
		maxX = row
	}

	return rocks, maxX + 1, maxY
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run day_14.go <filename>")
		return
	}
	filename := os.Args[1]
	rocks, maxX, _ := read_file(filename)

	// fmt.Println(maxX, maxY)
	// rocks.PrettyPrint(maxX, maxY)

	// fmt.Println(rocks[9].MoveNorth(rocks))

	for {
		modified := false
		sort.Sort(rocks)
		for i, _ := range rocks {
			moved := rocks[i].MoveNorth(rocks)
			if moved {
				// fmt.Printf("Moved %s (%d, %d)\n", rocks[i].Type, rocks[i].X, rocks[i].Y)
				modified = true
			}
		}
		if !modified {
			break
		}
	}

	// rocks.PrettyPrint(maxX, maxY)
	fmt.Printf("Southern Load: %d\n", rocks.LoadSouth())
	fmt.Printf("Northern Load: %d\n", rocks.LoadNorth(maxX))
}
