package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x, y, botNumber int
}
type Velocity struct {
	vx, vy int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error operning file ", err)
	}
	defer file.Close()
	mapBathroom := scanInputWithRegex(file)
	fmt.Println(mapBathroom)
	fmt.Println(calculatePositionsAfter100Seconds(mapBathroom, 101, 103))
}

func scanInputWithRegex(file *os.File) map[Point]Velocity {
	mapBathroom := make(map[Point]Velocity)

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	iteration := 0
	for scanner.Scan() {
		iteration += 1
		text := scanner.Text()
		parts := re.FindStringSubmatch(text)
		x, _ := strconv.Atoi(parts[1])
		y, _ := strconv.Atoi(parts[2])
		vx, _ := strconv.Atoi(parts[3])
		vy, _ := strconv.Atoi(parts[4])

		mapBathroom[Point{x, y, iteration}] = Velocity{vx, vy}
	}
	return mapBathroom
}

func calculatePositionsAfter100Seconds(mapBathroom map[Point]Velocity, width, height int) int {
	const MaxSeconds = 100
	quadrants := make([]int, 5)
	for point, velocity := range mapBathroom {
		newPointX := point.x + velocity.vx*MaxSeconds
		newPointY := point.y + velocity.vy*MaxSeconds

		for newPointX < 0 || newPointX > width-1 {
			if newPointX > width-1 {
				newPointX -= width
			} else if newPointX < 0 {
				newPointX += width
			}
		}

		for newPointY < 0 || newPointY > height-1 {
			if newPointY > height-1 {
				newPointY -= height
			} else if newPointY < 0 {
				newPointY += height
			}
		}
		quadrant := decideQuadrant(newPointX, newPointY, width, height)
		quadrants[quadrant] += 1

	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func decideQuadrant(x, y, width, height int) int {
	partitionX := width / 2
	partitionY := height / 2

	if x < partitionX && y < partitionY {
		return 0
	} else if x < partitionX && y > partitionY {
		return 1
	} else if x > partitionX && y < partitionY {
		return 2
	} else if x > partitionX && y > partitionY {
		return 3
	}
	return 4
}
