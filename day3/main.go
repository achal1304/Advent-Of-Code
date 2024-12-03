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
outerLoops:
	for i := 0; i < len(operations); i++ {
		// if i > 500 {
		// 	break
		// }
		// fmt.Println(i, " ", string(operations[i]))
		firstNum := 0
		secondNum := 0
		if string(operations[i]) == "m" {
			if i < len(operations)-7 {
				if string(operations[i+1]) == "u" && string(operations[i+2]) == "l" && string(operations[i+3]) == "(" {
					// fmt.Println("inside mul(")
					i = i + 3
				innerLoop1:
					for j := i + 1; j < i+6; j++ {
						// iteration := j - i - 1
						if string(operations[j]) == " " {
							// fmt.Println("empty spce skipping")
							i = i + 1
							continue outerLoops
						}
						if string(operations[j]) == "," {
							num, err := strconv.Atoi(string(operations[i+1 : j]))
							// fmt.Println("number check ", string(operations[i+1:j]))
							if err != nil {
								// fmt.Println("not a number ", string(operations[i+1:j]))
								i = i + 1
								continue outerLoops
							} else {
								// fmt.Println("firstnum ", num)
								firstNum = num
								i = j
								break innerLoop1
							}
						}
						if j == i+5 {
							i = i + 1
							continue outerLoops
						}
					}

				innerLoop2:
					for j := i + 1; j < i+6; j++ {
						// fmt.Println("seconf num iteration start from i and val", j, string(operations[j]))
						// iteration := j - i - 1
						if string(operations[j]) == " " {
							i = i + 1
							continue outerLoops
						}
						if string(operations[j]) == ")" {
							num, err := strconv.Atoi(string(operations[i+1 : j]))
							if err != nil {
								// fmt.Println("not a number ", string(operations[i+1:j]))
								i = i + 1
								continue outerLoops
							} else {
								// fmt.Println("secondNum ", num)
								secondNum = num
								i = j
								break innerLoop2
							}
						}
						if j == i+4 {
							i = i + 1
							continue outerLoops
						}
					}
					if firstNum != 0 && secondNum != 0 {
						// fmt.Println("new index iteration", i, " ", string(operations[i]))
						fmt.Println("index - ", i, " next ele - ", string(operations[i]), " first num -", firstNum, " second num -", secondNum)
						result += firstNum * secondNum
					}
				}
			}

		}
	}
	return result
}
