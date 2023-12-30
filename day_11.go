package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

type Point struct {
	X int
	Y int
}

func (p1 *Point) manh(p2 Point) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}

func insertDuplicateColumn(arr [][]string, columnIndex int) [][]string {
	for i := range arr {
		if columnIndex >= 0 && columnIndex < len(arr[i]) {
			// Duplicate the column
			duplicate := make([]string, len(arr[i][columnIndex:]))
			copy(duplicate, arr[i][columnIndex:])

			// Insert the duplicated column next to the original column
			arr[i] = append(arr[i][:columnIndex+1], append([]string{"."}, arr[i][columnIndex+1:]...)...)
			copy(arr[i][columnIndex+1:], duplicate)
		}
	}

	return arr
}

func insertDuplicateRow(arr [][]string, rowIndex int) [][]string {
	var dupArr [][]string

	for i, _ := range arr {
		duplicate := make([]string, len(arr[rowIndex]))
		copy(duplicate, arr[i])
		dupArr = append(dupArr, duplicate)
		if i == rowIndex {
			duplicate2 := make([]string, len(arr[rowIndex]))
			copy(duplicate2, arr[rowIndex])
			dupArr = append(dupArr, duplicate2)
		}
	}

	return dupArr
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

func gravitationalExpansion(universe [][]string) [][]string {
	eCols := emptyCols(universe)
	eRows := emptyRows(universe)
	for i, col := range eCols {
		universe = insertDuplicateColumn(universe, col+i)
	}
	for i, row := range eRows {
		universe = insertDuplicateRow(universe, row+i)
	}
	return universe
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
func read_file(filename string) [][]string {
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run day_11.go <filename>")
		return
	}
	filename := os.Args[1]
	universe := read_file(filename)
	universe = gravitationalExpansion(universe)
	// for _, seq := range universe {
	// 	fmt.Printf("%v\n", seq)
	// }

	galaxies := galaxyFinder(universe)
	fmt.Printf("Galaxies (%d): ", len(galaxies))
	for _, p := range galaxies {
		fmt.Printf("(%d, %d) ", p.X, p.Y)
	}
	fmt.Printf("\n")

	sumDistances := 0
	pairs := uniquePairs(galaxies)
	for _, pair := range pairs {
		p1, p2 := pair[0], pair[1]
		minDistance := p1.manh(p2)
		sumDistances += minDistance
		// fmt.Printf("Pair (%d, %d), (%d, %d) -> %d\n", p1.X, p1.Y, p2.X, p2.Y, minDistance)
	}
	fmt.Printf("Sum min distances: %d\n", sumDistances)
}
