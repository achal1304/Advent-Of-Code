package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/achal1304/Advent-Of-Code/utils"
)

const searchWord = "XMAS"
const searchInXWord = "MAS"
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
	fmt.Println(searchMASWordinX(wordarray, len(wordarray)))

}

func searchMASWordinX(wordArray []string, size int) int {
	totCount := 0
	dictMAS := make(map[string]int)
	var wg sync.WaitGroup
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if string(wordArray[row][col]) == string(searchInXWord[0]) {
				fmt.Println("************ A FOUNDDDD **************", col, row)
				wg.Add(4) // Start 4 goroutines for different directions

				// Parallelized search in four directions
				go func() {
					defer wg.Done()
					countLeftTop := searchXMAS(wordArray, 1, row-1, col-1, 1, LeftTop, searchInXWord)
					if countLeftTop == 1 {
						keyWord := strings.Join([]string{fmt.Sprint(row - 1), "-", fmt.Sprint(col - 1)}, "")
						fmt.Println("keyword update for finding MAS ", keyWord)
						utils.UpdateDict(dictMAS, keyWord)
					}
				}()

				go func() {
					defer wg.Done()
					countLeftBottom := searchXMAS(wordArray, 1, row+1, col-1, 1, LeftBottm, searchInXWord)
					if countLeftBottom == 1 {
						keyWord := strings.Join([]string{fmt.Sprint(row + 1), "-", fmt.Sprint(col - 1)}, "")
						fmt.Println("keyword update for finding MAS ", keyWord)
						utils.UpdateDict(dictMAS, keyWord)
					}
				}()

				go func() {
					defer wg.Done()
					countRightTop := searchXMAS(wordArray, 1, row-1, col+1, 1, RightTop, searchInXWord)
					if countRightTop == 1 {
						keyWord := strings.Join([]string{fmt.Sprint(row - 1), "-", fmt.Sprint(col + 1)}, "")
						fmt.Println("keyword update for finding MAS ", keyWord)
						utils.UpdateDict(dictMAS, keyWord)
					}
				}()

				go func() {
					defer wg.Done()
					countRightBottom := searchXMAS(wordArray, 1, row+1, col+1, 1, RightBottom, searchInXWord)
					if countRightBottom == 1 {
						keyWord := strings.Join([]string{fmt.Sprint(row + 1), "-", fmt.Sprint(col + 1)}, "")
						fmt.Println("keyword update for finding MAS ", keyWord)
						utils.UpdateDict(dictMAS, keyWord)
					}
				}()
			}
		}
	}
	wg.Wait()

	fmt.Println("finaldict", dictMAS)
	for _, v := range dictMAS {
		if v == 2 {
			totCount += 1
		}
	}
	return totCount
}

func searchXMASWord(wordArray []string, size int) int {
	totCount := 0
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if string(wordArray[row][col]) == string(searchWord[0]) {
				fmt.Println("************ X FOUNDDDD **************", col, row)
				totCount += searchXMAS(wordArray, 1, row, col-1, 1, Left, searchWord)
				fmt.Println("totalcount from left ", totCount)
				totCount += searchXMAS(wordArray, 1, row, col+1, 1, Right, searchWord)
				fmt.Println("totalcount from right ", totCount)
				totCount += searchXMAS(wordArray, 1, row-1, col, 1, Top, searchWord)
				fmt.Println("totalcount from top ", totCount)
				totCount += searchXMAS(wordArray, 1, row+1, col, 1, Bottom, searchWord)
				fmt.Println("totalcount from vottom ", totCount)
				totCount += searchXMAS(wordArray, 1, row-1, col-1, 1, LeftTop, searchWord)
				fmt.Println("totalcount from left top ", totCount)
				totCount += searchXMAS(wordArray, 1, row+1, col-1, 1, LeftBottm, searchWord)
				fmt.Println("totalcount from left bottom ", totCount)
				totCount += searchXMAS(wordArray, 1, row-1, col+1, 1, RightTop, searchWord)
				fmt.Println("totalcount from righttop ", totCount)
				totCount += searchXMAS(wordArray, 1, row+1, col+1, 1, RightBottom, searchWord)
				fmt.Println("totalcount from rightbottom ", totCount)
			}
		}
	}
	return totCount
}

func searchXMAS(wordArray []string, searchIndex int, row int, col int, currCount int, direction int, searchTarget string) int {
	totCount := 0
	if col < 0 || row < 0 || col+1 > len(wordArray) || row+1 > len(wordArray) {
		return totCount
	}
	if string(wordArray[row][col]) == string(searchTarget[searchIndex]) {
		fmt.Println("found word ", string(searchTarget[searchIndex]), " at x ", col, " at y ", row)
		fmt.Println("count  ", currCount)
		searchIndex += 1
		currCount += 1
		if currCount == len(searchTarget) {
			return 1
		}
		// Left
		if col-1 >= 0 && direction == Left {
			count := searchXMAS(wordArray, searchIndex, row, col-1, currCount, Left, searchTarget)
			fmt.Println("Left returned totalcount is ", count, " ", col, " ", row)
			totCount += count
		}

		// Right
		if col+1 < len(wordArray[0]) && direction == Right {
			count := searchXMAS(wordArray, searchIndex, row, col+1, currCount, Right, searchTarget)
			fmt.Println("Right returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count

		}

		// Top:
		if row-1 >= 0 && direction == Top {
			count := searchXMAS(wordArray, searchIndex, row-1, col, currCount, Top, searchTarget)
			fmt.Println("Top returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count

		}

		// Bottom:
		if row+1 < len(wordArray) && direction == Bottom {
			count := searchXMAS(wordArray, searchIndex, row+1, col, currCount, Bottom, searchTarget)
			fmt.Println("Bottom returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}

		// Top-left diagonal:
		if row-1 >= 0 && col-1 >= 0 && direction == LeftTop {
			count := searchXMAS(wordArray, searchIndex, row-1, col-1, currCount, LeftTop, searchTarget)
			fmt.Println("LeftTop returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}

		// Top-right diagonal
		if row-1 >= 0 && col+1 < len(wordArray[0]) && direction == RightTop {
			count := searchXMAS(wordArray, searchIndex, row-1, col+1, currCount, RightTop, searchTarget)
			fmt.Println("RightTop returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}

		// Bottom-left diagonal
		if row+1 < len(wordArray) && col-1 >= 0 && direction == LeftBottm {
			count := searchXMAS(wordArray, searchIndex, row+1, col-1, currCount, LeftBottm, searchTarget)
			fmt.Println(" LeftBottm returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}
		// Bottom right diagonal
		if row+1 < len(wordArray) && col+1 < len(wordArray[0]) && direction == RightBottom {
			count := searchXMAS(wordArray, searchIndex, row+1, col+1, currCount, RightBottom, searchTarget)
			fmt.Println("RightBottom returned totalcount is ", count, " ", col, " ", row)
			// if count == 4 {
			// 	totCount += 1
			// }
			totCount += count
		}
	}
	return totCount
}
