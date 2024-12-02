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

	fmt.Println("left", leftPath)
	fmt.Println("right", rightPath)

	fmt.Println("*** result is *****", part1(leftPath, rightPath))
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

// func part2(leftPath, rightPath []int) int {

// }
