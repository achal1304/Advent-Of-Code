package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

var DirectionMap = map[string]Point{
	"^": {-1, 0},
	">": {0, 1},
	"<": {0, -1},
	"v": {1, 0},
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("error opening file ", err)
		return
	}
	defer file.Close()
	botPosition, wallsMap, startObjPositions, endObjPositions, botMovements := scanRobotMap(file)
	fmt.Println(predictFinalPoisitions(botPosition, wallsMap, startObjPositions, endObjPositions, botMovements))

}

func scanRobotMap(file *os.File) (Point, map[Point]bool, [][]bool, [][]bool, []string) {
	scanner := bufio.NewScanner(file)
	wallsMap := make(map[Point]bool)
	startObjPositions := [][]bool{}
	endObjPositions := [][]bool{}
	botPosition := Point{0, 0}
	botMovements := []string{}

	iteration := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(iteration)
		if line != "" {
			startObjPos := make([]bool, len(line)*2)
			endObjPos := make([]bool, len(line)*2)
			for i, obj := range line {
				if string(obj) == "#" {
					wallsMap[Point{iteration, i * 2}] = true
					wallsMap[Point{iteration, (i * 2) + 1}] = true
				} else if string(obj) == "O" {
					// fmt.Println("found O at ", iteration, i)
					startObjPos[i*2] = true
					endObjPos[(i*2)+1] = true
					// fmt.Println(objPos)
				} else if string(obj) == "@" {
					botPosition = Point{iteration, i * 2}
				}
			}
			startObjPositions = append(startObjPositions, startObjPos)
			endObjPositions = append(endObjPositions, endObjPos)
			iteration++
		} else {
			for scanner.Scan() {
				line := scanner.Text()
				for _, ele := range line {
					botMovements = append(botMovements, string(ele))
				}
			}
		}
	}
	PrintObjectPositions(botPosition, startObjPositions, endObjPositions)
	return botPosition, wallsMap, startObjPositions, endObjPositions, botMovements
}

