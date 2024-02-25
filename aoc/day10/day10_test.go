package day10

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/day10a.txt", 4},
		{"testdata/day10b.txt", 8},
		{"testdata/day10c.txt", 10},
	}

	for _, tc := range testCases {
		pipeMap := readFile(tc.filename)
		path := findPath(pipeMap)
		pointsInside := 0
		for i, seq := range pipeMap {
			for j, _ := range seq {
				p := Point{i, j}
				if isPointPartOfPolygon(p, path[:len(path)-1]) && !p.isIn(path) {
					pointsInside += 1
				}
			}
		}
		if pointsInside != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, pointsInside)
		}
	}
}
