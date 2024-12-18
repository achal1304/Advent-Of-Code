package utils

import (
	"container/heap"
	"container/list"
)

type Point struct {
	X, Y int
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
			nx, ny := current.point.X+dir.X, current.point.Y+dir.Y
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

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
}

func isBfsValid(x, y, rows, cols int, grid [][]bool, visited [][]bool) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols && !grid[x][y] && !visited[x][y]
}

func BfsShortestPath(grid [][]bool, start Point, end Point) ([]Point, int) {
	rows := len(grid)
	cols := len(grid[0])
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	queue := list.New()
	queue.PushBack([]Point{start})

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	visited[start.X][start.Y] = true

	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]Point)
		current := path[len(path)-1]

		if current == end {
			return path, len(path) - 1
		}

		for _, dir := range directions {
			nx, ny := current.X+dir.X, current.Y+dir.Y
			if isBfsValid(nx, ny, rows, cols, grid, visited) {
				visited[nx][ny] = true
				newPath := append([]Point{}, path...)
				newPath = append(newPath, Point{nx, ny})
				queue.PushBack(newPath)
			}
		}
	}

	return nil, -1 // No path found
}
