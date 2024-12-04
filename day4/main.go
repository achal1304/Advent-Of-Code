package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const searchWord = "XMAS"
const (
	Left = iota
	Right
	Top
	Bottom
	LeftBottm
	RightBottom
	LeftTop
	RightTop
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordarray := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		wordarray = append(wordarray, line)
	}

	fmt.Println(wordarray)
	fmt.Println("x  length ", len(wordarray[0]))
	fmt.Println("y  length ", len(wordarray))
	// fmt.Println(searchXMASWord(wordarray, len(wordarray)))
	fmt.Println(searchXMASWord(wordarray, len(wordarray)))
}

func searchXMASWord(wordArray []string, size int) int {
	totCount := 0
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if string(wordArray[row][col]) == string(searchWord[0]) {
				fmt.Println("************ X FOUNDDDD **************", col, row)
				totCount += searchXMAS(wordArray, 1, row, col-1, 1, Left)
				fmt.Println("totalcount from left ", totCount)
				totCount += searchXMAS(wordArray, 1, row, col+1, 1, Right)
				fmt.Println("totalcount from right ", totCount)
				totCount += searchXMAS(wordArray, 1, row-1, col, 1, Top)
				fmt.Println("totalcount from top ", totCount)
				totCount += searchXMAS(wordArray, 1, row+1, col, 1, Bottom)
				fmt.Println("totalcount from vottom ", totCount)
				totCount += searchXMAS(wordArray, 1, row-1, col-1, 1, LeftTop)
				fmt.Println("totalcount from left top ", totCount)
				totCount += searchXMAS(wordArray, 1, row+1, col-1, 1, LeftBottm)
				fmt.Println("totalcount from left bottom ", totCount)
				totCount += searchXMAS(wordArray, 1, row-1, col+1, 1, RightTop)
				fmt.Println("totalcount from righttop ", totCount)
				totCount += searchXMAS(wordArray, 1, row+1, col+1, 1, RightBottom)
				fmt.Println("totalcount from rightbottom ", totCount)
			}
		}
	}
	return totCount
}

func searchXMAS(wordArray []string, searchIndex int, row int, col int, currCount int, direction int) int {
	totCount := 0
	if col < 0 || row < 0 || col+1 > len(wordArray) || row+1 > len(wordArray) {
		// fmt.Println("returning from here ", string(searchWord[searchIndex]), " at x ", col, " at y ", row)
		return totCount
	}
	// fmt.Println("checking ", string(wordArray[row][col]), " ", string(searchWord[searchIndex]))
	if string(wordArray[row][col]) == string(searchWord[searchIndex]) {
		fmt.Println("found word ", string(searchWord[searchIndex]), " at x ", col, " at y ", row)
		fmt.Println("count  ", currCount)
		searchIndex += 1
		currCount += 1
		if currCount == len(searchWord) {
			return 1
		}
		// Left
		if col-1 >= 0 && direction == Left {
			count := searchXMAS(wordArray, searchIndex, row, col-1, currCount, Left)
			fmt.Println("Left returned totalcount is ", count, " ", col, " ", row)
			totCount += count
		}

		// Right
		if col+1 < len(wordArray[0]) && direction == Right {
			count := searchXMAS(wordArray, searchIndex, row, col+1, currCount, Right)
			fmt.Println("Right returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count

		}

		// Top:
		if row-1 >= 0 && direction == Top {
			count := searchXMAS(wordArray, searchIndex, row-1, col, currCount, Top)
			fmt.Println("Top returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count

		}

		// Bottom:
		if row+1 < len(wordArray) && direction == Bottom {
			count := searchXMAS(wordArray, searchIndex, row+1, col, currCount, Bottom)
			fmt.Println("Bottom returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}

		// Top-left diagonal:
		if row-1 >= 0 && col-1 >= 0 && direction == LeftTop {
			count := searchXMAS(wordArray, searchIndex, row-1, col-1, currCount, LeftTop)
			fmt.Println("LeftTop returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}

		// Top-right diagonal
		if row-1 >= 0 && col+1 < len(wordArray[0]) && direction == RightTop {
			count := searchXMAS(wordArray, searchIndex, row-1, col+1, currCount, RightTop)
			fmt.Println("RightTop returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}

		// Bottom-left diagonal
		if row+1 < len(wordArray) && col-1 >= 0 && direction == LeftBottm {
			count := searchXMAS(wordArray, searchIndex, row+1, col-1, currCount, LeftBottm)
			fmt.Println(" LeftBottm returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}
		// Bottom right diagonal
		if row+1 < len(wordArray) && col+1 < len(wordArray[0]) && direction == RightBottom {
			count := searchXMAS(wordArray, searchIndex, row+1, col+1, currCount, RightBottom)
			fmt.Println("RightBottom returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}
	}
	return totCount
}
