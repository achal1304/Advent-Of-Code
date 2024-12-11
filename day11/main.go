package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func blink(stones []string) []string {
	var newStones []string
	for _, stone := range stones {
		stoneInt, err := strconv.Atoi(stone)
		if err != nil {
			fmt.Errorf("error converting stone to int64 %v ", err)
		}
		if stoneInt == 0 {
			newStones = append(newStones, "1")
		} else {
			if len(stone)%2 == 0 {
				mid := len(stone) / 2
				left := stone[:mid]
				_, err := strconv.Atoi(left)
				if err != nil {
					fmt.Errorf("error converting stone to int64 %v ", err)
				}
				right := stone[mid:]
				_, err = strconv.Atoi(right)
				if err != nil {
					fmt.Errorf("error converting stone to int64 %v ", err)
				}
				newStones = append(newStones, left, right)
			} else {
				newStone := strconv.Itoa(stoneInt * 2024)
				newStones = append(newStones, newStone)
			}
		}
	}
	return newStones
}

var memo = make(map[string]int64)

func transformWithMemo(stone string, iterations int) int64 {
	var result int64
	stoneInt, err := strconv.Atoi(stone)
	if err != nil {
		fmt.Errorf("error converting stone to int64 %v ", err)
	}
	stone = fmt.Sprint(stoneInt)
	if result, exists := memo[fmt.Sprint(stoneInt)+"_"+strconv.Itoa(iterations)]; exists {
		return result
	}

	if iterations == 0 {
		return 1
	}

	if stoneInt == 0 {
		result = transformWithMemo("1", iterations-1)
	} else {
		numDigits := len(stone)

		if numDigits%2 == 0 {
			mid := len(stone) / 2
			left := stone[:mid]
			_, err := strconv.Atoi(left)
			if err != nil {
				fmt.Errorf("error converting stone to int64 %v ", err)
			}
			right := stone[mid:]
			_, err = strconv.Atoi(right)
			if err != nil {
				fmt.Errorf("error converting stone to int64 %v ", err)
			}
			leftAns := transformWithMemo(left, iterations-1)
			rightAns := transformWithMemo(right, iterations-1)
			result = +leftAns + rightAns
		} else {
			newStone := strconv.Itoa(stoneInt * 2024)
			result = transformWithMemo(newStone, iterations-1)
		}
	}

	// Memoize the result for future reference
	memo[fmt.Sprint(stoneInt)+"_"+strconv.Itoa(iterations)] = result
	return result
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	initialStones := strings.Fields(string(data))

	// part 1
	// stones := initialStones
	// for i := 0; i < 75; i++ {
	// 	stones = blink(stones)
	// 	// fmt.Println("new stones ", stones)
	// }

	// part 2
	totalStones := int64(0)
	for _, stone := range initialStones {
		totalStones += transformWithMemo(stone, 75)
	}

	// fmt.Println("Number of stones after 25 blinks:", len(stones))
	fmt.Println("Number of stones after 75 blinks:", totalStones)
}
