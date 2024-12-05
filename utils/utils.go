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

func BinarySearch(arr []int, searchKey int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		// Find the middle index
		mid := low + (high-low)/2

		// Check if searchKey is at mid
		if arr[mid] == searchKey {
			return mid
		}

		// If searchKey is greater, ignore the left half
		if arr[mid] < searchKey {
			low = mid + 1
		} else { // If searchKey is smaller, ignore the right half
			high = mid - 1
		}
	}

	// If the element is not found, return -1
	return -1
}

func SwapElements(slice []int, i, j int) []int {
	// Check for valid indices
	if i < 0 || j < 0 || i >= len(slice) || j >= len(slice) {
		fmt.Println("Invalid indices")
		return slice
	}

	// Extract elements at indices i and j
	elementI := slice[i]
	elementJ := slice[j]

	// Remove elements at i and j
	slice = append(slice[:i], slice[i+1:]...) // Remove element at i
	// If j is after i, we need to adjust j because the slice has shifted
	if j > i {
		j--
	}
	slice = append(slice[:j], slice[j+1:]...) // Remove element at j

	// Insert element j before element i
	slice = append(slice[:i], append([]int{elementJ}, slice[i:]...)...)
	// Finally, insert element i at the end
	slice = append(slice[:i+1], append([]int{elementI}, slice[i+1:]...)...)

	return slice
}
