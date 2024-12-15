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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file ", err)
		return
	}
	defer file.Close()
	botPosition, wallsMap, objectPositions, botMovements := scanRobotMap(file)
	fmt.Println(predictFinalPoisitions(botPosition, wallsMap, objectPositions, botMovements))

}

func scanRobotMap(file *os.File) (Point, map[Point]bool, [][]bool, []string) {
	scanner := bufio.NewScanner(file)
	wallsMap := make(map[Point]bool)
	objectPositions := [][]bool{}
	botPosition := Point{0, 0}
	botMovements := []string{}

	iteration := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(iteration)
		if line != "" {
			objPos := make([]bool, len(line))
			for i, obj := range line {
				if string(obj) == "#" {
					wallsMap[Point{iteration, i}] = true
				} else if string(obj) == "O" {
					// fmt.Println("found O at ", iteration, i)
					objPos[i] = true
					// fmt.Println(objPos)
				} else if string(obj) == "@" {
					botPosition = Point{iteration, i}
				}
			}
			objectPositions = append(objectPositions, objPos)
			// fmt.Println(objectPositions)
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

	return botPosition, wallsMap, objectPositions, botMovements
}

func predictFinalPoisitions(botPosition Point,
	wallsMap map[Point]bool,
	objectPositions [][]bool,
	botMovements []string) int {
	for _, ele := range botMovements {
		// fmt.Println("iteration ", i)
		moveDirection := DirectionMap[ele]
		currX := botPosition.x + moveDirection.x
		currY := botPosition.y + moveDirection.y
		// fmt.Println("current positions ", currX, currY)
		if wallsMap[Point{currX, currY}] {
			continue
		} else if objectPositions[currX][currY] {
			nextObjX := currX + moveDirection.x
			nextObjY := currY + moveDirection.y
			for objectPositions[nextObjX][nextObjY] {
				nextObjX += moveDirection.x
				nextObjY += moveDirection.y
			}
			// fmt.Println("current positions ", currX, currY)
			// fmt.Println("found object ahead next posi ", nextObjX, nextObjY)
			if wallsMap[Point{nextObjX, nextObjY}] {
				continue
			} else {
				if nextObjX >= currX && nextObjY >= currY {
					for x := nextObjX; x >= currX; x-- {
						for y := nextObjY; y >= currY; y-- {
							// fmt.Println("updating next positions ", x, y)
							objectPositions[x][y] = objectPositions[x-moveDirection.x][y-moveDirection.y]
							objectPositions[x-moveDirection.x][y-moveDirection.y] = false
						}
					}
				} else {
					for x := nextObjX; x <= currX; x++ {
						for y := nextObjY; y <= currY; y++ {
							// fmt.Println("updating next positions ", x, y)
							objectPositions[x][y] = objectPositions[x-moveDirection.x][y-moveDirection.y]
							objectPositions[x-moveDirection.x][y-moveDirection.y] = false
						}
					}
				}
			}
		}
		botPosition.x = currX
		botPosition.y = currY
	}

	sum := calculateGPSSum(objectPositions)
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
