package day12

import (
	"testing"
)

func TestP1(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/day12.txt", 21},
	}

	for _, tc := range testCases {
		springRecords := readSpringRecords(tc.filename)
		totalMatches := 0
		for _, cr := range springRecords {
			matches := cr.numArrangements(1)
			totalMatches += matches
		}
		if totalMatches != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, totalMatches)
		}
	}
}

func TestP2(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/day12.txt", 525152},
	}

	for _, tc := range testCases {
		springRecords := readSpringRecords(tc.filename)
		totalMatches := 0
		for _, cr := range springRecords {
			matches := cr.numArrangements(5)
			totalMatches += matches
		}
		if totalMatches != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, totalMatches)
		}
	}
}
