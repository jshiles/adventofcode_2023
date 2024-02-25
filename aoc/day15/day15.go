package day15

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Label       string
	FocalLength int
}

type Box struct {
	LensSlots []Lens
}

func (b *Box) AddLens(label string, focalLength int) {
	var amendedLenses []Lens
	updated := false
	for _, lens := range b.LensSlots {
		if lens.Label != label {
			amendedLenses = append(amendedLenses, lens)
		} else {
			amendedLenses = append(amendedLenses, Lens{label, focalLength})
			updated = true
		}
	}
	if !updated {
		amendedLenses = append(amendedLenses, Lens{label, focalLength})
	}
	b.LensSlots = amendedLenses
}

func (b *Box) RemoveLens(labelToRemove string) {
	var filteredLenses []Lens
	for _, lens := range b.LensSlots {
		if lens.Label != labelToRemove {
			filteredLenses = append(filteredLenses, lens)
		}
	}
	b.LensSlots = filteredLenses
}

func focusingPower(boxes [256]Box) int {
	power := 0
	for boxnum, box := range boxes {
		for slot, lens := range box.LensSlots {
			focalPower := (1 + boxnum) * (1 + slot) * lens.FocalLength
			// fmt.Printf("%s -> %d\n", lens.Label, focalPower)
			power += focalPower
		}
	}
	return power
}

func evaluateHash(s string) int {
	currentVal := 0
	for _, char := range s {
		currentVal += int(char)
		currentVal *= 17
		currentVal = currentVal % 256
		// fmt.Printf("%c: %d -> %d\n", char, char, currentVal)
	}
	return currentVal
}

func extractLabel(input string) (string, int, string, int) {
	if strings.Contains(input, "-") || strings.Contains(input, "=") {
		parts := strings.FieldsFunc(input, func(r rune) bool {
			return r == '-' || r == '='
		})

		label := parts[0]
		boxID := evaluateHash(label)
		action := string(input[len(parts[0])])
		focalLength := 0
		if len(parts) > 1 {
			focalLength, _ = strconv.Atoi(parts[1])
		}

		return label, boxID, action, focalLength
	}

	return input, 0, "", 0
}

func readInput(filename string) []string {
	var data []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return data
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var originalString string
	for scanner.Scan() {
		originalString += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return data
	}

	stringWithoutNewlines := strings.ReplaceAll(originalString, "\n", "")
	parts := strings.Split(stringWithoutNewlines, ",")
	data = append(data, parts...)

	return data
}

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
