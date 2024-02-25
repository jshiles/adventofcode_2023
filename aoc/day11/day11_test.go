package day11

import (
	"testing"
)

func TestP1(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/day11.txt", 374},
	}

	for _, tc := range testCases {
		universe := readUniverse(tc.filename)
		galaxies := galaxyFinder(universe)
		sumDistancesGravity := 0
		pairs := uniquePairs(galaxies)
		eCols := emptyCols(universe)
		eRows := emptyRows(universe)
		for _, pair := range pairs {
			p1, p2 := pair[0], pair[1]
			minDistance := p1.manhTrunc(p2, 2, eRows, eCols)
			sumDistancesGravity += minDistance
		}
		if sumDistancesGravity != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, sumDistancesGravity)
		}
	}
}

func TestP2(t *testing.T) {
	testCases := []struct {
		filename    string
		timesLarger int
		expected    int
	}{
		{"testdata/day11.txt", 10, 1030},
		{"testdata/day11.txt", 100, 8410},
	}

	for _, tc := range testCases {
		universe := readUniverse(tc.filename)
		galaxies := galaxyFinder(universe)
		sumDistancesGravity := 0
		pairs := uniquePairs(galaxies)
		eCols := emptyCols(universe)
		eRows := emptyRows(universe)
		for _, pair := range pairs {
			p1, p2 := pair[0], pair[1]
			minDistance := p1.manhTrunc(p2, tc.timesLarger, eRows, eCols)
			sumDistancesGravity += minDistance
		}
		if sumDistancesGravity != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, sumDistancesGravity)
		}
	}
}
