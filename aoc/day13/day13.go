package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func transpose(data [][]int) [][]int {
	if len(data) == 0 || len(data[0]) == 0 {
		return data
	}

	rows, cols := len(data), len(data[0])

	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = data[i][j]
		}
	}

	return result
}

func subtractRow(data [][]int, r1, r2 int) int {
	if r1 >= len(data) || r1 < 0 || r2 >= len(data) || r2 < 0 {
		return 0
	}
	diff := 0
	for colIndex := 0; colIndex < len(data[0]); colIndex++ {
		diff += int(math.Abs(float64(data[r1][colIndex] - data[r2][colIndex])))
	}
	return diff
}

func twoRowsIdentical(data [][]int, r1, r2 int) bool {
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

func isMirror(data [][]int, p1, p2, smudges int) bool {
	if (p1 < 0 || p1 >= len(data) || p2 < 0 || p2 >= len(data)) && smudges == 0 {
		return true
	}
	if smudges > 0 && subtractRow(data, p1, p2) == 1 {
		return isMirror(data, p1-1, p2+1, smudges-1)
	} else if twoRowsIdentical(data, p1, p2) {
		return isMirror(data, p1-1, p2+1, smudges)
	}
	return false
}

func findMirrors(data [][]int, smudges int) int {
	for rowIndex := 0; rowIndex < len(data)-1; rowIndex++ {
		if isMirror(data, rowIndex, rowIndex+1, smudges) {
			// fmt.Printf("Horizontal Mirror found: %d, %d\n", rowIndex, rowIndex+1)
			return rowIndex + 1
		}
	}

	return 0
}

func convertStringSliceToIntSlice(strSlice []string) ([]int, error) {
	intSlice := make([]int, len(strSlice))

	for i, str := range strSlice {
		// Parse string as integer
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		// Append integer to the new slice
		intSlice[i] = num
	}

	return intSlice, nil
}

func readLavaIsland(filename string) [][][]int {
	file, _ := os.Open(filename)
	var island [][][]int
	var pattern [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			island = append(island, pattern)
			pattern = nil
		} else {
			resultString := strings.Replace(line, string("."), string("0"), -1)
			resultString = strings.Replace(resultString, string("#"), string("1"), -1)
			re := regexp.MustCompile(`.{1}`)
			matches := re.FindAllString(resultString, -1)
			matchesInt, err := convertStringSliceToIntSlice(matches)
			if err == nil {
				pattern = append(pattern, matchesInt)
			}
		}
	}
	if pattern != nil {
		island = append(island, pattern)
	}

	return island
}

func Run(filename string) {
	lavaIsland := readLavaIsland(filename)

	fmt.Println(len(lavaIsland))

	summary := 0
	for _, pattern := range lavaIsland {
		// for _, seq := range pattern {
		// 	fmt.Printf("%v\n", seq)
		// }
		horizontal := 100 * findMirrors(pattern, 0)
		vertical := findMirrors(transpose(pattern), 0)
		// fmt.Printf("Pattern summary: %d (%d, %d)\n", horizontal+vertical, horizontal, vertical)
		summary += horizontal + vertical
	}
	fmt.Printf("Part 1: %d \n", summary)

	summary = 0
	for _, pattern := range lavaIsland {
		// for _, seq := range pattern {
		// 	fmt.Printf("%v\n", seq)
		// }
		horizontal := 100 * findMirrors(pattern, 1)
		vertical := findMirrors(transpose(pattern), 1)
		// fmt.Printf("Pattern summary: %d (%d, %d)\n", horizontal+vertical, horizontal, vertical)
		summary += horizontal + vertical
	}
	fmt.Printf("Part 2: %d \n", summary)
}
