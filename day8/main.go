package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error reading file ", err)
		return
	}
	defer file.Close()

	antennasMap, lenMap := scanInputAntennas(file)
	fmt.Println("antennasMap ", antennasMap)
	fmt.Println("lenMap ", lenMap)
	fmt.Println(scanTheMap(antennasMap, lenMap))

}

func scanInputAntennas(file *os.File) (map[string][][]int, int) {
	scanner := bufio.NewScanner(file)
	antennasMap := make(map[string][][]int)
	iteration := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, ele := range line {
			if string(ele) != "." {
				newLocation := []int{iteration, i}
				antennasMap[string(ele)] = append(antennasMap[string(ele)], newLocation)
			}
		}
		iteration += 1
	}
	// fmt.Println(antennasMap)
	return antennasMap, iteration
}

func scanTheMap(antennasMap map[string][][]int, lenMap int) int {
	distinctAntiNodes := make(map[string]int)
	totCount := 0
	for _, locations := range antennasMap {
		// PART 1 - calculateAntinodes
		// count := calculateAntinodes(locations, lenMap, distinctAntiNodes)
		// PART 2 - calculateAntinodesResonantHarmonics
		count := calculateAntinodesResonantHarmonics(locations, lenMap, distinctAntiNodes)
		if count > totCount {
			totCount = count
		}
	}
	fmt.Println(distinctAntiNodes)
	return totCount
}

func calculateAntinodes(locations [][]int, lenMap int, distinctAntiNodes map[string]int) int {
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			xDistance := locations[i][0] - locations[j][0]
			yDistance := locations[i][1] - locations[j][1]

			loc1X := locations[i][0] + xDistance
			loc1Y := locations[i][1] + yDistance
			updateLocationsMap(loc1X, loc1Y, lenMap, distinctAntiNodes)

			loc2X := locations[j][0] - xDistance
			loc2Y := locations[j][1] - yDistance
			updateLocationsMap(loc2X, loc2Y, lenMap, distinctAntiNodes)
		}
	}
	// fmt.Println(distinctAntiNodes)
	return len(distinctAntiNodes)
}

func calculateAntinodesResonantHarmonics(locations [][]int, lenMap int, distinctAntiNodes map[string]int) int {
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			updateLocationsMap(locations[i][0], locations[i][1], lenMap, distinctAntiNodes)
			updateLocationsMap(locations[j][0], locations[j][1], lenMap, distinctAntiNodes)
			xDistance := locations[i][0] - locations[j][0]
			yDistance := locations[i][1] - locations[j][1]

			checkHarmonics := true
			x := xDistance
			y := yDistance
			for checkHarmonics {
				loc1X := locations[i][0] + x
				loc1Y := locations[i][1] + y
				checkHarmonics = updateLocationsMap(loc1X, loc1Y, lenMap, distinctAntiNodes)
				x += xDistance
				y += yDistance
			}

			checkHarmonics = true
			x = xDistance
			y = yDistance
			for checkHarmonics {
				loc2X := locations[j][0] - x
				loc2Y := locations[j][1] - y
				checkHarmonics = updateLocationsMap(loc2X, loc2Y, lenMap, distinctAntiNodes)
				x += xDistance
				y += yDistance
			}
		}
	}
	// fmt.Println(distinctAntiNodes)
	return len(distinctAntiNodes)
}

func updateLocationsMap(x, y, lenMap int, distinctAntiNodes map[string]int) bool {
	if x >= 0 && x <= lenMap-1 && y >= 0 && y <= lenMap-1 {
		antinodeLocation := strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y)}, "-")
		distinctAntiNodes[antinodeLocation] = 1
		return true
	}
	return false
}
