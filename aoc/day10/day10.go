package day10

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
)

type Point struct {
	row int
	col int
}

func (p *Point) equals(other Point) bool {
	return p.row == other.row && p.col == other.col
}

func (p *Point) adjacent(other Point) bool {
	return (int(math.Abs(float64(p.row-other.row))) == 1 && p.col == other.col) ||
		(int(math.Abs(float64(p.col-other.col))) == 1 && p.row == other.row)
}

func (p *Point) isIn(arr []Point) bool {
	for _, other := range arr {
		if p.row == other.row && p.col == other.col {
			return true
		}
	}
	return false
}

func findIndex(matrix [][]string, target string) (Point, error) {
	for i, row := range matrix {
		for j, element := range row {
			if element == target {
				return Point{row: i, col: j}, nil
			}
		}
	}
	return Point{row: -1, col: -1}, errors.New("S not found in matrix")
}

func valueAt(m [][]string, p Point) (string, error) {
	if validPoint(m, p) {
		return m[p.row][p.col], nil
	}
	return "", nil
}

func validPoint(m [][]string, p Point) bool {
	if p.row < 0 || p.row >= len(m) || p.col < 0 || p.col >= len(m[0]) {
		return false
	}
	return true
}

func isStart(m [][]string, p Point) bool {
	startPoint, _ := findIndex(m, "S")
	return p.equals(startPoint)
}

func canMoveEast(m [][]string, p Point) bool {
	currVal, _ := valueAt(m, p)
	east := Point{p.row, p.col + 1}
	eastVal, err := valueAt(m, east)
	if err == nil &&
		(currVal == "S" || currVal == "-" || currVal == "F" || currVal == "L") &&
		(eastVal == "S" || eastVal == "-" || eastVal == "7" || eastVal == "J") {
		return true
	}
	return false
}

func canMoveWest(m [][]string, p Point) bool {
	currVal, _ := valueAt(m, p)
	west := Point{p.row, p.col - 1}
	westVal, err := valueAt(m, west)
	if err == nil &&
		(currVal == "S" || currVal == "-" || currVal == "7" || currVal == "J") &&
		(westVal == "S" || westVal == "-" || westVal == "F" || westVal == "L") {
		return true
	}
	return false
}

func canMoveNorth(m [][]string, p Point) bool {
	currVal, _ := valueAt(m, p)
	north := Point{p.row - 1, p.col}
	northVal, err := valueAt(m, north)
	if err == nil &&
		(currVal == "S" || currVal == "|" || currVal == "L" || currVal == "J") &&
		(northVal == "S" || northVal == "|" || northVal == "F" || northVal == "7") {
		return true
	}
	return false
}

func canMoveSouth(m [][]string, p Point) bool {
	currVal, _ := valueAt(m, p)
	south := Point{p.row + 1, p.col}
	southVal, err := valueAt(m, south)
	if err == nil &&
		(currVal == "S" || currVal == "|" || currVal == "F" || currVal == "7") &&
		(southVal == "S" || southVal == "|" || southVal == "L" || southVal == "J") {
		return true
	}
	return false
}

func camMoveThere(m [][]string, start Point, end Point) bool {
	if start.adjacent(end) {
		if start.row > end.row {
			return canMoveNorth(m, start)
		} else if start.row < end.row {
			return canMoveSouth(m, start)
		} else if start.col > end.col {
			return canMoveWest(m, start)
		} else if start.col < end.col {
			return canMoveEast(m, start)
		}
	}
	return false
}

func traverse(m [][]string, p Point, trail []Point) ([]Point, error) {
	// base case
	endPoint, _ := findIndex(m, "S")
	if camMoveThere(m, p, endPoint) && len(trail) > 2 {
		trail = append(trail, p)
		fmt.Printf("Trail (%d): ", len(trail))
		for _, p := range trail {
			fmt.Printf("(%d, %d) ", p.row, p.col)
		}
		fmt.Printf("\n")
		return trail, nil
	}

	north := Point{p.row - 1, p.col}
	if canMoveNorth(m, p) && !north.isIn(trail) {
		// fmt.Printf("Going north from (%d, %d) to (%d, %d)\n", p.row, p.col, north.row, north.col)
		retTrail, error := traverse(m, north, append(trail, p))
		if error == nil {
			return retTrail, nil
		}
	}

	south := Point{p.row + 1, p.col}
	if canMoveSouth(m, p) && !south.isIn(trail) {
		// fmt.Printf("Going south from (%d, %d) to (%d, %d)\n", p.row, p.col, south.row, south.col)
		retTrail, error := traverse(m, south, append(trail, p))
		if error == nil {
			return retTrail, nil
		}
	}

	west := Point{p.row, p.col - 1}
	if canMoveWest(m, p) && !west.isIn(trail) {
		// fmt.Printf("Going west from (%d, %d) to (%d, %d)\n", p.row, p.col, west.row, west.col)
		retTrail, error := traverse(m, west, append(trail, p))
		if error == nil {
			return retTrail, nil
		}
	}

	east := Point{p.row, p.col + 1}
	if canMoveEast(m, p) && !east.isIn(trail) {
		// fmt.Printf("Going east from (%d, %d) to (%d, %d)\n", p.row, p.col, east.row, east.col)
		retTrail, error := traverse(m, east, append(trail, p))
		if error == nil {
			return retTrail, nil
		}
	}

	return []Point{}, nil
}

func findPath(m [][]string) []Point {
	start, _ := findIndex(m, "S")
	fmt.Printf("S @ (%d,%d)\n", start.row, start.col)

	path, _ := traverse(m, start, []Point{})
	return path
}

// Ray casting algorithm
func isPointPartOfPolygon(point Point, polygon []Point) bool {
	intersections := 0
	n := len(polygon)

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		if (polygon[i].col > point.col) != (polygon[j].col > point.col) &&
			(point.row < (polygon[j].row-polygon[i].row)*(point.col-polygon[i].col)/(polygon[j].col-polygon[i].col)+polygon[i].row) {
			intersections++
		}
	}

	return intersections%2 != 0
}

// Read the pipe map into slice of slice of strings.
func readFile(filename string) [][]string {
	file, _ := os.Open(filename)
	var sequences [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`.{1}`)
		matches := re.FindAllString(line, -1)
		sequences = append(sequences, matches)
	}

	return sequences
}

func Run(filename string) {
	pipeMap := readFile(filename)

	// Part 1
	path := findPath(pipeMap)
	stepsInTrail := len(path)
	fmt.Printf("Steps to farthest point: %d\n", int(stepsInTrail/2))

	// Part 2
	pointsInside := 0
	for i, seq := range pipeMap {
		for j, _ := range seq {
			p := Point{i, j}
			if isPointPartOfPolygon(p, path[:len(path)-1]) && !p.isIn(path) {
				// fmt.Printf("Inside (%d, %d)\n", i, j)
				pointsInside += 1
			}
		}
	}
	fmt.Printf("Points inside the pipe path: %d\n", pointsInside)
}
