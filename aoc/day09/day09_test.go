package day09

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/day09.txt", 2},
	}

	for _, tc := range testCases {
		sum := 0
		sequences := readIntegers(tc.filename)
		for _, sequence := range sequences {
			val := extrapolateLeft(sequence)
			sum += val
		}
		if sum != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, sum)
		}
	}
}
