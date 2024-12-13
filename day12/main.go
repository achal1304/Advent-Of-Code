package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y int
	num  string
}

type EdgePoint struct {
	x, y float32
}

var directions = []Point{
	{-1, 0, ""}, // Up
	{1, 0, ""},  // Down
	{0, -1, ""}, // Left
	{0, 1, ""},  // Right
}

// Part 2
var memo map[Point]int

// Part 1
// BFS to calculate the score from a trailhead (height 0 position)
func bfs(mapData [][]string, start Point, rows, cols int, visited map[Point]bool, regions map[string]map[Point]bool) int {
	queue := []Point{start}
	// visited := make(map[Point]bool)
	visited[start] = true
	perimeter := 4
	area := 1

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Explore all four directions
		for _, direction := range directions {
			nx, ny := current.x+direction.x, current.y+direction.y
			if nx >= 0 && ny >= 0 && nx < rows && ny < cols && !visited[Point{nx, ny, mapData[nx][ny]}] {
				if mapData[nx][ny] == mapData[current.x][current.y] {
					checkPoint := Point{nx, ny, mapData[nx][ny]}
					visitedPerimeter := calculateVisitedPerimter(mapData, checkPoint, rows, cols, visited)
					// fmt.Println("visited Perimeter for ", checkPoint, visitedPerimeter, visited)
					area += 1
					perimeter += (4 - (2 * visitedPerimeter))
					visited[checkPoint] = true
					regions[mapData[start.x][start.y]+fmt.Sprint(start.x)+fmt.Sprint(start.y)][checkPoint] = true
					queue = append(queue, checkPoint)
				}
			}
		}
	}
	// fmt.Println("calcualted for ", mapData[start.x][start.y])
	// fmt.Println("calcualted permieter for ", perimeter)
	// fmt.Println("calcualted area for ", area)
	return perimeter * area
}

func calculateVisitedPerimter(mapData [][]string, start Point, rows, cols int, visited map[Point]bool) int {
	count := 0
	for _, direction := range directions {
		nx, ny := start.x+direction.x, start.y+direction.y
		if nx >= 0 && ny >= 0 && nx < rows && ny < cols && mapData[nx][ny] == mapData[start.x][start.y] {
			if _, ok := visited[Point{nx, ny, mapData[nx][ny]}]; ok {
				// fmt.Println("** visited ** ", Point{nx, ny, mapData[nx][ny]})
				count += 1
			}
		}
	}
	return count
}

// Function to calculate the total score from all trailheads
func calculatePriceForFencing(mapData [][]string) int {
	rows := len(mapData)
	cols := len(mapData[0])
	totalScore := 0
	visitedMap := make(map[Point]bool, len(mapData)*len(mapData))
	regions := make(map[string]map[Point]bool)

	// Iterate over all positions to find trailheads (height 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			p := Point{i, j, mapData[i][j]}
			_, ok := visitedMap[p]
			if !ok {
				if _, ok := regions[mapData[i][j]]; ok {
					regions[mapData[i][j]+fmt.Sprint(i)+fmt.Sprint(j)][p] = true
				} else {
					regions[mapData[i][j]+fmt.Sprint(i)+fmt.Sprint(j)] = map[Point]bool{p: true}
				}
				// For each trailhead, use BFS to calculate the score

				// Part 1
				score := bfs(mapData, p, rows, cols, visitedMap, regions)
				totalScore += score
			}
		}
	}
	// fmt.Println("regions ", regions)

	// Part 1
	// fmt.Println(totalScore)

	// Part 2
	fmt.Println(calculatePriceWithSides(regions, rows, cols))

	return totalScore
}

func main() {
	// Calculate and print the total score
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error while opening file ", err)
	}
	defer file.Close()
	gardenMap := scanInputGardenMap(file)

	calculatePriceForFencing(gardenMap)
}

func calculatePriceWithSides(regions map[string]map[Point]bool, rows, cols int) int {
	totCount := 0
	for k, v := range regions {
		// fmt.Println(k, v)
		visitedEdges := make(map[EdgePoint]bool)
		edges := make(map[EdgePoint]Point)
		area := len(v)
		for point, _ := range v {
			for _, direction := range directions {
				nx, ny := point.x+direction.x, point.y+direction.y
				if _, ok := v[Point{nx, ny, point.num}]; ok {
					continue
				} else {
					edgeX := float32(point.x+nx) / 2
					edgeY := float32(point.y+ny) / 2
					edges[EdgePoint{edgeX, edgeY}] = Point{nx - int(edgeX), ny - int(edgeY), k}
				}
			}
		}

		// fmt.Println(edges)
		edgeCount := 0
		for edge, direction := range edges {
			if _, ok := visitedEdges[edge]; !ok {
				// fmt.Println("searching for new edge ", edge)
				edgeCount += 1
				// fmt.Println("conditions ", edge.x-float32(math.Floor(float64(edge.x))))
				if edge.x-float32(math.Floor(float64(edge.x))) != 0.5 {
					for _, rowSearchIndex := range []int{-1, 1} {
						nx, ny := edge.x+float32(rowSearchIndex), edge.y
						for edges[EdgePoint{nx, ny}] == direction {
							// fmt.Println(direction)
							// fmt.Println(nx, ny)
							visitedEdges[EdgePoint{nx, ny}] = true
							nx += float32(rowSearchIndex)
						}
					}
				} else {
					for _, colSearchIndex := range []int{-1, 1} {
						nx, ny := edge.x, edge.y+float32(colSearchIndex)
						// fmt.Println(direction)
						for edges[EdgePoint{nx, ny}] == direction {
							// fmt.Println(nx, ny)
							visitedEdges[EdgePoint{nx, ny}] = true
							ny += float32(colSearchIndex)
						}
					}
				}
			}
		}
		// fmt.Println("edge count ", edgeCount)
		// fmt.Println("area count ", area)
		totCount += (area * edgeCount)
	}
	return totCount
}

func scanInputGardenMap(file *os.File) [][]string {
	var gardenMap [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var plants []string
		text := scanner.Text()
		if text == "" {
			break
		}
		parts := strings.Split(text, "")
		for _, ele := range parts {
			plants = append(plants, ele)
		}
		gardenMap = append(gardenMap, plants)
	}

	fmt.Println(gardenMap)
	return gardenMap
}
