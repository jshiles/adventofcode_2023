package day14

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

func (r *Rock) MoveWest(rocks Rocks) bool {
	if r.Type == "O" && r.Y > 0 && !rocks.Contains(r.X, r.Y-1) {
		r.Y -= 1
		return true
	}
	return false
}

func (r *Rock) MoveSouth(rocks Rocks, maxX int) bool {
	if r.Type == "O" && r.X < maxX-1 && !rocks.Contains(r.X+1, r.Y) {
		r.X += 1
		return true
	}
	return false
}

func (r *Rock) MoveEast(rocks Rocks, maxY int) bool {
	if r.Type == "O" && r.Y < maxY-1 && !rocks.Contains(r.X, r.Y+1) {
		r.Y += 1
		return true
	}
	return false
}

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

func (rocks Rocks) Cycle(maxX, maxY int) {
	for {
		modified := false
		sort.Sort(rocks)
		for n, _ := range rocks {
			moved := rocks[n].MoveNorth(rocks)
			if moved {
				// fmt.Printf("Moved %s (%d, %d)\n", rocks[i].Type, rocks[i].X, rocks[i].Y)
				modified = true
			}
		}
		if !modified {
			break
		}
	}
	for {
		modified := false
		sort.Sort(rocks)
		for w, _ := range rocks {
			moved := rocks[w].MoveWest(rocks)
			if moved {
				// fmt.Printf("Moved %s (%d, %d)\n", rocks[i].Type, rocks[i].X, rocks[i].Y)
				modified = true
			}
		}
		if !modified {
			break
		}
	}
	for {
		modified := false
		sort.Sort(rocks)
		for s := len(rocks) - 1; s >= 0; s-- {
			moved := rocks[s].MoveSouth(rocks, maxX)
			if moved {
				// fmt.Printf("Moved %s (%d, %d)\n", rocks[s].Type, rocks[s].X, rocks[s].Y)
				modified = true
			}
		}
		if !modified {
			break
		}
	}
	for {
		modified := false
		sort.Sort(rocks)
		for e := len(rocks) - 1; e >= 0; e-- {
			moved := rocks[e].MoveEast(rocks, maxY)
			if moved {
				// fmt.Printf("Moved %s (%d, %d)\n", rocks[i].Type, rocks[i].X, rocks[i].Y)
				modified = true
			}
		}
		if !modified {
			break
		}
	}
}

func areListsIdentical(list1, list2 Rocks) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

func deepCopy(original Rocks) Rocks {
	copyOfArray := make(Rocks, len(original))

	for i, v := range original {
		copyOfArray[i] = Rock{
			Type: v.Type,
			X:    v.X,
			Y:    v.Y,
		}
	}

	return copyOfArray
}

func readRocks(filename string) (Rocks, int, int) {
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

func Run(filename string) {
	rocks, maxX, maxY := readRocks(filename)

	var prev []Rocks
	loops := 1000000000
	cycleDetected := false
	for cycle := 0; cycle < loops; cycle++ {
		prev = append(prev, deepCopy(rocks))
		rocks.Cycle(maxX, maxY)
		fmt.Printf("Cycle [%d]: Northern: %d\n", cycle, rocks.LoadNorth(maxX))

		loopStart := 0
		for i, prior := range prev {
			if areListsIdentical(prior, rocks) {
				cycleDetected = true
				loopStart = i
				break
			}
		}
		if cycleDetected {
			fmt.Printf("At cycle %d, we detcted a cycle because it matched %d\n", cycle, loopStart)
			fmt.Printf("Answer at cycle: %d\n", loopStart+int((loops-loopStart)%(cycle-loopStart+1))-1)
			break
		}
	}
	if !cycleDetected {
		rocks.PrettyPrint(maxX, maxY)
		fmt.Printf("Northern Load: %d\n", rocks.LoadNorth(maxX))
	}
}
