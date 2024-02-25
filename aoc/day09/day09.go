package day09

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Check if slice of ints are all 0
func allZeros(seq []int) bool {
	for _, x := range seq {
		if x != 0 {
			return false
		}
	}
	return true
}

// Logic described in part 1: https://adventofcode.com/2023/day/9
func extrapolateRight(seq []int) int {
	if allZeros(seq) || len(seq) <= 1 {
		return 0
	}
	var nextSeq []int
	for i := 1; i < len(seq); i++ {
		nextSeq = append(nextSeq, seq[i]-seq[i-1])
	}
	return extrapolateRight(nextSeq) + seq[len(seq)-1]
}

// Logic described in part 2: https://adventofcode.com/2023/day/9
func extrapolateLeft(seq []int) int {
	if allZeros(seq) || len(seq) <= 1 {
		return 0
	}
	var nextSeq []int
	for i := 1; i < len(seq); i++ {
		nextSeq = append(nextSeq, seq[i]-seq[i-1])
	}
	return seq[0] - extrapolateLeft(nextSeq)
}

// Read file of integeters into slice of slice of ints.
func readIntegers(filename string) [][]int {
	file, _ := os.Open(filename)
	var sequences [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var seq []int
		line := scanner.Text()
		re := regexp.MustCompile(`-{0,1}\d+`)
		matches := re.FindAllString(line, -1)
		for _, i := range matches {
			num, _ := strconv.Atoi(i)
			seq = append(seq, num)
		}
		sequences = append(sequences, seq)
	}

	return sequences
}

func Run(filename string) {
	sum := 0
	sequences := readIntegers(filename)
	for _, sequence := range sequences {
		val := extrapolateLeft(sequence)
		sum += val
		// fmt.Printf("Next in sequence(%v) -> %d\n", sequence, val)
	}
	fmt.Printf("Sum of extrapolated values: %d\n", sum)
}
