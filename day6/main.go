package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/achal1304/Advent-Of-Code/utils"
)

var directionSwitch = map[string]string{
	"Left":   "Top",
	"Top":    "Right",
	"Right":  "Bottom",
	"Bottom": "Left",
}

const Top = "Top"
const Bottom = "Bottom"
const Right = "Right"
const Left = "Left"

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
	startX := x
	startY := y
	visitedDict := make(map[string]int)
	visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y)}, "-")] = 0
outerLoop:
	for x > 0 && x < len(pathAndObstacles)-1 && y > 0 && y < len(pathAndObstacles)-1 {
		// fmt.Println("index ", x, "-", y)
		switch startDirection {
		case "Top":
			for x-1 >= 0 {
				// fmt.Println("Top index ", x, "-", y)
				if pathAndObstacles[x-1][y] == "." {
					visitedDict[strings.Join([]string{fmt.Sprint(x - 1), fmt.Sprint(y)}, "-")] = 0
					x -= 1
				} else {
					startDirection = directionSwitch[startDirection]
					// fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}
		case "Right":
			for y+1 <= len(pathAndObstacles)-1 {
				// fmt.Println("Right index ", x, "-", y)
				if pathAndObstacles[x][y+1] == "." {
					visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y + 1)}, "-")] = 0
					y += 1
				} else {
					startDirection = directionSwitch[startDirection]
					// fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}

		case "Bottom":
			for x+1 <= len(pathAndObstacles)-1 {
				// fmt.Println("Bottom index ", x, "-", y)
				if pathAndObstacles[x+1][y] == "." {
					visitedDict[strings.Join([]string{fmt.Sprint(x + 1), fmt.Sprint(y)}, "-")] = 0
					x += 1
				} else {
					startDirection = directionSwitch[startDirection]
					// fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}
		case "Left":
			for y-1 >= 0 {
				// fmt.Println("Left index ", x, "-", y)
				if pathAndObstacles[x][y-1] == "." {
					visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y - 1)}, "-")] = 0
					y -= 1
				} else {
					startDirection = directionSwitch[startDirection]
					// fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}
		default:
			fmt.Println("no caswe found")
		}
	}
	fmt.Println("answer part 1", len(visitedDict))

	return blockTheGuardLoops(pathAndObstacles, visitedDict, startX, startY)
}

func checkIfBlockedOrLoop(pathAndObstacles [][]string, x, y int) bool {
	startDirection := Top
	visitedDict := make(map[string]string)
	visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y)}, "-")] = Top
outerLoop:
	for x > 0 && x < len(pathAndObstacles)-1 && y > 0 && y < len(pathAndObstacles)-1 {
		// fmt.Println("index ", x, "-", y)
		switch startDirection {
		case Top:
			for x-1 >= 0 {
				// fmt.Println("Top index ", x, "-", y)
				if pathAndObstacles[x-1][y] == "." {
					if visitedDir, ok := visitedDict[strings.Join([]string{fmt.Sprint(x - 1), fmt.Sprint(y)}, "-")]; ok {
						if visitedDir == Top {
							return true
						}
					} else {
						visitedDict[strings.Join([]string{fmt.Sprint(x - 1), fmt.Sprint(y)}, "-")] = Top
					}
					x -= 1
				} else {
					startDirection = directionSwitch[startDirection]
					// fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}
		case Right:
			for y+1 <= len(pathAndObstacles)-1 {
				// fmt.Println("Right index ", x, "-", y)
				if pathAndObstacles[x][y+1] == "." {
					if visitedDir, ok := visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y + 1)}, "-")]; ok {
						if visitedDir == Right {
							return true
						}
					} else {
						visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y + 1)}, "-")] = Right
					}
					y += 1
				} else {
					startDirection = directionSwitch[startDirection]
					// fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}

		case Bottom:
			for x+1 <= len(pathAndObstacles)-1 {
				// fmt.Println("Bottom index ", x, "-", y)
				if pathAndObstacles[x+1][y] == "." {
					if visitedDir, ok := visitedDict[strings.Join([]string{fmt.Sprint(x + 1), fmt.Sprint(y)}, "-")]; ok {
						if visitedDir == Bottom {
							return true
						}
					} else {
						visitedDict[strings.Join([]string{fmt.Sprint(x + 1), fmt.Sprint(y)}, "-")] = Bottom
					}
					x += 1
				} else {
					startDirection = directionSwitch[startDirection]
					// fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}
		case Left:
			for y-1 >= 0 {
				// fmt.Println("Left index ", x, "-", y)
				if pathAndObstacles[x][y-1] == "." {
					if visitedDir, ok := visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y - 1)}, "-")]; ok {
						if visitedDir == Left {
							return true
						}
					} else {
						visitedDict[strings.Join([]string{fmt.Sprint(x), fmt.Sprint(y - 1)}, "-")] = Left
					}
					y -= 1
				} else {
					startDirection = directionSwitch[startDirection]
					// fmt.Println("startdir ", startDirection)
					continue outerLoop
				}
			}
		default:
			fmt.Println("no caswe found")
		}
	}
	// fmt.Println("answer part 1", len(visitedDict))
	return false
}

func blockTheGuardLoops(pathAndObstacles [][]string, pathOfGuard map[string]int, startX, startY int) int {
	fmt.Println(startX)
	fmt.Println(startY)
	fmt.Println(pathOfGuard)
	var wg sync.WaitGroup
	ch := make(chan int, len(pathAndObstacles)*len(pathAndObstacles))
	count := 0
	for pathIndexes, _ := range pathOfGuard {
		blockX, blockY := utils.GetXAndYFromString(pathIndexes)
		if blockX == startX && blockY == startY {
			continue
		}

		var newBlockedPathAndObstacles [][]string

		for _, row := range pathAndObstacles {
			newRow := make([]string, len(row))
			copy(newRow, row)
			newBlockedPathAndObstacles = append(newBlockedPathAndObstacles, newRow)
		}
		newBlockedPathAndObstacles[blockX][blockY] = "#"
		wg.Add(1)

		go func(blockX int, blockY int) {
			defer wg.Done()
			if checkIfBlockedOrLoop(newBlockedPathAndObstacles, startX, startY) {
				fmt.Println("** WOOHOOO Guard blocked at position ", blockX, " ", blockY)
				ch <- 1
			}
		}(blockX, blockY)

	}
	wg.Wait()
	close(ch)

	for ans := range ch {
		count += ans
	}
	return count
}
