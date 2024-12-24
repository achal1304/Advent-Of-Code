package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func checkOperationsGates(operations map[string][]string, values map[string]int) {
	zGates := []string{}
	for k, _ := range operations {
		if _, ok := values[k]; !ok {
			checkEachOperation(k, operations, values)
		}
		if strings.HasPrefix(k, "z") {
			zGates = append(zGates, k)
		}
	}

	sort.Slice(zGates, func(i, j int) bool {
		return zGates[i] > zGates[j]
	})
	result := ""
	// Iterate through the sorted keys and append the corresponding values to the result
	for _, key := range zGates {
		// Convert value to string for appending
		result += fmt.Sprintf("%d", values[key])
	}
	fmt.Println("result is ", result)

	decimalInt, err := strconv.ParseInt(result, 2, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the integer value
	fmt.Println("The decimal value is:", decimalInt)
}

func checkEachOperation(gate string,
	operations map[string][]string,
	values map[string]int) int {
	if v, ok := values[gate]; ok {
		return v
	}

	ele := operations[gate]
	gate1, gate2, op := ele[0], ele[2], ele[1]
	gate1Int, gate2Int := 0, 0
	if _, ok := values[gate1]; !ok {
		gate1Int = checkEachOperation(gate1, operations, values)
	} else {
		gate1Int = values[gate1]
	}

	if _, ok := values[gate2]; !ok {
		gate2Int = checkEachOperation(gate2, operations, values)
	} else {
		gate2Int = values[gate2]
	}

	output := 0
	switch op {
	case "XOR":
		output = gate1Int ^ gate2Int
	case "OR":
		output = gate1Int | gate2Int
	case "AND":
		output = gate1Int & gate2Int
	default:
		fmt.Println("no matching oeprations ")
	}
	values[gate] = output

	return output
}

func main() {
	// Read the input file
	lines, err := readInputFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Process the value assignments
	values := processValues(lines)
	fmt.Println("Values Map:", values)

	// Process the operations
	operations := processOperations(lines)
	// fmt.Println("Operations Map:", operations)

	checkOperationsGates(operations, values)
}
