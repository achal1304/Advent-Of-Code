package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/achal1304/Advent-Of-Code/utils"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dict := make(map[int][]int)
	predecessorCount := make(map[int]int)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		parts := strings.Split(text, "|")
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("unable to parse parts ", err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("unable to parse parts ", err)
		}

		predecessorCount[num2] += 1
		if _, ok := predecessorCount[num1]; !ok {
			predecessorCount[num1] = 0
		}

		utils.UpdateListDict(dict, num2, num1)
	}

	var printers [][]int
	for scanner.Scan() {
		var printer []int
		text := scanner.Text()
		if text == "" {
			break
		}
		parts := strings.Split(text, ",")
		for _, ele := range parts {
			num, err := strconv.Atoi(ele)
			if err != nil {
				log.Fatal("unable to parse parts ", err)
			}
			printer = append(printer, num)
		}
		printers = append(printers, printer)
	}

	for k, _ := range dict {
		sort.Ints(dict[k])
	}
	correctCount, incorrectPages := validatePages(printers, dict, predecessorCount)
	fmt.Println("correctcount", correctCount)
	fmt.Println("incorrectPages", incorrectPages)
	totalIncorrect := processIncorrectPages(incorrectPages, dict)
	fmt.Println("totalinccorect", totalIncorrect)

}

func validatePages(printerPages [][]int, graph map[int][]int, predecessorMap map[int]int) (int, [][]int) {
	indexTotalCount := 0
	incorrectList := [][]int{}
	for _, printerPage := range printerPages {
		receivedCount := validateEachPage(printerPage, graph)
		if receivedCount == 0 {
			indexTotalCount += receivedCount
			incorrectList = append(incorrectList, printerPage)
		}
	}
	return indexTotalCount, incorrectList
}

func validateEachPage(printerPage []int, graph map[int][]int) int {
	isEveryIterationValid := true
	for i := 0; i < len(printerPage)-1; i++ {
		searchedElements := graph[printerPage[i]]
		for j := i + 1; j < len(printerPage); j++ {
			isPresentInList := utils.BinarySearch(searchedElements, printerPage[j])
			if isPresentInList > -1 {
				isEveryIterationValid = false
				break
			}
		}
	}
	if isEveryIterationValid {
		return printerPage[len(printerPage)/2]
	}
	return 0
}

func processIncorrectPages(printerPages [][]int, graph map[int][]int) int {
	indexTotalCount := 0
	for _, printerPage := range printerPages {
		indexTotalCount += fixAndValidateIncorrectPage(printerPage, graph)
	}
	return indexTotalCount
}

func fixAndValidateIncorrectPage(printerPage []int, graph map[int][]int) int {
	for i := 0; i < len(printerPage)-1; i++ {
		for j := i + 1; j < len(printerPage); j++ {
			searchedElements := graph[printerPage[i]]
			presentIndex := utils.BinarySearch(searchedElements, printerPage[j])
			if presentIndex > -1 {
				currentIndexList := []int{printerPage[j]}
				printerPage = append(printerPage[:j], printerPage[j+1:]...)
				printerPage = append(printerPage[:i], append(currentIndexList, printerPage[i:]...)...)
				printerPage = append(printerPage)
			}
		}
	}

	fmt.Println("** final updated list is **", printerPage)
	return printerPage[len(printerPage)/2]
}
