package utils

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
