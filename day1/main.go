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

// Historian Hysteria
func main() {

	leftPath := []int{}
	rightPath := []int{}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("unable to parse parts ", err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("unable to parse parts ", err)
		}

		leftPath = append(leftPath, num1)
		rightPath = append(rightPath, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("left", leftPath)
	// fmt.Println("right", rightPath)

	part2leftpath := make([]int, len(leftPath))
	part2rightpath := make([]int, len(rightPath))
	copy(part2leftpath, leftPath)
	copy(part2rightpath, rightPath)

	fmt.Println("*** result part1 is *****", part1(leftPath, rightPath))
	fmt.Println("*** result part2 is *****", part2(part2leftpath, part2rightpath))
}

func part1(leftPath, righPath []int) int {
	sort.Ints(leftPath)
	sort.Ints(righPath)
	res := 0

	for i, _ := range leftPath {
		res += utils.AbsInt(leftPath[i] - righPath[i])
	}
	return res
}

func part2(leftPath, rightPath []int) int {
	countMap := make(map[int]int, len(leftPath))
	for i, _ := range leftPath {
		utils.UpdateDict(countMap, rightPath[i])
	}

	result := 0
	for _, v := range leftPath {
		result += (v * utils.FindInDict(countMap, v))
	}
	return result
}
