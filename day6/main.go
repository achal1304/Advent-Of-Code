package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var directionSwitch = map[string]string{
	"Left":   "Top",
	"Top":    "Right",
	"Right":  "Bottom",
	"Bottom": "Left",
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pathAndObstacles := [][]string{}
	startLocationX := 0
	startLocationY := 0
	iteration := 0
	for scanner.Scan() {
		pathAndObstacle := []string{}
		line := scanner.Text()
		paths := strings.Split(line, "")
		for j, path := range paths {
			if path == "^" {
				startLocationX = iteration
				startLocationY = j
			}
			pathAndObstacle = append(pathAndObstacle, path)
		}
		pathAndObstacles = append(pathAndObstacles, pathAndObstacle)
		iteration += 1
	}
	fmt.Println(startLocationX, " ", startLocationY)
	pathAndObstacles[startLocationX][startLocationY] = "."
	fmt.Println(pathAndObstacles)
	fmt.Println(calculateUniquePositions(pathAndObstacles, startLocationX, startLocationY))
}

func calculateUniquePositions(pathAndObstacles [][]string, x, y int) int {
	startDirection := "Top"
	visitedDict := make(map[string]int)
	visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y)}, "-")] = 0
outerLoop:
	for x > 0 && x < len(pathAndObstacles)-1 && y > 0 && y < len(pathAndObstacles)-1 {
		// fmt.Println("index ", x, "-", y)
		switch startDirection {
		case "Top":
			for x-1 >= 0 {
				fmt.Println("Top index ", x, "-", y)
				if pathAndObstacles[x-1][y] == "." {
					visitedDict[strings.Join([]string{fmt.Sprint(x - 1), fmt.Sprint(y)}, "-")] = 0
					x -= 1
				} else {
					startDirection = directionSwitch[startDirection]
					fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}
		case "Right":
			for y+1 <= len(pathAndObstacles)-1 {
				fmt.Println("Right index ", x, "-", y)
				if pathAndObstacles[x][y+1] == "." {
					visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y + 1)}, "-")] = 0
					y += 1
				} else {
					startDirection = directionSwitch[startDirection]
					fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}

		case "Bottom":
			for x+1 <= len(pathAndObstacles)-1 {
				fmt.Println("Bottom index ", x, "-", y)
				if pathAndObstacles[x+1][y] == "." {
					visitedDict[strings.Join([]string{fmt.Sprint(x + 1), fmt.Sprint(y)}, "-")] = 0
					x += 1
				} else {
					startDirection = directionSwitch[startDirection]
					fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}
		case "Left":
			for y-1 >= 0 {
				fmt.Println("Left index ", x, "-", y)
				if pathAndObstacles[x][y-1] == "." {
					visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y - 1)}, "-")] = 0
					y -= 1
				} else {
					startDirection = directionSwitch[startDirection]
					fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}
		default:
			fmt.Println("no caswe found")
		}
	}
	return len(visitedDict)
}
