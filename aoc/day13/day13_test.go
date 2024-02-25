package day13

import (
	"testing"
)

func TestP1(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/day13.txt", 405},
	}

	for _, tc := range testCases {
		lavaIsland := readLavaIsland(tc.filename)
		summary := 0
		for _, pattern := range lavaIsland {
			horizontal := 100 * findMirrors(pattern, 0)
			vertical := findMirrors(transpose(pattern), 0)
			summary += horizontal + vertical
		}
		if summary != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, summary)
		}
	}
}

func TestP2(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/day13.txt", 400},
	}

	for _, tc := range testCases {
		lavaIsland := readLavaIsland(tc.filename)
		summary := 0
		for _, pattern := range lavaIsland {
			horizontal := 100 * findMirrors(pattern, 1)
			vertical := findMirrors(transpose(pattern), 1)
			summary += horizontal + vertical
		}
		if summary != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, summary)
		}
	}
}
