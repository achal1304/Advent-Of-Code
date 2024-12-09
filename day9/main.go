package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/achal1304/Advent-Of-Code/utils"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error while opening file ", err)
	}
	defer file.Close()
	diskMapsBlocks, freeSpaceBlocks, blockIds := scanInput(file)
	fmt.Println(diskMapsBlocks)
	fmt.Println(freeSpaceBlocks)
	// part 1
	// blockMaps := PrepareBlocks(diskMapsBlocks, freeSpaceBlocks)
	// part 2
	blockMaps := PrepareBlocksWithFSFragmentation(diskMapsBlocks, freeSpaceBlocks, blockIds)
	// fmt.Println(blockMaps)
	fmt.Println(calculateChecksum(blockMaps))
}

func scanInput(file *os.File) ([]int, []int, []int) {
	diskMapBlocks := []int{}
	freeSpaceBlocks := []int{}
	blockIds := []int{}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputText := scanner.Text()
	iteration := 0
	isDiskMapBlocks := true
	for i, ele := range inputText {
		num, err := strconv.Atoi(string(ele))
		if err != nil {
			fmt.Errorf("error reading input element %d %d", ele, i)
		}
		if isDiskMapBlocks {
			diskMapBlocks = append(diskMapBlocks, num)
			blockIds = append(blockIds, iteration)
			iteration += 1
			isDiskMapBlocks = false
		} else {
			freeSpaceBlocks = append(freeSpaceBlocks, num)
			isDiskMapBlocks = true
		}

	}

	if len(freeSpaceBlocks) < len(diskMapBlocks) {
		freeSpaceBlocks = append(freeSpaceBlocks, 0)
	}
	return diskMapBlocks, freeSpaceBlocks, blockIds
}

// Part 1
func PrepareBlocks(diskMapBlocks []int, freeSpaceBlocks []int) []int {
	blocksMappings := []int{}
	r := len(diskMapBlocks) - 1
	for l := 0; l < len(diskMapBlocks); l++ {
		for i := 0; i < diskMapBlocks[l]; i++ {
			blocksMappings = append(blocksMappings, l)
		}
		// fmt.Println("block mappings after diskblocks", blocksMappings)
		if l == r {
			return blocksMappings
		}

		for i := 0; i < freeSpaceBlocks[l]; i++ {
			if l == r {
				return blocksMappings
			}
			if diskMapBlocks[r] > 0 {
				blocksMappings = append(blocksMappings, r)
				diskMapBlocks[r] -= 1
				if diskMapBlocks[r] <= 0 {
					r -= 1
				}
			}
		}
		// fmt.Println("block mappings after freeblocks", blocksMappings)

	}
	return blocksMappings
}

// Part 2
func PrepareBlocksWithFSFragmentation(diskMapBlocks, freeSpaceBlocks, blockIds []int) []int {
	shiftedElements := make(map[int]int, len(diskMapBlocks))
	shiftedElements[0] = 1
	r := len(diskMapBlocks) - 1
outerLoops:
	for r >= 0 && len(shiftedElements) <= len(diskMapBlocks) {
		_, isShifted := shiftedElements[blockIds[r]]
		if isShifted {
			r -= 1
			continue outerLoops
		} else {
			shiftedElements[blockIds[r]] = 1
		}
		for l := 0; l < r; l++ {
			if diskMapBlocks[r] <= freeSpaceBlocks[l] {
				// update the freespace by adding the diskmapblcok of the element which is to be swapped
				// as after swapping it will leave empty space behind
				freeSpaceBlocks[r-1] = freeSpaceBlocks[r-1] + diskMapBlocks[r] + freeSpaceBlocks[r]

				// update the free space block of the swapped element as the place where it is swapped needs
				// to adjust the remaining freespace of previous element
				freeSpaceBlocks[r] = freeSpaceBlocks[l] - diskMapBlocks[r]

				// make freespace block of the element before the swapped element as 0 as the next element which
				// is actually swapped will adjust the free space for further operations
				freeSpaceBlocks[l] = 0

				diskMapBlocks = utils.ShiftElementOfArray(diskMapBlocks, diskMapBlocks[r], l+1, r)
				freeSpaceBlocks = utils.ShiftElementOfArray(freeSpaceBlocks, freeSpaceBlocks[r], l+1, r)
				blockIds = utils.ShiftElementOfArray(blockIds, blockIds[r], l+1, r)
				continue outerLoops
			}
		}
		r -= 1
	}

	blockMappings := formBlocksAfterFSFragmentation(diskMapBlocks, freeSpaceBlocks, blockIds)

	return blockMappings
}

func formBlocksAfterFSFragmentation(diskMapBlocks, freeSpaceBlocks, blockIds []int) []int {
	fSFragmentationBlocks := []int{}

	for i := 0; i < len(diskMapBlocks); i++ {
		for j := 0; j < diskMapBlocks[i]; j++ {
			fSFragmentationBlocks = append(fSFragmentationBlocks, blockIds[i])
		}
		for j := 0; j < freeSpaceBlocks[i]; j++ {
			fSFragmentationBlocks = append(fSFragmentationBlocks, 0)
		}
	}
	return fSFragmentationBlocks
}

func calculateChecksum(blockMappings []int) int {
	totCount := 0
	for i := 0; i < len(blockMappings); i++ {
		count := i * blockMappings[i]
		totCount += count
	}
	return totCount
}
