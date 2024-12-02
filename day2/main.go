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
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var levels [][]int
	for scanner.Scan() {
		var levelStrings []string
		var level []int
		input := scanner.Text()
		levelStrings = strings.Fields(input)
		for _, v := range levelStrings {
			levelInt, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("cannot convert levels to int ", level)
			}
			level = append(level, levelInt)
		}
		levels = append(levels, level)
	}

	fmt.Println(GetSafeReports(levels))
	fmt.Println(SafeReportsWithProblemDampener(levels))

}

func GetSafeReports(levels [][]int) int {
	result := 0
	for _, level := range levels {
		if isSafe, _ := isReportSafe(level); isSafe {
			result += 1
		}
	}
	return result
}

func isReportSafe(level []int) (bool, int) {
	var incremental bool
	var decremental bool
	for j := 1; j < len(level); j++ {

		if level[j-1] > level[j] {
			if !(0 < level[j-1]-level[j] && level[j-1]-level[j] < 4) || incremental {
				return false, j
			}
			decremental = true
		} else {
			if !(0 < level[j]-level[j-1] && level[j]-level[j-1] < 4) || decremental {
				return false, j
			}
			incremental = true
		}
	}
	return (incremental || decremental) && !(incremental && decremental), 0
}

func SafeReportsWithProblemDampener(levels [][]int) int {
	result := 0
outerLoop:
	for _, level := range levels {
		isSafe, unsafeIndex := isReportSafe(level)
		if !isSafe {
			for i := utils.MaxNumber(0, unsafeIndex-2); i < len(level); i++ {
				newlevel := make([]int, len(level))
				copy(newlevel, level)
				newlevel = append(newlevel[:i], newlevel[i+1:]...)

				isSafe, _ = isReportSafe(newlevel)
				if !isSafe {
					continue
				} else {
					fmt.Println("level", level)
					result += 1
					continue outerLoop
				}
			}
		} else {
			fmt.Println("level", level)
			result += 1
		}
	}
	return result
}
