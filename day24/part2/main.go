package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to read lines from a text file
func readInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Function to process the value assignments (e.g., x00: 1)
func processValues(lines []string) map[string]int {
	values := make(map[string]int)
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			valueInt, _ := strconv.Atoi(value)
			// Convert value to integer
			values[key] = valueInt
		}
	}
	return values
}

// Function to process the operations (e.g., ntg XOR fgs -> mjb)
func processOperations(lines []string) map[string][]string {
	operations := make(map[string][]string)
	for _, line := range lines {
		if strings.Contains(line, "->") {
			// Split the operation part from the result part
			parts := strings.Split(line, "->")
			result := strings.TrimSpace(parts[1])

			// Now split the left side (e.g., ntg XOR fgs)
			leftParts := strings.Split(parts[0], " ")
			operations[result] = []string{
				strings.TrimSpace(leftParts[0]),
				strings.TrimSpace(leftParts[1]),
				strings.TrimSpace(leftParts[2]),
			}
		}
	}
	return operations
}

func validateGates(operations map[string][]string, values map[string]int) []string {
	faultyGates := []string{}
	zOutputs := map[string]bool{}
	xyAdds := map[string]bool{}
	xyCarries := map[string]bool{}
	usedInXOR := map[string]bool{}
	usedInOR := map[string]bool{}
	usedInAnd := map[string]bool{}

	// Collect data about gates
	for result, op := range operations {
		if strings.HasPrefix(result, "z") {
			zOutputs[result] = true
		}
		if op[1] == "XOR" {
			usedInXOR[result] = true
			if (strings.HasPrefix(op[0], "x") && strings.HasPrefix(op[2], "y")) ||
				(strings.HasPrefix(op[0], "y") && strings.HasPrefix(op[2], "x")) {
				xyAdds[result] = true
			}
		}
		if op[1] == "AND" {
			usedInAnd[result] = true
			if (strings.HasPrefix(op[0], "x") && strings.HasPrefix(op[2], "y")) ||
				(strings.HasPrefix(op[0], "y") && strings.HasPrefix(op[2], "x")) {
				xyCarries[result] = true
			}
		}
		if op[1] == "OR" {
			usedInOR[result] = true
		}
	}

	faultyGatesMap := make(map[string]bool)

	// Validate gates
	for result, op := range operations {
		// Skip gates with inputs exclusively from x and y
		if op[2] == "x00" && op[0] == "y00" {
			continue
		}

		if result == "z45" {
			continue
		}

		// NO z gates should be directly associated to x and y as XOR, AND and OR
		// are used for full adders and if x and y directly associated to z, chain breaks
		if (strings.HasPrefix(op[0], "x") && strings.HasPrefix(op[2], "y")) ||
			(strings.HasPrefix(op[0], "y") && strings.HasPrefix(op[2], "x")) {
			if strings.HasPrefix(result, "z") {
				faultyGatesMap[result] = true
				faultyGates = append(faultyGates, fmt.Sprintf("%s: faulty x y z", result))
				continue
			} else {
				continue
			}
		}

		if strings.HasPrefix(result, "z") {
			// Ouptut operations to Z should always b XOR and one another operation should be the carry or OR oepration
			if op[1] != "XOR" {
				faultyGatesMap[result] = true
				faultyGates = append(faultyGates, fmt.Sprintf("%s: invalid Operation NON XOR to z", result))
				continue
			} else if !(usedInXOR[op[0]] || usedInXOR[op[2]]) {
				if !(xyCarries[op[0]] || xyCarries[op[2]]) {
					faultyGatesMap[result] = true
					faultyGates = append(faultyGates, fmt.Sprintf("%s: invalid INPUT, one CARRY OR/AND expected z", result))
					continue
				}
			}

			// If one of the operation is OR operation which means it is a carry operation,
			// the other operation must be XOR as Z gate is formed by using XOR or previous XOR and a CARRY
			if usedInOR[op[0]] {
				if !usedInXOR[op[2]] {
					faultyGatesMap[op[2]] = true
					faultyGates = append(faultyGates, fmt.Sprintf("%s: invalid XOR input to Z XOR ", op[2]))
				}
			} else if usedInOR[op[2]] {
				if !usedInXOR[op[0]] {
					faultyGatesMap[op[0]] = true
					faultyGates = append(faultyGates, fmt.Sprintf("%s: invalid XOR input to Z XOR ", op[0]))
				}
			}

		} else {
			// In Full adders, an OR operation only takes both AND inputs and no other inputs
			if op[1] == "OR" {
				if !(usedInAnd[op[0]] && (usedInAnd[op[2]])) {
					if !usedInAnd[op[0]] {
						faultyGatesMap[op[0]] = true
						faultyGates = append(faultyGates, fmt.Sprintf("%s: invalid OR output", op[0]))
					} else {
						faultyGatesMap[op[2]] = true
						faultyGates = append(faultyGates, fmt.Sprintf("%s: invalid OR output", op[2]))
					}
				}
				// in NON Z operations, XOR cannot be used as XOR should be specifically used to output the final
				// value to the Z Gate
				// if XOR is on the x and y gate and assigned to non Z gate, then that is fine as it will
				// become part of the chain, but non x and non y input XOR will break the chain
			} else if op[1] == "XOR" {
				faultyGatesMap[result] = true
				faultyGates = append(faultyGates, fmt.Sprintf("%s: faulty XOR not allowed for non z gate", result))
				continue
			}
		}
	}

	fmt.Println("fault gates ", faultyGatesMap)
	return faultyGates
}

func main() {
	// Read the input file
	lines, err := readInputFile("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Process the value assignments
	values := processValues(lines)
	// fmt.Println("Values Map:", values)

	// Process the operations
	operations := processOperations(lines)

	// Validate gates based on the conditions
	faultyGates := validateGates(operations, values)

	// Print faulty gates
	fmt.Println("Faulty Gates description:")
	for _, gate := range faultyGates {
		fmt.Println(gate)
	}
}
