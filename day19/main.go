package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkOnsenTowelsDesigns(design string, patterns map[string]bool, memo map[string][]string, maxDesign int) ([]string, bool) {
	if design == "" {
		return []string{}, true // Empty design is trivially possible
	}

	if val, found := memo[design]; found {
		return val, len(val) > 0 // Use cached result
	}

	finalDesign := map[string][][]string{}

	found := false
	for pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			// Check if the rest of the design can be constructed
			remainder := design[len(pattern):]
			if sequence, ok := checkOnsenTowelsDesigns(remainder, patterns, memo, maxDesign); ok {
				found = true
				memo[design] = append([]string{pattern}, sequence...)
				// fmt.Println("len ele and design ", len(design), maxDesign, memo[design])
				if len(design) == maxDesign {
					finalDesign[design] = append(finalDesign[design], memo[design])
					// fmt.Println(finalDesign)
				}
			}
		}
	}

	return memo[design], found
}

func countWays(design string, patterns map[string]bool, memo map[string]int) int {
	if design == "" {
		return 1 // Base case: empty design can be made in one way (do nothing)
	}

	if val, found := memo[design]; found {
		return val // Use cached result
	}

	totalWays := 0
	for pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			remaining := design[len(pattern):]
			totalWays += countWays(remaining, patterns, memo)
		}
	}

	memo[design] = totalWays // Cache the result
	return totalWays
}

func main() {
	// Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Step 1: Read the first line into a map
	var elementMap map[string]bool
	maxPatternLength := 0
	if scanner.Scan() {
		elementMap = make(map[string]bool)
		firstLine := scanner.Text()
		elements := strings.Split(firstLine, ",")
		for _, element := range elements {
			element = strings.TrimSpace(element)
			if len(element) > maxPatternLength {
				maxPatternLength = len(element)
			}
			elementMap[element] = true
		}
	}
	scanner.Scan()

	// Step 2: Read the remaining lines into a [][]string
	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Output the results
	fmt.Println("Element Map:", maxPatternLength)
	for key := range elementMap {
		fmt.Printf("%s: %v\n", key, elementMap[key])
	}

	fmt.Println("\nGrid:")
	for _, row := range grid {
		fmt.Println(row)
	}

	// PART 1
	// patternsUsed := make(map[string][]string)
	// countSuccess := 0
	// for _, ele := range grid {
	// 	if _, found := checkOnsenTowelsDesigns(ele, elementMap, patternsUsed, len(ele)); found {
	// 		countSuccess++
	// 	}
	// }
	// fmt.Println("patternsUsed ", patternsUsed)

	// PART 2
	patternsUsedPart2 := make(map[string]int)
	countSuccess := 0
	for _, ele := range grid {
		count := countWays(ele, elementMap, patternsUsedPart2)
		countSuccess += count
	}
	fmt.Println("successCount ", countSuccess)
}
