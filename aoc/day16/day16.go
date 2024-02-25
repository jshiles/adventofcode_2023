package day16

import (
	"fmt"
)

// func directionChange (dir int, mirror string) int {
// 	switch mirror {
// 	case "/":
// 		switch dir {
// 			case
// 		}
// 	}
// }

type Space struct {
	X, Y     int
	Contents string
}

// > / = -90
// v / = -90 or 270
// ^ / = +90
// < / = -90

func (s *Space) OutputDirection(incomingDir int) []int {
	var outgoing []int

	if s.Contents == "." {
		outgoing = append(outgoing, incomingDir)
	} else if (s.Contents == "/" && incomingDir == 270) || (s.Contents == "\\" && incomingDir == 180) {
		outgoing = append(outgoing, incomingDir-90)
	} else if s.Contents == "\\" && incomingDir == 270 || (s.Contents == "/" && incomingDir == 90) {
		outgoing = append(outgoing, incomingDir+90)
	}

	return outgoing
}

func engergizedTiles(initseq []string) int {
	return 0
}

func readInput(filename string) []string {
	var seq []string
	return seq
}

func Run(filename string) {
	initseq := readInput(filename)
	energized := engergizedTiles(initseq)
	fmt.Printf("Energized tiles: %d\n", energized)
}
