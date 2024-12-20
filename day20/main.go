package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/achal1304/Advent-Of-Code/utils"
)

type Position struct {
	row, col int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error while reading file ", err)
	}
	defer file.Close()

	maze, start, end := readMaze(file)
	// Part 1
	// fmt.Println(maze)
	pathMapSeconds := scanPath(maze, start, end)
	// skipPaths(maze, pathMapSeconds)

	// PART 2
	part2(maze, start, end, pathMapSeconds)
}

func skipPaths(maze [][]rune, pathMapSeconds map[Position]int) {
	savingsCount := make(map[int]int)

	for key, _ := range pathMapSeconds {
		for _, ele := range [][]int{{0, 2}, {2, 0}, {-2, 0}, {0, -2}} {
			newele := Position{key.row + ele[0], key.col + ele[1]}
			if isValid(maze, newele.row, newele.col) && pathMapSeconds[newele] != 0 &&
				pathMapSeconds[newele] > pathMapSeconds[key]+2 {
				diff := pathMapSeconds[newele] - pathMapSeconds[key] - 2
				if diff >= 100 {
					savingsCount[diff]++
				}
			}
		}
	}

	fmt.Println(savingsCount)

	count := 0
	for _, v := range savingsCount {
		count += v
	}
	fmt.Println(count)
}

func scanPath(maze [][]rune, start Position, end Position) map[Position]int {
	s := start
	pathMap := make(map[Position]int)
	pathMap[s] = 0
	pathCount := 0
	visited := make(map[Position]bool)
	visited[s] = true
outerLoop:
	for s != end {
		for _, ele := range [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			newele := Position{s.row + ele[0], s.col + ele[1]}
			if isValid(maze, newele.row, newele.col) && !visited[newele] {
				pathCount++
				pathMap[newele] = pathCount
				s = newele
				visited[newele] = true
				continue outerLoop
			}
			visited[newele] = true
		}
	}
	return pathMap
}

func isValid(maze [][]rune, r, c int) bool {
	return r >= 0 && r < len(maze) && c >= 0 && c < len(maze[0]) && string(maze[r][c]) != "#"
}

func readMaze(file *os.File) ([][]rune, Position, Position) {
	var maze [][]rune
	var start, end Position
	scanner := bufio.NewScanner(file)

	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		maze = append(maze, []rune(line))
		for c, ch := range line {
			if ch == 'S' {
				start = Position{row: r, col: c}
			} else if ch == 'E' {
				end = Position{row: r, col: c}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	return maze, start, end
}

func part2(maze [][]rune, start, end Position, pathMapsDistance map[Position]int) {
	cheatScore := make(map[int]int)
	for k, v := range pathMapsDistance {
		manhattanPoints := manhattanRange(k, 20, len(maze[0]), len(maze))
		for _, cheats := range manhattanPoints {
			if _, ok := pathMapsDistance[cheats]; ok {
				if (string(maze[cheats.row][cheats.col]) == "." || cheats == end) &&
					pathMapsDistance[cheats] > v+1 {
					savedSec := pathMapsDistance[cheats] - v - manhattanDistance(k, cheats)
					cheatScore[savedSec]++
				}
			}
		}
	}

	fmt.Println(cheatScore)
	totScore := 0
	for k, v := range cheatScore {
		if k >= 100 {
			totScore += v
		}
	}
	fmt.Println(totScore)
}

func manhattanDistance(p1, p2 Position) int {
	return utils.AbsInt(p1.row-p2.row) + utils.AbsInt(p1.col-p2.col)
}

func manhattanRange(pos Position, distance int, maxX, maxY int) []Position {
	var points []Position
	for dy := -distance; dy <= distance; dy++ {
		for dx := -distance; dx <= distance; dx++ {
			if utils.AbsInt(dy)+utils.AbsInt(dx) <= distance {
				nx, ny := pos.row+dx, pos.col+dy
				if nx >= 0 && ny >= 0 && nx < maxX && ny < maxY {
					points = append(points, Position{nx, ny})
				}
			}
		}
	}
	return points
}
