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

// func readInput(filename string) []string {
//
// }

func Run(filename string) {
	initseq := readInput(filename)

	var boxes [256]Box
	for _, seq := range initseq {
		label, boxID, action, focalLength := extractLabel(seq)
		if action == "=" {
			boxes[boxID].AddLens(label, focalLength)
		} else if action == "-" {
			boxes[boxID].RemoveLens(label)
		}
	}

	fmt.Printf("Focal Power: %d\n", focusingPower(boxes))
}
