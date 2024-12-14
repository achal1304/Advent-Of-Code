package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// State represents a position in the 2D grid (X, Y) and the total cost incurred to reach that position
type State struct {
	x, y, cost, aPresses, bPresses int
}

// PriorityQueue implements heap.Interface and holds States
type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// Prioritize the state with the lowest cost
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// The claw machine problem using a priority queue to minimize the cost
func solveClawMachine(XA, YA, XB, YB, XP, YP int) int {
	// Min-heap (priority queue)
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Start at position (0, 0) with 0 cost and 0 button presses
	heap.Push(pq, &State{x: 0, y: 0, cost: 0, aPresses: 0, bPresses: 0})

	// A set to track visited positions to avoid reprocessing
	visited := map[[2]int]bool{}

	// Directions: Button A (XA, YA) and Button B (XB, YB)
	directions := []struct {
		dx, dy, cost, buttonType int
	}{
		{XA, YA, 3, 1}, // Button A
		{XB, YB, 1, 2}, // Button B
	}

	// While the queue is not empty
	for pq.Len() > 0 {
		// Pop the state with the smallest cost
		currentState := heap.Pop(pq).(*State)

		// If we've reached the prize, return the cost
		if currentState.x == XP && currentState.y == YP {
			return currentState.cost
		}

		// Mark the current position as visited
		visitedKey := [2]int{currentState.x, currentState.y}
		if visited[visitedKey] {
			continue
		}
		visited[visitedKey] = true

		// Explore both Button A and Button B moves
		for _, dir := range directions {
			// Calculate new position
			newX := currentState.x + dir.dx
			newY := currentState.y + dir.dy
			newCost := currentState.cost + dir.cost

			// Check if the button press limit is exceeded
			newAPresses := currentState.aPresses
			newBPresses := currentState.bPresses
			if dir.buttonType == 1 {
				newAPresses++
			} else {
				newBPresses++
			}

			// If the new position has not been visited, add it to the priority queue
			if !visited[[2]int{newX, newY}] {
				newState := &State{
					x:        newX,
					y:        newY,
					cost:     newCost,
					aPresses: newAPresses,
					bPresses: newBPresses,
				}
				heap.Push(pq, newState)
			}
		}
	}

	// If no solution found (should not happen in theory)
	return -1
}

func processClawMachineInput(filename string) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	totCount1 := 0
	totCount2 := 0

	iteration := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		iteration += 1
		// Read Button A line (e.g., "Button A: X+94, Y+34")
		buttonALine := scanner.Text()
		if buttonALine == "" {
			break
		}
		scanner.Scan() // Read Button B line
		buttonBLine := scanner.Text()
		scanner.Scan() // Read Prize line
		prizeLine := scanner.Text()

		// Parse Button A and B values
		reButtonA := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
		// Parse Button A
		matchA := reButtonA.FindStringSubmatch(buttonALine)
		XA, _ := strconv.Atoi(matchA[1])
		YA, _ := strconv.Atoi(matchA[2])

		// Correct regex to parse Button B: X+22, Y+67
		reButtonB := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
		// Parse Button B
		matchB := reButtonB.FindStringSubmatch(buttonBLine)
		XB, _ := strconv.Atoi(matchB[1])
		YB, _ := strconv.Atoi(matchB[2])

		// Correct regex to parse Prize: X=8400, Y=5400
		rePrize := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
		// Parse Prize location
		prizeMatch := rePrize.FindStringSubmatch(prizeLine)
		XP, _ := strconv.Atoi(prizeMatch[1])
		YP, _ := strconv.Atoi(prizeMatch[2])

		// Part 1
		// result := solveClawMachine(XA, YA, XB, YB, XP, YP)
		// fmt.Println("result is ", result)

		result2 := solveClawMachinePart2(XA, YA, XB, YB, XP+10000000000000, YP+10000000000000)

		// Part 1
		// if result != -1 {
		// 	totCount1 += result
		// }

		if result2 != 0 {
			fmt.Println("winning prize on machine ", iteration)
			totCount2 += result2
		}
		scanner.Scan()
	}
	fmt.Println(totCount1)
	fmt.Println(totCount2)
}

func solveClawMachinePart2(XA, YA, XB, YB, XP, YP int) int {
	m := ((YP * XA) - (XP * YA)) / ((YB * XA) - (XB * YA))
	n := ((XP) - (m * XB)) / XA

	fmt.Println("m and n for machines ", m, n)

	if XA*n+XB*m == XP && YA*n+YB*m == YP {
		return 3*n + m
	}
	return 0
}

// Main function to start the program
func main() {
	// Specify the input file name
	filename := "input.txt"
	processClawMachineInput(filename)
}