func predictFinalPoisitions(botPosition Point,
	wallsMap map[Point]bool,
	startObjPositions [][]bool,
	endObjPositions [][]bool,
	botMovements []string) int {
outerLoop:
	for _, ele := range botMovements {
		// fmt.Println("iteration ", i)
		moveDirection := DirectionMap[ele]
		currX := botPosition.x + moveDirection.x
		currY := botPosition.y + moveDirection.y
		if wallsMap[Point{currX, currY}] {
			continue
		}
		if moveDirection.x == 0 {
			if startObjPositions[currX][currY] || endObjPositions[currX][currY] {
				nextObjX := currX + moveDirection.x
				nextObjY := currY + moveDirection.y
				for startObjPositions[nextObjX][nextObjY] || endObjPositions[nextObjX][nextObjY] {
					nextObjX += moveDirection.x
					nextObjY += moveDirection.y
				}
				if wallsMap[Point{nextObjX, nextObjY}] {
					continue
				} else {
					xStep, yStep := 1, 1
					if nextObjX >= currX {
						xStep = -1
					}
					if nextObjY >= currY {
						yStep = -1
					}

					// Loop over the range based on the calculated step directions
					for x := nextObjX; x != currX+xStep; x += xStep {
						for y := nextObjY; y != currY+yStep; y += yStep {
							if startObjPositions[x-moveDirection.x][y-moveDirection.y] {
								startObjPositions[x][y] = startObjPositions[x-moveDirection.x][y-moveDirection.y]
								startObjPositions[x-moveDirection.x][y-moveDirection.y] = false
							} else if endObjPositions[x-moveDirection.x][y-moveDirection.y] {
								endObjPositions[x][y] = endObjPositions[x-moveDirection.x][y-moveDirection.y]
								endObjPositions[x-moveDirection.x][y-moveDirection.y] = false
							}
						}
					}
				}
			}
		} else if moveDirection.y == 0 {
			if startObjPositions[currX][currY] || endObjPositions[currX][currY] {
				queue := []Point{}
				elementsToBeMoved := []Point{}
				visitedInQueue := make(map[Point]bool)
				if endObjPositions[currX][currY] {
					queue = append(queue, Point{currX, currY})
					queue = append(queue, Point{currX, currY - 1})
				} else {
					queue = append(queue, Point{currX, currY})
					queue = append(queue, Point{currX, currY + 1})
				}
				for len(queue) > 0 {
					qLen := len(queue)
					elementsToBeMoved = append(elementsToBeMoved, queue...)
				innerLoop:
					for i := 0; i < qLen; i++ {
						nextX := queue[i].x + moveDirection.x
						nextY := queue[i].y + moveDirection.y
						if wallsMap[Point{nextX, nextY}] {
							continue outerLoop
						}
						if visitedInQueue[Point{nextX, nextY}] {
							continue innerLoop
						}
						if endObjPositions[nextX][nextY] {
							queue = append(queue, Point{nextX, nextY})
							queue = append(queue, Point{nextX, nextY - 1})
							visitedInQueue[Point{nextX, nextY}] = true
							visitedInQueue[Point{nextX, nextY - 1}] = true
						} else if startObjPositions[nextX][nextY] {
							queue = append(queue, Point{nextX, nextY})
							queue = append(queue, Point{nextX, nextY + 1})
							visitedInQueue[Point{nextX, nextY}] = true
							visitedInQueue[Point{nextX, nextY + 1}] = true
						}
					}
					queue = queue[qLen:]
					// fmt.Println("updated queue ", queue)
				}
				// fmt.Println("ele to be move ", elementsToBeMoved)

				for i := len(elementsToBeMoved) - 1; i >= 0; i-- {
					if startObjPositions[elementsToBeMoved[i].x][elementsToBeMoved[i].y] {
						startObjPositions[elementsToBeMoved[i].x+moveDirection.x][elementsToBeMoved[i].y+moveDirection.y] = startObjPositions[elementsToBeMoved[i].x][elementsToBeMoved[i].y]
						startObjPositions[elementsToBeMoved[i].x][elementsToBeMoved[i].y] = false
					} else if endObjPositions[elementsToBeMoved[i].x][elementsToBeMoved[i].y] {
						endObjPositions[elementsToBeMoved[i].x+moveDirection.x][elementsToBeMoved[i].y+moveDirection.y] = endObjPositions[elementsToBeMoved[i].x][elementsToBeMoved[i].y]
						endObjPositions[elementsToBeMoved[i].x][elementsToBeMoved[i].y] = false
					}
				}
			}
		}
		botPosition.x = currX
		botPosition.y = currY
		// PrintSingleObjPositions(botPosition, startObjPositions)
		// PrintSingleObjPositions(botPosition, endObjPositions)
		// PrintObjectPositions(botPosition, startObjPositions, endObjPositions)
	}

	sum := calculateGPSSum(startObjPositions)
	return sum
}

func calculateGPSSum(objPositions [][]bool) int {
	totCount := 0
	for i := 0; i < len(objPositions); i++ {
		for j := 0; j < len(objPositions[0]); j++ {
			if objPositions[i][j] {
				totCount += (100 * i) + j
			}
		}
	}
	return totCount
}

func PrintObjectPositions(botPosition Point, startObjPositions, endObjPositions [][]bool) {
	fmt.Println("bot Position ", botPosition)
	for i, ele := range startObjPositions {
		for j, _ := range ele {
			if i == botPosition.x && j == botPosition.y {
				fmt.Print("@")
			} else {
				if startObjPositions[i][j] || endObjPositions[i][j] {
					fmt.Print("O")
				} else {
					fmt.Print("-")
				}
			}
		}
		fmt.Println()
	}
}

func PrintSingleObjPositions(botPosition Point, objPos [][]bool) {
	fmt.Println("bot Position ", botPosition)
	for i, ele := range objPos {
		for j, _ := range ele {
			if i == botPosition.x && j == botPosition.y {
				fmt.Print("@")
			} else {
				if objPos[i][j] {
					fmt.Print("O")
				} else {
					fmt.Print("-")
				}
			}
		}
		fmt.Println()
	}
}
