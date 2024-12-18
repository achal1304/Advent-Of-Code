package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Item struct {
	point Point
	cost  int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
}

func Dijkstra(grid [][]bool, start Point, end Point) int {
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	rows := len(grid)
	cols := len(grid[0])
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{point: start, cost: 0})

	dist := make(map[Point]int)
	dist[start] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		if current.point == end {
			return current.cost
		}

		for _, dir := range directions {
			nx, ny := current.point.x+dir.x, current.point.y+dir.y
			if isValid(nx, ny, rows, cols) && !grid[nx][ny] {
				neighbor := Point{nx, ny}
				newDist := current.cost + 1
				if oldDist, ok := dist[neighbor]; !ok || newDist < oldDist {
					dist[neighbor] = newDist
					heap.Push(pq, &Item{point: neighbor, cost: newDist})
				}
			}
		}
	}

	return -1 // unreachable
}

func main() {
	file, err := os.Open("input.txt")
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
	for scanner.Scan() {
		iteration++
		if iteration > 1024 {
			break
		}
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
		grid[x][y] = true
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// fmt.Println(grid)
	start := Point{0, 0}
	end := Point{gridSize - 1, gridSize - 1}
	steps := Dijkstra(grid, start, end)
	if steps == -1 {
		fmt.Println("The destination is unreachable.")
	} else {
		fmt.Printf("Minimum steps required: %d\n", steps)
	}
}
