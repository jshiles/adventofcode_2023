package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func transpose(data [][]string) [][]string {
	if len(data) == 0 || len(data[0]) == 0 {
		return data
	}

	rows, cols := len(data), len(data[0])

	result := make([][]string, cols)
	for i := range result {
		result[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = data[i][j]
		}
	}

	return result
}

func twoRowsIdentical(data [][]string, r1, r2 int) bool {
	if r1 >= len(data) || r1 < 0 || r2 >= len(data) || r2 < 0 {
		return false
	}
	for colIndex := 0; colIndex < len(data[0]); colIndex++ {
		if data[r1][colIndex] != data[r2][colIndex] {
			return false
		}
	}
	return true
}

func isMirror(data [][]string, p1, p2 int) bool {
	if p1 < 0 || p1 >= len(data) || p2 < 0 || p2 >= len(data) {
		return true
	}
	if twoRowsIdentical(data, p1, p2) {
		return isMirror(data, p1-1, p2+1)
	}
	return false
}

func findMirrors(data [][]string) int {
	for rowIndex := 0; rowIndex < len(data)-1; rowIndex++ {
		if isMirror(data, rowIndex, rowIndex+1) {
			// fmt.Printf("Horizontal Mirror found: %d, %d\n", rowIndex, rowIndex+1)
			return rowIndex + 1
		}
	}

	return 0
}

func read_file(filename string) [][][]string {
	file, _ := os.Open(filename)
	var island [][][]string
	var pattern [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			island = append(island, pattern)
			pattern = nil
		} else {
			re := regexp.MustCompile(`.{1}`)
			matches := re.FindAllString(line, -1)
			pattern = append(pattern, matches)
		}
	}
	if pattern != nil {
		island = append(island, pattern)
	}

	return island
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run day_13.go <filename>")
		return
	}
	filename := os.Args[1]
	lavaIsland := read_file(filename)

	fmt.Println(len(lavaIsland))

	summary := 0
	for _, pattern := range lavaIsland {
		for _, seq := range pattern {
			fmt.Printf("%v\n", seq)
		}
		horizontal := 100 * findMirrors(pattern)
		vertical := findMirrors(transpose(pattern))
		fmt.Printf("Pattern summary: %d (%d, %d)\n", horizontal+vertical, horizontal, vertical)
		fmt.Println("")
		summary += horizontal + vertical
	}
	fmt.Printf("Part 1: %d \n", summary)
}
