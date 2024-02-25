package day16

import (
	"testing"
)

func TestP1(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/day16.txt", 46},
	}

	for _, tc := range testCases {
		initseq := readInput(tc.filename)
		energized := engergizedTiles(initseq)
		if energized != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, energized)
		}
	}
}
