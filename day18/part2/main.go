package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/achal1304/Advent-Of-Code/utils"
)

type Position struct {
	Row, Col, score int
}

const STARTFROM = 2878

type State struct {
	cost   int
	r, c   int
	dr, dc int
}

var directions = [4][2]int{
	{-1, 0}, // North
	{0, 1},  // East
	{1, 0},  // South
	{0, -1}, // West
}

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

func readMaze(filename string) ([][]bool, Position, Position, []Position) {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gridSize := 71 // Define grid size
	grid := make([][]bool, gridSize)
	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}

	scanner := bufio.NewScanner(file)
	iteration := 0
	allBytes := []Position{}
	for scanner.Scan() {
		iteration++
		line := scanner.Text()
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			panic("Invalid input format")
		}
		x, err1 := strconv.Atoi(coords[0])
		y, err2 := strconv.Atoi(coords[1])
		if err1 != nil || err2 != nil {
			panic("Invalid coordinate values")
		}
		if iteration <= STARTFROM {
			grid[y][x] = true
		} else {
			allBytes = append(allBytes, Position{y, x, 0})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// fmt.Println(grid)
	start := Position{0, 0, 0}
	end := Position{gridSize - 1, gridSize - 1, 0}
	// part2(grid, start, end)
	return grid, start, end, allBytes
}

func isValid(maze [][]bool, r, c int) bool {
	return r >= 0 && r < len(maze) && c >= 0 && c < len(maze[0]) && maze[r][c] == false
}

func part2(grid [][]bool, start, end Position) map[[2]int]bool {
	pq := &MinHeap{}
	heap.Init(pq)

	lowestCost := make(map[[2]int]int)
	bestCost := math.MaxInt
	endStates := [][2]int{}
	backtrack := make(map[[2]int][][2]int)

	// Example starting state
	heap.Push(pq, &State{cost: 0, r: start.Row, c: start.Col, dr: 0, dc: 1})
	lowestCost[[2]int{start.Row, start.Col}] = math.MaxInt

	for pq.Len() > 0 {
		state := heap.Pop(pq).(*State)
		cost := state.cost
		r, c := state.r, state.c
		dr, dc := state.dr, state.dc

		// fmt.Println(lowestCost)
		costFound, exisits := lowestCost[[2]int{r, c}]
		if !exisits || cost <= costFound {
			lowestCost[[2]int{r, c}] = cost
		} else {
			continue
		}
		// Check if reached destination or certain condition
		if r == end.Row && c == end.Col {
			if cost <= bestCost {
				bestCost = cost
				endStates = append(endStates, [2]int{r, c})
			}
			break
		}

		for _, direction := range []struct {
			newCost, nr, nc, ndr, ndc int
		}{
			{cost + 1, r + dr, c + dc, 0, 1},
			{cost + 1, r + dr, c + dc, 1, 0},
			{cost + 1, r + dr, c + dc, -1, 0},
			{cost + 1, r + dr, c + dc, 0, -1},
		} {
			if !isValid(grid, direction.nr, direction.nc) || grid[direction.nr][direction.nc] == true {
				continue
			}

			lowest, exists := lowestCost[[2]int{direction.nr, direction.nc}]
			if !exists || direction.newCost <= lowest {
				lowestCost[[2]int{direction.nr, direction.nc}] = direction.newCost
				heap.Push(pq, &State{cost: direction.newCost, r: direction.nr, c: direction.nc, dr: direction.ndr, dc: direction.ndc})
				backtrack[[2]int{direction.nr, direction.nc}] =
					append(backtrack[[2]int{direction.nr, direction.nc}],
						[2]int{r, c})
			}
		}
	}

	seen := make(map[[2]int]bool)
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

	fmt.Println("Safe sit positions:", seenIndexes)
	fmt.Println("Safe sit positions:", len(seenIndexes))
	return seenIndexes
}

// To found shortest possible path -

func main() {
	maze, _, _, allBytes := readMaze("input.txt")
	// Calculate the minimal score path using Dijkstra's algorithm
	// dijkstra(maze, start, end)
	// fmt.Println(maze)
	seen, _ := utils.BfsShortestPath(maze, utils.Point{0, 0}, utils.Point{len(maze) - 1, len(maze) - 1})

	seenIndexes := make(map[[2]int]bool)
	for _, ele := range seen {
		seenIndexes[[2]int{ele.X, ele.Y}] = true
	}
	fmt.Println("seendindexes ", seenIndexes, len(seenIndexes))

	printMaze(maze, seenIndexes)
	// fmt.Println("allBytes ", allBytes)
	for i, indexes := range allBytes {
		if !seenIndexes[[2]int{indexes.Row, indexes.Col}] {
			continue
		}

		maze[indexes.Row][indexes.Col] = true
		isPossible := utils.Dijkstra(maze, utils.Point{0, 0}, utils.Point{len(maze) - 1, len(maze) - 1})
		if isPossible == -1 {
			fmt.Println(indexes.Col, indexes.Row, "index is ", i+STARTFROM)
			break
		}
	}

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

func printMaze(maze [][]bool, seenIndexes map[[2]int]bool) {
	rows := len(maze)
	cols := len(maze[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if maze[i][j] {
				// Obstacle
				fmt.Print("#")
			} else if seenIndexes[[2]int{i, j}] {
				// Visited point
				fmt.Print("O")
			} else {
				// Unvisited open point
				fmt.Print(".")
			}
		}
		fmt.Println() // Newline after each row
	}
}

// func calculateUpdateLowestCost(lowestCost map[Cost]int, current Position) bool {
// 	p := Cost{current.Row, current.Col, current.direction}
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
