package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var directions = []Point{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

// Part 2
var memo map[Point]int

// DFS function to count the distinct hiking trails starting from a given position
func dfs(memo map[Point]int, mapData [][]int, x, y, rows, cols int) int {
	if x < 0 || y < 0 || x >= rows || y >= cols {
		return 0 // Out of bounds
	}

	if mapData[x][y] == 9 {
		return 1 // Reached height 9, a valid trail
	}

	if mapData[x][y] == -1 {
		return 0 // Invalid tile or visited
	}

	// Memoization check
	if val, found := memo[Point{x, y}]; found {
		return val // Return the previously calculated number of paths
	}

	temp := mapData[x][y]

	totalPaths := 0
	// Explore all four directions
	for _, direction := range directions {
		nx, ny := x+direction.x, y+direction.y
		if nx >= 0 && ny >= 0 && nx < rows && ny < cols && mapData[nx][ny] == temp+1 {
			totalPaths += dfs(memo, mapData, nx, ny, rows, cols)
		}
	}

	// Memoize the result
	memo[Point{x, y}] = totalPaths

	return totalPaths
}

// Part 1
// BFS to calculate the score from a trailhead (height 0 position)
func bfs(mapData [][]int, start Point, rows, cols int) int {
	queue := []Point{start}
	visited := make(map[Point]bool)
	visited[start] = true
	reachable9s := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// If the current tile is 9, increment the score
		if mapData[current.x][current.y] == 9 {
			reachable9s++
		}

		// Explore all four directions
		for _, direction := range directions {
			nx, ny := current.x+direction.x, current.y+direction.y
			if nx >= 0 && ny >= 0 && nx < rows && ny < cols && !visited[Point{nx, ny}] {
				if mapData[nx][ny] == mapData[current.x][current.y]+1 {
					visited[Point{nx, ny}] = true
					queue = append(queue, Point{nx, ny})
				}
			}
		}
	}
	return reachable9s
}

// Function to calculate the total score from all trailheads
func calculateTrailScores(mapData [][]int) int {
	rows := len(mapData)
	cols := len(mapData[0])
	totalScore := 0

	// Iterate over all positions to find trailheads (height 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mapData[i][j] == 0 {
				// For each trailhead, use BFS to calculate the score

				// Part 1
				// score := bfs(mapData, Point{i, j}, rows, cols)

				// Part 2
				memo := make(map[Point]int)
				score := dfs(memo, mapData, i, j, rows, cols)
				totalScore += score
			}
		}
	}

	return totalScore
}

func main() {
	// Calculate and print the total score
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error while opening file ", err)
	}
	defer file.Close()
	hikingMap := scanInputTrailheads(file)
	totalScore := calculateTrailScores(hikingMap)
	fmt.Println("Total Score:", totalScore)
}

func scanInputTrailheads(file *os.File) [][]int {
	var hikingMap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var hiking []int
		text := scanner.Text()
		if text == "" {
			break
		}
		parts := strings.Split(text, "")
		for _, ele := range parts {
			if ele == "." {
				hiking = append(hiking, -999)
				continue
			}
			num, err := strconv.Atoi(ele)
			if err != nil {
				log.Fatal("unable to parse parts ", err)
			}
			hiking = append(hiking, num)
		}
		hikingMap = append(hikingMap, hiking)
	}

	fmt.Println(hikingMap)
	return hikingMap
}
