package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
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

	// Simulate the process and calculate the sum of the 2000th secret numbers
	sum := 0
	for _, initialSecret := range initialSecrets {
		secret := initialSecret
		for i := 0; i < 2000; i++ {
			secret = nextSecretNumber(secret)
		}
		fmt.Println("secret for secret is ", initialSecret, secret)
		sum += secret
	}

	// Print the result
	fmt.Println("Sum of the 2000th secret numbers:", sum)
}
