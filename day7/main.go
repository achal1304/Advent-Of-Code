package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error in reading file ", file)
	}
	defer file.Close()

	result := processInput(*file)
	fmt.Println(result)
}

func checkValidCombinations(testValue int, numbers []int) int {
	fmt.Println("new num and queue ", testValue, " ", numbers)
	queue := []int{numbers[0]}
	allCalculations := make(map[int]int)
	for i := 1; i < len(numbers); i++ {
		for _, ele := range queue {
			calculation1 := ele * numbers[i]
			if calculation1 <= testValue {
				queue = append(queue, calculation1)
			}

			calculation2 := ele + numbers[i]
			if calculation2 <= testValue {
				queue = append(queue, calculation2)
			}

			calculation3String := fmt.Sprint(ele) + fmt.Sprint(numbers[i])
			calculation3, err := strconv.Atoi(calculation3String)
			if err != nil {
				log.Fatalln("invalid calculation 3", calculation3String)
				return 0
			}
			if calculation3 <= testValue {
				queue = append(queue, calculation3)
			}
			if i == len(numbers)-1 {
				if calculation1 == testValue || calculation2 == testValue || calculation3 == testValue {
					fmt.Println("** found valid operations **", numbers)
					return testValue
				}
				allCalculations[calculation1] = 1
				allCalculations[calculation2] = 1
				allCalculations[calculation3] = 1
			}
		}
		queue = queue[1:]
	}
	return 0
}

func checkCombinationsInGoroutines(allNumbers [][]int, allTargets []int) int {
	var wg sync.WaitGroup
	totalCount := 0
	ch := make(chan int, len(allTargets))
	for i, ele := range allNumbers {
		j := i
		go func(numbers []int, testValue int) {
			defer wg.Done()
			wg.Add(1)
			count := checkValidCombinations(testValue, numbers)
			ch <- count
		}(ele, allTargets[j])
	}

	wg.Wait()
	close(ch)

	for ans := range ch {
		totalCount += ans
	}
	return totalCount
}

func processInput(file os.File) int {
	validEquationsCount := 0

	scanner := bufio.NewScanner(&file)
	allNumbers := [][]int{}
	allTargets := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		testValueStr := strings.TrimSpace(parts[0])
		numbersStr := strings.TrimSpace(parts[1])

		testValue, _ := strconv.Atoi(testValueStr)

		numbersStrList := strings.Fields(numbersStr)
		numbers := make([]int, len(numbersStrList))

		for i, numStr := range numbersStrList {
			num, _ := strconv.Atoi(numStr)
			numbers[i] = num
		}

		allNumbers = append(allNumbers, numbers)
		allTargets = append(allTargets, testValue)
	}

	validEquationsCount += checkCombinationsInGoroutines(allNumbers, allTargets)

	return validEquationsCount
}
