package day11

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Point struct {
	X int
	Y int
}

func (p1 *Point) manhTrunc(p2 Point, truncFactor int, truncRows, truncCols []int) int {
	distance := 0

	minX, maxX := min(p1.X, p2.X), max(p1.X, p2.X)
	for i := minX; i < maxX; i++ {
		if isIn(i, truncRows) {
			distance += truncFactor
		} else {
			distance += 1
		}
	}

	minY, maxY := min(p1.Y, p2.Y), max(p1.Y, p2.Y)
	for i := minY; i < maxY; i++ {
		if isIn(i, truncCols) {
			distance += truncFactor
		} else {
			distance += 1
		}
	}

	return distance
}

func isIn(x int, arr []int) bool {
	for _, a := range arr {
		if a == x {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func emptyCols(arr [][]string) []int {
	var emptyIdx []int
	for cIdx, _ := range arr[0] {
		empty := true
		for rIdx, _ := range arr {
			if arr[rIdx][cIdx] != "." {
				empty = false
			}
		}
		if empty {
			emptyIdx = append(emptyIdx, cIdx)
		}
	}
	return emptyIdx
}

func emptyRows(arr [][]string) []int {
	var emptyIdx []int
	for rIdx, row := range arr {
		empty := true
		for cIdx, _ := range row {
			if arr[rIdx][cIdx] != "." {
				empty = false
			}
		}
		if empty {
			emptyIdx = append(emptyIdx, rIdx)
		}
	}
	return emptyIdx
}

func galaxyFinder(universe [][]string) []Point {
	var galaxies []Point

	for i, seq := range universe {
		for j, space := range seq {
			if space == "#" {
				galaxies = append(galaxies, Point{i, j})
			}
		}
	}

	return galaxies
}

func uniquePairs(points []Point) [][]Point {
	var pairs [][]Point

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pair := []Point{points[i], points[j]}
			pairs = append(pairs, pair)
		}
	}

	return pairs
}

// Read the pipe map into slice of slice of strings.
func readUniverse(filename string) [][]string {
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
	universe := readUniverse(filename)

	galaxies := galaxyFinder(universe)
	fmt.Printf("Galaxies (%d): ", len(galaxies))
	for _, p := range galaxies {
		fmt.Printf("(%d, %d) ", p.X, p.Y)
	}
	fmt.Printf("\n")

	sumDistancesGravity := 0
	pairs := uniquePairs(galaxies)
	eCols := emptyCols(universe)
	eRows := emptyRows(universe)
	for _, pair := range pairs {
		p1, p2 := pair[0], pair[1]
		minDistance := p1.manhTrunc(p2, 2, eRows, eCols)
		sumDistancesGravity += minDistance
	}
	fmt.Printf("P1: Sum min distances (2x): %d\n", sumDistancesGravity)

	sumDistancesGravity = 0
	for _, pair := range pairs {
		p1, p2 := pair[0], pair[1]
		minDistance := p1.manhTrunc(p2, 1000000, eRows, eCols)
		sumDistancesGravity += minDistance
	}
	fmt.Printf("P2: Sum min distances(1Mx): %d\n", sumDistancesGravity)
}
