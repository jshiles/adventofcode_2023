package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

//
// Part 1 - Human Movement
//

func tailRecSearch(node string, instructions []string, instr_idx int, network map[string]map[string]string) int {
	// fmt.Printf("Node: %s\n", node)
	if node == "ZZZ" {
		return 0
	}
	if instr_idx >= len(instructions) {
		instr_idx = 0
	}
	next_node := network[node][instructions[instr_idx]]
	return 1 + tailRecSearch(next_node, instructions, instr_idx+1, network)
}

func stepsToZZZ(instructions []string, network map[string]map[string]string) int {
	return tailRecSearch("AAA", instructions, 0, network)
}

//
// Part 2 - Ghost Movement
//

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return int(math.Abs(float64(a)))
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b))) / gcd(a, b)
}

func calculateLCM(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Handle empty array case
	}
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

func endsWithZ(node string) bool {
	if len(node) > 0 && node[len(node)-1] != 'Z' {
		return false
	}
	return true
}

func extractStringsEndingWithA(inputStrings []string) []string {
	var result []string
	for _, str := range inputStrings {
		if len(str) > 0 && str[len(str)-1] == 'A' {
			result = append(result, str)
		}
	}
	return result
}

func extractKeysIntoArray(netowrk map[string]map[string]string) []string {
	var nodes []string
	for key := range netowrk {
		nodes = append(nodes, key)
	}
	return nodes
}

func searchPathGhost(node string, instructions []string, instr_idx int, network map[string]map[string]string) int {
	nextNode := node
	steps := 1
	for {
		if endsWithZ(nextNode) {
			break
		}
		if instr_idx >= len(instructions) {
			instr_idx = 0
		}
		nextNode = network[nextNode][instructions[instr_idx]]
		steps += 1
		instr_idx += 1
	}
	return steps
}

// Each path to Z is a cylce from the start, even in the test example. :/
// Brute forcing it would run for quite some time. But, it should be the
// LCM of the cycle lengths (path -1).
func stepsToZZZGhostStyle(instructions []string, network map[string]map[string]string) int {
	currNodes := extractStringsEndingWithA(extractKeysIntoArray(network))
	pathSteps := make([]int, len(currNodes))
	for idx, node := range currNodes {
		steps := searchPathGhost(node, instructions, 0, network)
		pathSteps[idx] = steps - 1
		fmt.Printf("%s -> %d\n", node, steps)
	}
	return calculateLCM(pathSteps)
}

func read_file(filename string) ([]string, map[string]map[string]string) {
	file, _ := os.Open(filename)
	var instructions []string
	objectMap := make(map[string]map[string]string)

	lineCnt := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if lineCnt == 0 {
			for _, b := range []byte(line) {
				instructions = append(instructions, string(b))
			}
		} else {
			re := regexp.MustCompile(`\b[0-9A-Z]{3}\b`)
			matches := re.FindAllString(line, -1)
			if len(matches) == 3 {
				objectMap[matches[0]] = map[string]string{
					"L": matches[1],
					"R": matches[2],
				}
			}
		}
		lineCnt += 1
	}

	return instructions, objectMap
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run day_08.go <filename>")
		return
	}
	filename := os.Args[1]
	instructions, objectMap := read_file(filename)
	steps := stepsToZZZGhostStyle(instructions, objectMap)
	fmt.Printf("Steps: %d\n", steps)
}
