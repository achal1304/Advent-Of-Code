package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("not able to open file input.txt", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	// scanner.Scan()
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)
		result += findMulInstructions(input)
	}
	fmt.Println("result ", result)
}

func findMulInstructions(operations string) int {
	result := 0
	for i := 0; i < len(operations); i++ {
		firstNum := 0
		secondNum := 0
		if string(operations[i]) == "m" {
			if i < len(operations)-7 {
				if string(operations[i+1]) == "u" && string(operations[i+2]) == "l" && string(operations[i+3]) == "(" {
					i = i + 3
					i, firstNum = detectNumberInMul(operations, i, ",")
					i, secondNum = detectNumberInMul(operations, i, ")")

					if firstNum != 0 && secondNum != 0 {
						fmt.Println("index - ", i, " next ele - ", string(operations[i]), " first num -", firstNum, " second num -", secondNum)
						result += firstNum * secondNum
						fmt.Println("result  ", result)
					}
				}
			}
		}
	}
	return result
}

func disableMulCheck(operations string, currentIndex int) int {
	if len(operations)-5 > currentIndex {
		if string(operations[currentIndex:currentIndex+5]) == "don't" {
			fmt.Println("dont detected ")
			for i := currentIndex + 5; i < len(operations)-5; i++ {
				if string(operations[i:i+5]) != "don't" && string(operations[i:i+2]) == "do" {
					fmt.Println("do detected near ", string(operations[i-1:i+3]))
					return i + 3
				}
			}
			return len(operations) - 5
		}
	}
	return currentIndex
}

func detectNumberInMul(operations string, currentIndex int, searchCharacter string) (int, int) {
	outputNum := 0
	for j := currentIndex + 1; j < currentIndex+6; j++ {
		if string(operations[j]) == " " {
			return currentIndex + 1, 0
		}
		if string(operations[j]) == searchCharacter {
			num, err := strconv.Atoi(string(operations[currentIndex+1 : j]))
			if err != nil {
				return currentIndex + 1, 0
			} else {
				outputNum = num
				currentIndex = j
				return currentIndex, outputNum
			}
		}
		if j == currentIndex+4 {
			return currentIndex + 1, 0
		}
	}
	return currentIndex + 1, 0
}
