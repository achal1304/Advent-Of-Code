package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		utils.UpdateListDict(dict, num1, num2)
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

	// fmt.Println(dict)
	// fmt.Println(printers)

	// Copy the inDegree map to avoid modifying the original map

	fmt.Println(dict)
	for k, v := range dict {
		fmt.Println("key ", k, " val ", v)
	}
	fmt.Println(validatePages(printers, dict, predecessorCount))
}

func validateEachPage(printerPage []int, graph map[int][]int, predecessorMap map[int]int) int {
	// count := 0
	queue := []int{}
	for page, predecessorCount := range predecessorMap {
		if predecessorCount == 0 {
			queue = append(queue, page)
		}
	}

	processedPages := 0

	// fmt.Println("searching in ", printerPage)
	for len(queue) > 0 && processedPages < len(printerPage) {
		fmt.Println("queue ", queue)
		// fmt.Println("currrent index", processedPages)
		checkPage := queue[0]
		queue = queue[1:]

		if printerPage[processedPages] == checkPage {
			processedPages += 1
		}

		for _, successor := range graph[checkPage] {
			predecessorMap[successor]--
			if predecessorMap[successor] == 0 {
				queue = append(queue, successor)
			}
		}
	}
	// fmt.Println("processedpages and len", processedPages, " ", len(printerPage))
	if processedPages == len(printerPage) {
		return printerPage[len(printerPage)/2]
	}
	return 0
}

func validatePages(printerPages [][]int, graph map[int][]int, predecessorMap map[int]int) int {
	indexTotalCount := 0
	for _, printerPage := range printerPages {
		predecessorMapCopy := make(map[int]int)
		for k, v := range predecessorMap {
			predecessorMapCopy[k] = v
		}
		printerPagesCopy := make(map[int][]int)
		for k, v := range graph {
			printerPagesCopy[k] = v
		}
		indexTotalCount += validateEachPage(printerPage, printerPagesCopy, predecessorMapCopy)
	}
	return indexTotalCount
}
