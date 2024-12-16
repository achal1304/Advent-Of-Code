package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Position struct {
	row, col, direction, score int
}

const (
	North = iota
	East
	South
	West
)

var directions = [4][2]int{
	{-1, 0}, // North
	{0, 1},  // East
	{1, 0},  // South
	{0, -1}, // West
}

var turnRight = [4]int{East, South, West, North}
var turnLeft = [4]int{West, North, East, South}

// Min-Heap to prioritize positions with the lowest score
type MinHeap []Position

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i].score < (*h)[j].score }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Position))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func readMaze(filename string) ([][]rune, Position, Position) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, Position{}, Position{}
	}
	defer file.Close()

	var maze [][]rune
	var start, end Position
	scanner := bufio.NewScanner(file)

	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		maze = append(maze, []rune(line))
		for c, ch := range line {
			if ch == 'S' {
				start = Position{row: r, col: c, direction: East, score: 0}
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

func isValid(maze [][]rune, r, c int) bool {
	return r >= 0 && r < len(maze) && c >= 0 && c < len(maze[0]) && maze[r][c] != '#'
}

func dijkstra(maze [][]rune, start, end Position) int {
	h := &MinHeap{start}
	heap.Init(h)
	visited := make(map[string]bool)
	for h.Len() > 0 {
		current := heap.Pop(h).(Position)

		if current.row == end.row && current.col == end.col {
			return current.score
		}

		state := fmt.Sprintf("%d,%d,%d", current.row, current.col, current.direction)
		if visited[state] {
			continue
		}
		visited[state] = true

		forwardR := current.row + directions[current.direction][0]
		forwardC := current.col + directions[current.direction][1]
		if isValid(maze, forwardR, forwardC) {
			heap.Push(h, Position{row: forwardR, col: forwardC, direction: current.direction, score: current.score + 1})
		}

		rightDir := turnRight[current.direction]
		heap.Push(h, Position{row: current.row, col: current.col, direction: rightDir, score: current.score + 1000})

		leftDir := turnLeft[current.direction]
		heap.Push(h, Position{row: current.row, col: current.col, direction: leftDir, score: current.score + 1000})
	}
	return -1
}

func dfs(maze [][]rune, current Position, end Position, bestScore int, visited map[string]bool, path map[Position]bool) {
	// If we reach the end tile with the best score, mark the path
	if current.row == end.row && current.col == end.col && current.score == bestScore {
		path[current] = true
		return
	}

	// Mark the current position as visited
	visitedKey := fmt.Sprintf("%d,%d,%d", current.row, current.col, current.direction)
	if visited[visitedKey] {
		return
	}
	visited[visitedKey] = true
	path[current] = true

	// Move forward in the current direction
	forwardRow := current.row + directions[current.direction][0]
	forwardCol := current.col + directions[current.direction][1]
	if isValid(maze, forwardRow, forwardCol) {
		dfs(maze, Position{row: forwardRow, col: forwardCol, direction: current.direction, score: current.score + 1}, end, bestScore, visited, path)
	}

	// Turn right and explore
	rightDirection := turnRight[current.direction]
	dfs(maze, Position{row: current.row, col: current.col, direction: rightDirection, score: current.score + 1000}, end, bestScore, visited, path)

	// Turn left and explore
	leftDirection := turnLeft[current.direction]
	dfs(maze, Position{row: current.row, col: current.col, direction: leftDirection, score: current.score + 1000}, end, bestScore, visited, path)
}

func main() {
	maze, start, end := readMaze("input.txt")

	if (start == Position{}) || (end == Position{}) {
		fmt.Println("Start or End not found in the maze.")
		return
	}

	// Calculate the minimal score path using Dijkstra's algorithm
	bestScore := dijkstra(maze, start, end)

	visited := make(map[string]bool) // To keep track of visited nodes in DFS
	path := make(map[Position]bool)  // To store the positions that are part of any best path

	// Start DFS from the start tile
	dfs(maze, start, end, bestScore, visited, path)

	count := len(path)
	fmt.Printf("Total number of tiles on the best path(s): %d\n", count)
	// Print the marked maze with 'O' indicating best path tiles
	// count := 0
	for r := 0; r < len(maze); r++ {
		for c := 0; c < len(maze[r]); c++ {
			if maze[r][c] == 'O' {
				count++
			}
			fmt.Print(string(maze[r][c]))
		}
		fmt.Println()
	}

	fmt.Printf("Total number of tiles on the best path(s): %d\n", count)
}
