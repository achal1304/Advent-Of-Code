package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

const MaxDepth = 24

var ReversePress = map[Point]string{
	{-1, 0}: "^",
	{0, 1}:  ">",
	{0, -1}: "<",
	{1, 0}:  "v",
}

var PasswordPress = map[string]Point{
	"1": {2, 0},
	"2": {2, 1},
	"3": {2, 2},
	"4": {1, 0},
	"5": {1, 1},
	"6": {1, 2},
	"7": {0, 0},
	"8": {0, 1},
	"9": {0, 2},
	"0": {3, 1},
	"A": {3, 2},
	"X": {3, 0},
}

var KeyPad = map[string]Point{
	"^": {0, 1},
	">": {1, 2},
	"<": {1, 0},
	"v": {1, 1},
	"A": {0, 2},
}

type Memo struct {
	cache map[string]int
}

func NewMemo() *Memo {
	return &Memo{cache: make(map[string]int)}
}

func (m *Memo) Get(key string) (int, bool) {
	val, exists := m.cache[key]
	return val, exists
}

func (m *Memo) Set(key string, value int) {
	m.cache[key] = value
}

func computeLength(x, y string, startInd, depth int, memo *Memo) int {
	startKey := "A"
	if startInd != 0 {
		startKey = x
	}

	if depth == 1 {
		sequences := checkForNextRobot(KeyPad[startKey], []string{x + y}, 0, 0, KeyPad)[0]
		return len(sequences) - 1
	}

	key := fmt.Sprintf("%s-%s-%s-%d", startKey, x, y, depth)
	if val, found := memo.Get(key); found {
		// fmt.Println("********************* found in memo *************************", key, val)
		return val
	}

	optimal := int(^uint(0) >> 1) // Max int value
	sequences := checkForNextRobot(KeyPad[startKey], []string{x + y}, 0, 0, KeyPad)

	for _, seq := range sequences {
		length := 0
		for i := 0; i < len(seq)-1; i++ {
			length += computeLength(string(seq[i]), string(seq[i+1]), i, depth-1, memo)
		}
		if length < optimal {
			optimal = length
		}
	}

	memo.Set(key, optimal)
	return optimal
}

func specialCase(dX, dY int) []string {
	path := []string{}
	for dY > 0 {
		path = append(path, ReversePress[Point{0, 1}])
		dY--
	}
	for dX > 0 {
		path = append(path, ReversePress[Point{1, 0}])
		dX--
	}
	for dX < 0 {
		path = append(path, ReversePress[Point{-1, 0}])
		dX++
	}
	for dY < 0 {
		path = append(path, ReversePress[Point{0, -1}])
		dY++
	}
	return path
}

func indexDistance(p1, p2 Point) (int, int) {
	return p2.x - p1.x, p2.y - p1.y
}

func checkForRobot(inputTextStrings []string, startPos Point, checkX, checkY int, checkPress map[string]Point, phases int) [][]string {
	finalEleList := [][]string{}

	for _, inputText := range inputTextStrings {
		eleList := make([][]string, phases)

		for phase := 0; phase < phases; phase++ {
			start := startPos
			for _, ele := range inputText {
				passPoint := checkPress[string(ele)]
				dX, dY := indexDistance(start, passPoint)

				if (passPoint.x == checkX || passPoint.y == checkY) &&
					(start.x == checkX || start.y == checkY) {
					eleList[phase] = append(eleList[phase], specialCase(dX, dY)...) // Direct path
				} else {
					eleList[phase] = append(eleList[phase], generatePath(dX, dY, phase)...) // Alternate path
				}

				eleList[phase] = append(eleList[phase], "A")
				start = passPoint
			}
		}
		finalEleList = append(finalEleList, eleList...)
	}

	return finalEleList
}

func checkForNextRobot(startPos Point, inputTextStrings []string, checkX, checkY int, checkPress map[string]Point) []string {
	paths := checkForRobot(inputTextStrings, startPos, checkX, checkY, checkPress, 4) // Reuse checkForRobot logic
	stringPaths := []string{}
	for _, ele := range paths {
		stringPaths = append(stringPaths, strings.Join(ele, ""))
	}
	return stringPaths
}

func checkForFirstAndNextRobot(inputTextStrings []string, checkX, checkY int, checkPress map[string]Point) [][]string {
	startPos := checkPress["A"]
	return checkForRobot(inputTextStrings, startPos, checkX, checkY, checkPress, 4) // 4 phases
}

func generatePath(dX, dY, phase int) []string {
	path := []string{}

	order := [][]Point{
		{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}, // Phase 0
		{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}, // Phase 1
		{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}, // Phase 2
		{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}, // Phase 3
	}

	for _, dir := range order[phase] {
		for (dX > 0 && dir == (Point{1, 0})) ||
			(dX < 0 && dir == (Point{-1, 0})) ||
			(dY > 0 && dir == (Point{0, 1})) ||
			(dY < 0 && dir == (Point{0, -1})) {
			path = append(path, ReversePress[dir])
			if dir.x != 0 {
				dX -= dir.x
			} else {
				dY -= dir.y
			}
		}
	}

	return path
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Errorf("error occurred")
		os.Exit(1)
	}
	defer file.Close()

	inputNum := []int{382, 463, 935, 279, 480}
	inputText := []string{
		"382A",
		"463A",
		"935A",
		"279A",
		"480A",
	}
	ans := 0
	for i, ele := range inputText {
		ansList := checkForFirstAndNextRobot([]string{ele}, 3, 0, PasswordPress)
		ansList1 := checkUniqueLists(ansList)

		memo := NewMemo()
		totLen := 9999999999999
		for j := 0; j < len(ansList1); j++ {
			minLen := 0
			for k := 0; k < len(ansList1[j])-1; k++ {
				optimal := computeLength(string(ansList1[j][k]), string(ansList1[j][k+1]), k, MaxDepth, memo)
				minLen += optimal
			}
			if minLen < totLen {
				totLen = minLen
			}
		}
		fmt.Println("TotalLength for", ele, "After 24 iterations is", "length is", totLen+1)
		ans += (totLen + 1) * inputNum[i]
	}
	fmt.Println(ans)
}

func getMinListLength2(input [][]string) int {
	minLen := 999999999
	for _, ele := range input {
		if len(ele) < minLen {
			minLen = len(ele)
		}
	}
	return minLen
}

func checkUniqueLists(input [][]string) []string {
	uniqueMap := make(map[string]bool)
	var distinct [][]string

	minLength := getMinListLength2(input)

	for _, sequence := range input {
		key := fmt.Sprint(sequence)
		if len(sequence) > minLength {
			continue
		}
		if !uniqueMap[key] {
			uniqueMap[key] = true
			distinct = append(distinct, sequence)
		}
	}
	finallist := []string{}
	for _, ele := range distinct {
		finallist = append(finallist, strings.Join(ele, ""))
	}
	return finallist
}
