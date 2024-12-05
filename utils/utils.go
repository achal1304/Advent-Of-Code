package utils

import (
	"fmt"
)

func AbsInt(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func UpdateDict[T comparable](countPath map[T]int, path T) {
	if _, ok := countPath[path]; ok {
		countPath[path] += 1
		return
	}
	countPath[path] = 1
}

func FindInDict[T comparable](countPath map[T]int, path T) int {
	return countPath[path]
}

func RemoveSliceElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		fmt.Println("Index out of range")
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}
func RemoveSliceElementByValue(slice []int, value int) []int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func MaxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func UpdateListDict(dict map[int][]int, key int, value int) {
	if _, ok := dict[key]; !ok {
		dict[key] = []int{value}
	} else {
		dict[key] = append(dict[key], value)
	}
}
