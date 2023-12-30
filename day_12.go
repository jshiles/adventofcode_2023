package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ConditionalRecord struct {
	Springs string
	Pattern []int
}

func (cr *ConditionalRecord) numArrangements() int {
	regex := generateRegex(cr.Pattern)
	// fmt.Println(regex)

	var permutations []string
	generatePermutations(cr.Springs, 0, "", &permutations)

	validArrangments := 0
	for _, perm := range permutations {
		match, err := regexp.MatchString(regex, perm)
		if err == nil {
			// fmt.Printf("%s %v\n", perm, match)
			if match {
				validArrangments += 1
			}
		}
	}

	return validArrangments
}

func generateRegex(pattern []int) string {
	var regexPattern strings.Builder
	regexPattern.WriteString("^\\.*")
	for i, num := range pattern {
		if i == len(pattern)-1 {
			regexPattern.WriteString(fmt.Sprintf("#{%d}", num))
		} else {
			regexPattern.WriteString(fmt.Sprintf("#{%d}\\.+", num))
		}
	}
	regexPattern.WriteString("\\.*$")
	return regexPattern.String()
}

func generatePermutations(s string, index int, current string, result *[]string) {
	if index == len(s) {
		*result = append(*result, current)
		return
	}

	if s[index] == '?' {
		generatePermutations(s, index+1, current+"#", result)
		generatePermutations(s, index+1, current+".", result)
	} else {
		generatePermutations(s, index+1, current+string(s[index]), result)
	}
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
func read_file(filename string) []ConditionalRecord {
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run day_12.go <filename>")
		return
	}
	filename := os.Args[1]
	springRecords := read_file(filename)
	totalMatches := 0
	for _, cr := range springRecords {
		matches := cr.numArrangements()
		totalMatches += matches
		// fmt.Printf("Matches: %d (%s <- %v)\n", matches, cr.Springs, cr.Pattern)
	}
	fmt.Printf("Total matches: %d\n", totalMatches)
}
