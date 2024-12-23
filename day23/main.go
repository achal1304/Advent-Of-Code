package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// ReadConnections reads the input file and constructs the connections map
func ReadConnections(filename string) (map[string]map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	connections := make(map[string]map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		computer1, computer2 := parts[0], parts[1]

		if connections[computer1] == nil {
			connections[computer1] = make(map[string]bool)
		}
		if connections[computer2] == nil {
			connections[computer2] = make(map[string]bool)
		}
		connections[computer1][computer2] = true
		connections[computer2][computer1] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return connections, nil
}

// FindThreeInterconnected finds all sets of three interconnected computers
func FindThreeInterconnected(connections map[string]map[string]bool) [][]string {
	resultSets := [][]string{}
	computers := []string{}

	for computer := range connections {
		computers = append(computers, computer)
	}

	for i := 0; i < len(computers); i++ {
		for j := i + 1; j < len(computers); j++ {
			for k := j + 1; k < len(computers); k++ {
				if isConnected(computers[i], computers[j], computers[k], connections) {
					resultSets = append(resultSets, []string{computers[i], computers[j], computers[k]})
				}
			}
		}
	}
	return resultSets
}

// CountSetsWithT counts sets containing at least one computer starting with 't' and prints them
func CountSetsWithT(sets [][]string) int {
	count := 0
	for _, set := range sets {
		if containsT(set) {
			fmt.Printf("Set with 't': %v\n", set)
			count++
		}
	}
	return count
}

// Helper function to check if all elements in a slice are connected
func isConnected(a, b, c string, connections map[string]map[string]bool) bool {
	return connections[a][b] && connections[a][c] && connections[b][c]
}

// Helper function to check if a slice contains an element starting with 't'
func containsT(computers []string) bool {
	for _, computer := range computers {
		if strings.HasPrefix(computer, "t") {
			return true
		}
	}
	return false
}

func SortConnnectionsWithLength(connections map[string]map[string]bool) []string {
	keys := []string{}
	for k, _ := range connections {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return len(connections[keys[i]]) > len(connections[keys[j]])
	})
	return keys
}

func FindMaxInterconnected(connections map[string]map[string]bool, sortedKeys []string) []string {
	maxPair := []string{}
	for _, key := range sortedKeys {
		currentPair := []string{key}
		neighbours := connections[key]

		// as we are going from sortdkeys and for the next keys the length of map
		// is similar to max no. of pairs it can have, we skip all next iterations as we are already at max pairs
		if len(maxPair) >= len(neighbours) {
			return maxPair
		}

	innterLoop:
		for neighbour, _ := range neighbours {
			mapOfNeighbor := connections[neighbour]
			for _, ele := range currentPair {
				if mapOfNeighbor[ele] != true {
					continue innterLoop
				}
			}
			currentPair = append(currentPair, neighbour)
		}
		if len(currentPair) > len(maxPair) {
			maxPair = currentPair
		}
	}
	fmt.Println(maxPair)
	return maxPair
}

func main() {
	connections, err := ReadConnections("input.txt")
	if err != nil {
		fmt.Println("Error reading connections:", err)
		return
	}

	sortedKeys := SortConnnectionsWithLength(connections)
	fmt.Println(sortedKeys)

	// Part1
	// resultSets := FindThreeInterconnected(connections)
	// count := CountSetsWithT(resultSets)
	// fmt.Printf("Total sets of three inter-connected computers containing at least one computer starting with 't': %d\n", count)

	maxPair := FindMaxInterconnected(connections, sortedKeys)
	sort.Strings(maxPair)

	fmt.Println("sorted pairs ", strings.Join(maxPair, ","))
}
