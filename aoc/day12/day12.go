package day12

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ConditionalRecord struct {
	Springs string
	Pattern []int
}

// Increase length of Springs by repeating it "times" with "sep" in between
func (cr *ConditionalRecord) expandString(times int, sep string) string {
	if times == 1 {
		return cr.Springs
	}
	return strings.TrimSuffix(strings.Repeat(cr.Springs+sep, times), sep)
}

// Increase length of Pattern by repeating "times"
func (cr *ConditionalRecord) expandPattern(times int) []int {
	var newPattern []int
	for i := 0; i < times; i++ {
		newPattern = append(newPattern, cr.Pattern...)
	}
	return newPattern
}

// https://adventofcode.com/2023/day/12
// Part 1 fold == 1
// Part 2 fold == 5
func (cr *ConditionalRecord) numArrangements(fold int) int {
	pattern := cr.expandPattern(fold)
	record := cr.expandString(fold, "?")

	positions := make(map[int]int)
	positions[0] = 1
	for i, contiguous := range pattern {
		updatedPositions := make(map[int]int)
		for key, val := range positions {
			for n := key; n < len(record)-sum(pattern, i+1, len(pattern))+len(pattern[i+1:]); n++ {
				if n+contiguous-1 < len(record) && !isIn(".", record[n:n+contiguous]) {
					if (i == len(pattern)-1 && !isIn("#", record[n+contiguous:])) ||
						(i < len(pattern)-1 && n+contiguous < len(record) && string(record[n+contiguous]) != "#") {
						posvalue, exists := updatedPositions[n+contiguous+1]
						if exists {
							updatedPositions[n+contiguous+1] = posvalue + val
						} else {
							updatedPositions[n+contiguous+1] = val
						}
					}
				}
				if string(record[n]) == "#" {
					break
				}
			}
		}
		positions = updatedPositions
	}

	// sum valid arrangements
	validArrangments := 0
	for _, v := range positions {
		validArrangments += v
	}

	return validArrangments
}

// sum part of a slice
func sum(arr []int, start, stop int) int {
	subsetSum := 0
	for i := start; i < stop; i++ {
		subsetSum += arr[i]
	}
	return subsetSum
}

// is x a substring of arr
func isIn(x string, arr string) bool {
	for _, a := range arr {
		if string(a) == x {
			return true
		}
	}
	return false
}

func inputToSlice(s string) []int {
	strSlice := strings.Split(s, ",")

	var intSlice []int
	for _, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return intSlice
		}
		intSlice = append(intSlice, num)
	}

	return intSlice
}

// Read the pipe map into slice of slice of strings.
func readSpringRecords(filename string) []ConditionalRecord {
	file, _ := os.Open(filename)
	var sequences []ConditionalRecord

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := strings.Split(line, " ")
		sequences = append(sequences, ConditionalRecord{matches[0], inputToSlice(matches[1])})
	}

	return sequences
}

func Run(filename string) {
	springRecords := readSpringRecords(filename)

	totalMatches := 0
	for _, cr := range springRecords {
		matches := cr.numArrangements(1)
		totalMatches += matches
		// fmt.Printf("Matches: %d (%s <- %v)\n", matches, cr.Springs, cr.Pattern)
	}
	fmt.Printf("Total matches (Part 1): %d\n", totalMatches)

	totalMatches = 0
	for _, cr := range springRecords {
		matches := cr.numArrangements(5)
		totalMatches += matches
		// fmt.Printf("Matches: %d (%s <- %v)\n", matches, cr.Springs, cr.Pattern)
	}
	fmt.Printf("Total matches (Part 2): %d\n", totalMatches)
}
