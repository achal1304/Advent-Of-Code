package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Memo struct {
	cache map[string]int
}

func NewMemo() *Memo {
	return &Memo{cache: make(map[string]int)}
}

func (m *Memo) Get(key string) (int, bool) {
	val, exists := m.cache[key]
	return val, exists
}

func (m *Memo) Set(key string, value int) {
	m.cache[key] = value
}

// Function to calculate the next secret number
func nextSecretNumber(secret int) int {
	// Step 1: Multiply by 64, mix, and prune
	secret ^= (secret * 64)
	secret %= 16777216

	// Step 2: Divide by 32, round down, mix, and prune
	secret ^= (secret / 32)
	secret %= 16777216

	// Step 3: Multiply by 2048, mix, and prune
	secret ^= (secret * 2048)
	secret %= 16777216

	return secret
}

// Function to generate prices and changes for a buyer
func generatePricesAndChanges(initialSecret int, count int, memo *Memo) ([]int, []int) {
	prices := make([]int, count)
	changes := make([]int, count)
	secret := initialSecret
	prices[0] = secret % 10
	visited := NewMemo()
	changes[0] = -9999

	for i := 1; i < count; i++ {
		secret = nextSecretNumber(secret)
		prices[i] = secret % 10
		if i > 0 {
			changes[i] = prices[i] - prices[i-1]
		}

	}

	for i := 1; i < len(changes)-3; i++ {
		key := fmt.Sprintf("%d-%d-%d-%d", changes[i], changes[i+1], changes[i+2], changes[i+3])
		if _, ok := visited.cache[key]; !ok {
			memo.cache[key] += prices[i+3]
			visited.cache[key] += 1
		}
	}
	return prices, changes
}

// Function to find the best sequence of changes
func findBestSequence(initialSecrets []int, memoCache *Memo) (string, int) {
	for _, initialSecret := range initialSecrets {
		generatePricesAndChanges(initialSecret, 2001, memoCache)
	}

	maxBananas := 0
	bestSequence := ""
	for sequence, value := range memoCache.cache {
		if value > maxBananas {
			maxBananas = value
			bestSequence = sequence
		}
	}

	// fmt.Println(memoCache)
	return bestSequence, maxBananas
}

func main() {
	// Open the input file
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	// Read the input file using bufio.NewScanner
	scanner := bufio.NewScanner(file)
	var initialSecrets []int

	for scanner.Scan() {
		line := scanner.Text()
		if secret, err := strconv.Atoi(line); err == nil {
			initialSecrets = append(initialSecrets, secret)
		} else {
			fmt.Println("Error parsing line:", line)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	memoCache := NewMemo()

	// Find the best sequence and the maximum bananas
	bestSequence, maxBananas := findBestSequence(initialSecrets, memoCache)

	// Print the results
	fmt.Println("Best sequence:", bestSequence)
	fmt.Println("Maximum bananas:", maxBananas)
}
