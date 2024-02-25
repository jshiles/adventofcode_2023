package day08

import (
	"fmt"
	"os"
	"testing"
)

func TestGCD(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{8, 12, 4},   // GCD of 8 and 12 is 4
		{17, 23, 1},  // GCD of 17 and 23 is 1
		{30, 42, 6},  // GCD of 30 and 42 is 6
		{0, 5, 5},    // GCD of 0 and 5 is 5
		{0, 0, 0},    // GCD of 0 and 0 is 0
		{-10, 5, 5},  // GCD of -10 and 5 is 5
		{-8, -12, 4}, // GCD of -8 and -12 is 4
		{8, -12, 4},  // GCD of 8 and -12 is 4
	}

	for _, tc := range testCases {
		actual := gcd(tc.a, tc.b)
		if actual != tc.expected {
			t.Errorf("GCD(%d, %d): expected %d, got %d", tc.a, tc.b, tc.expected, actual)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/day08.txt", 6},
	}

	fmt.Println(os.Getwd())
	for _, tc := range testCases {
		instructions, objectMap := readInstructionsMap(tc.filename)
		actual := stepsToZZZGhostStyle(instructions, objectMap)
		if actual != tc.expected {
			t.Errorf("stepsToZZZGhostStyle: expected %d, got %d", tc.expected, actual)
		}
	}
}
