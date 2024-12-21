package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/achal1304/Advent-Of-Code/utils"
)

type Point struct {
	x, y int
}

var DirectionalPress = map[string]Point{
	"^": Point{-1, 0},
	">": Point{0, 1},
	"<": Point{0, -1},
	"v": Point{1, 0},
}

var ReversePress = map[Point]string{
	Point{-1, 0}: "^",
	Point{0, 1}:  ">",
	Point{0, -1}: "<",
	Point{1, 0}:  "v",
}

var PasswordPress = map[string]Point{
	"1": Point{2, 0},
	"2": Point{2, 1},
	"3": Point{2, 2},
	"4": Point{1, 0},
	"5": Point{1, 1},
	"6": Point{1, 2},
	"7": Point{0, 0},
	"8": Point{0, 1},
	"9": Point{0, 2},
	"0": Point{3, 1},
	"A": Point{3, 2},
	"X": Point{3, 0},
}

var KeyPad = map[string]Point{
	"^": Point{0, 1},
	">": Point{1, 2},
	"<": Point{1, 0},
	"v": Point{1, 1},
	"A": Point{0, 2},
}

func specialCase(dX, dY int) []string {
	eleList := []string{}
	for dY > 0 {
		eleList = append(eleList, ReversePress[Point{0, 1}])
		dY -= 1
	}
	for dX > 0 {
		eleList = append(eleList, ReversePress[Point{1, 0}])
		dX -= 1
	}
	for dX < 0 {
		eleList = append(eleList, ReversePress[Point{-1, 0}])
		dX += 1
	}
	for dY < 0 {
		eleList = append(eleList, ReversePress[Point{0, -1}])
		dY++
	}
	return eleList
}

func checkForFirstRobot(inputTextStrings []string, checkX, checkY int, checkPress map[string]Point) [][]string {
	start := checkPress["A"]
	finalEleList := [][]string{}
	for _, inputText := range inputTextStrings {
		eleList := [][]string{{}, {}, {}, {}}
		for _, ele := range inputText {
			dX, dY := indexDistance(start, checkPress[string(ele)])
			specialList := specialCase(dX, dY)
			eleList[0] = append(eleList[0], specialList...)
			eleList[0] = append(eleList[0], "A")
			start = checkPress[string(ele)]
		}
		start = checkPress["A"]
		for _, ele := range inputText {
			passPoint := checkPress[string(ele)]
			dX, dY := indexDistance(start, checkPress[string(ele)])
			if (passPoint.x == checkX || passPoint.y == checkY) &&
				(start.x == checkX || start.y == checkY) {
				specialList := specialCase(dX, dY)
				eleList[1] = append(eleList[1], specialList...)
				eleList[1] = append(eleList[1], "A")
			} else {
				for dX > 0 {
					eleList[1] = append(eleList[1], ReversePress[Point{1, 0}])
					dX -= 1
				}
				for dY > 0 {
					eleList[1] = append(eleList[1], ReversePress[Point{0, 1}])
					dY -= 1
				}
				for dY < 0 {
					eleList[1] = append(eleList[1], ReversePress[Point{0, -1}])
					dY++
				}
				for dX < 0 {
					eleList[1] = append(eleList[1], ReversePress[Point{-1, 0}])
					dX += 1
				}
				eleList[1] = append(eleList[1], "A")
			}
			start = checkPress[string(ele)]
		}
		start = checkPress["A"]
		for _, ele := range inputText {
			passPoint := checkPress[string(ele)]
			dX, dY := indexDistance(start, checkPress[string(ele)])
			if (passPoint.x == checkX || passPoint.y == checkY) &&
				(start.x == checkX || start.y == checkY) {
				specialList := specialCase(dX, dY)
				eleList[2] = append(eleList[2], specialList...)
				eleList[2] = append(eleList[2], "A")
			} else {
				for dX > 0 {
					eleList[2] = append(eleList[2], ReversePress[Point{1, 0}])
					dX -= 1
				}
				for dY > 0 {
					eleList[2] = append(eleList[2], ReversePress[Point{0, 1}])
					dY -= 1
				}
				for dX < 0 {
					eleList[2] = append(eleList[2], ReversePress[Point{-1, 0}])
					dX += 1
				}
				for dY < 0 {
					eleList[2] = append(eleList[2], ReversePress[Point{0, -1}])
					dY++
				}
				eleList[2] = append(eleList[2], "A")
			}
			start = checkPress[string(ele)]
		}

		start = checkPress["A"]
		for _, ele := range inputText {
			passPoint := checkPress[string(ele)]
			dX, dY := indexDistance(start, checkPress[string(ele)])
			if (passPoint.x == checkX || passPoint.y == checkY) &&
				(start.x == checkX || start.y == checkY) {
				specialList := specialCase(dX, dY)
				eleList[3] = append(eleList[3], specialList...)
				eleList[3] = append(eleList[3], "A")
			} else {
				for dY < 0 {
					eleList[3] = append(eleList[3], ReversePress[Point{0, -1}])
					dY++
				}
				for dX < 0 {
					eleList[3] = append(eleList[3], ReversePress[Point{-1, 0}])
					dX += 1
				}
				for dY > 0 {
					eleList[3] = append(eleList[3], ReversePress[Point{0, 1}])
					dY -= 1
				}
				for dX > 0 {
					eleList[3] = append(eleList[3], ReversePress[Point{1, 0}])
					dX -= 1
				}
				eleList[3] = append(eleList[3], "A")
			}
			start = checkPress[string(ele)]
		}
		finalEleList = append(finalEleList, eleList...)
	}
	return finalEleList
}

func manhattanDistance(p1, p2 Point) int {
	return utils.AbsInt(p1.x-p2.x) + utils.AbsInt(p1.y-p2.y)
}
func indexDistance(p1, p2 Point) (int, int) {
	return p2.x - p1.x, p2.y - p1.y
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Errorf("error occured ")
		os.Exit(1)
	}
	defer file.Close()

	inputNum := []int{29, 980, 179, 456, 379}
	inputText := []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}
	ans := 0
	for i, ele := range inputText {
		ansList := checkForFirstRobot([]string{ele}, 3, 0, PasswordPress)
		ansList1 := checkUniqueLists(ansList)
		fmt.Println("Pass 1:", ansList1)

		// ansList1 = []string{"<A^A^>^AvvvA"}
		for j := 0; j < 3; j++ {
			ansList = checkForFirstRobot(ansList1, 0, 0, KeyPad)
			ansList1 = checkUniqueLists(ansList)
			fmt.Printf("Pass %d %v, Length: %d\n", j, ansList1, len(ansList1))
		}

		min := getMinListLength(ansList1)
		fmt.Println("Minimum length is", min)
		ans += min * inputNum[i]
	}
	fmt.Println(ans)
}

func getMinListLength(input []string) int {
	minLen := 999999999
	for _, ele := range input {
		if len(ele) < minLen {
			minLen = len(ele)
		}
	}
	return minLen
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
