package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Position struct {
	row, col, direction, score int
}

type State struct {
	cost   int
	r, c   int
	dr, dc int
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
type MinHeap []*State

func (pq MinHeap) Len() int           { return len(pq) }
func (pq MinHeap) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq MinHeap) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *MinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *MinHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
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
	h := &MinHeap{}
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

func part2(grid [][]rune, start, end Position) int {
	pq := &MinHeap{}
	heap.Init(pq)

	lowestCost := make(map[[4]int]int)
	bestCost := math.MaxInt
	endStates := [][4]int{}
	backtrack := make(map[[4]int][][4]int)

	// Example starting state
	heap.Push(pq, &State{cost: 0, r: start.row, c: start.col, dr: 0, dc: 1})
	lowestCost[[4]int{start.row, start.col, 0, 1}] = math.MaxInt

	for pq.Len() > 0 {
		state := heap.Pop(pq).(*State)
		cost := state.cost
		r, c := state.r, state.c
		dr, dc := state.dr, state.dc

		costFound, exisits := lowestCost[[4]int{r, c, dr, dc}]
		if !exisits || cost <= costFound {
			lowestCost[[4]int{r, c, dr, dc}] = cost
		} else {
			continue
		}
		// Check if reached destination or certain condition
		if string(grid[r][c]) == "E" {
			if cost <= bestCost {
				bestCost = cost
				endStates = append(endStates, [4]int{r, c, dr, dc})
			}
			break
		}

		for _, direction := range []struct {
			newCost, nr, nc, ndr, ndc int
		}{
			{cost + 1, r + dr, c + dc, dr, dc},
			{cost + 1000, r, c, dc, -dr},
			{cost + 1000, r, c, -dc, dr},
		} {
			if string(grid[direction.nr][direction.nc]) == "#" {
				continue
			}

			lowest, exists := lowestCost[[4]int{direction.nr, direction.nc, direction.ndr, direction.ndc}]
			if !exists || direction.newCost <= lowest {
				lowestCost[[4]int{direction.nr, direction.nc, direction.ndr, direction.ndc}] = direction.newCost
				heap.Push(pq, &State{cost: direction.newCost, r: direction.nr, c: direction.nc, dr: direction.ndr, dc: direction.ndc})
				backtrack[[4]int{direction.nr, direction.nc, direction.ndr, direction.ndc}] =
					append(backtrack[[4]int{direction.nr, direction.nc, direction.ndr, direction.ndc}],
						[4]int{r, c, dr, dc})
			}
		}
	}

	seen := make(map[[4]int]bool)
	seenIndexes := make(map[[2]int]bool)
	for len(endStates) > 0 {
		val := endStates[0]
		endStates = endStates[1:]
		foundState, _ := backtrack[val]
		for _, ele := range foundState {
			if exist := seen[ele]; exist {
				continue
			}
			seen[ele] = true
			seenIndexes[[2]int{ele[0], ele[1]}] = true
			endStates = append(endStates, ele)
		}
	}

	fmt.Println("Safe sit positions:", len(seenIndexes)+1)
	return len(seenIndexes) + 1
}

func main() {
	maze, start, end := readMaze("input.txt")

	if (start == Position{}) || (end == Position{}) {
		fmt.Println("Start or End not found in the maze.")
		return
	}

	// Calculate the minimal score path using Dijkstra's algorithm
	dijkstra(maze, start, end)
	part2(maze, start, end)
	// count := 0
	// fmt.Printf("Total number of tiles on the best path(s): %d\n", count)
	// Print the marked maze with 'O' indicating best path tiles
	// count := 0
	// for r := 0; r < len(maze); r++ {
	// 	for c := 0; c < len(maze[r]); c++ {
	// 		if ans[Position{r, c, 0, 0}] {
	// 			maze[r][c] = 'O'
	// 		}
	// 		fmt.Print(string(maze[r][c]))
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Printf("Total number of tiles on the best path(s): %d\n", len(ans))
}

// func calculateUpdateLowestCost(lowestCost map[Cost]int, current Position) bool {
// 	p := Cost{current.row, current.col, current.direction}
// 	if val, ok := lowestCost[p]; !ok {
// 		lowestCost[p] = current.score
// 		return true
// 	} else {
// 		if current.score <= val {
// 			lowestCost[p] = current.score
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// }
