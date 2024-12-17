package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Register map[string]int

func readInput(file *os.File) (Register, []int) {
	registerRe := regexp.MustCompile(`Register (A|B|C): (\d+)`)
	programRe := regexp.MustCompile(`Program: (.+)`)
	scanner := bufio.NewScanner(file)
	registers := make(Register)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		output := registerRe.FindStringSubmatch(line)
		value, err := strconv.Atoi(output[2])
		if err != nil {
			log.Fatal("error reading value ", err)
		}
		registers[output[1]] = value
	}
	scanner.Scan()
	inputStr := scanner.Text()
	var program []int
	match := programRe.FindStringSubmatch(inputStr)
	if match != nil {
		numbersStr := match[1]

		strNumbers := strings.Split(numbersStr, ",")

		for _, numStr := range strNumbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				break
			}
			program = append(program, num)
		}

	}

	return registers, program

}

func calculateOutputs(registers Register, program []int) []int {
	ip := 0

	// Output collection
	var output []int

	// Run the program
	iteration := 0
	for ip < len(program)-1 {
		iteration += 1
		// Fetch the opcode and operand
		opcode := program[ip]
		operand := program[ip+1]
		ip += 2 // Move instruction pointer past opcode and operand

		// if iteration > 10 {
		// 	break
		// }
		// Execute the corresponding operation
		switch opcode {
		case 0: // adv - division (store result in A)
			registers["A"] = div(registers, registers["A"], operand, registers["B"], registers["C"])
		case 1: // bxl - bitwise XOR (store result in B)
			registers["B"] = bxl(registers, registers["B"], operand)
		case 2: // bst - store operand % 8 in B
			registers["B"] = bst(registers, operand)
		case 3: // jnz - jump if A != 0
			ip = jnz(registers, registers["A"], operand, ip)
		case 4: // bxc - bitwise XOR (store result in B)
			registers["B"] = bxc(registers["B"], registers["C"])
		case 5: // out - output operand % 8
			output = append(output, out(registers, operand, registers["C"]))
		case 6: // bdv - division (store result in B)
			registers["B"] = div(registers, registers["A"], operand, registers["B"], registers["C"])
		case 7: // cdv - division (store result in C)
			registers["C"] = div(registers, registers["A"], operand, registers["B"], registers["C"])
		}
	}

	// Output the final result as a comma-separated string
	// fmt.Println("output for A ", output)
	return output
}

// Helper functions for each opcode operation

// General division function for adv, bdv, cdv
func div(registers Register, A, operand, B, C int) int {
	var divisor int
	divisor = getValueForOperand(registers, operand)

	return A / int(math.Pow(2, float64(divisor)))
}

func getValueForOperand(registers Register, operand int) int {
	value := 0
	switch operand {
	case 0, 1, 2, 3, 7:
		value = operand // 2^operand (for literal operands)
	case 4:
		value = registers["A"] // 2^A (operand refers to register A)
	case 5:
		value = registers["B"] // 2^B (operand refers to register B)
	case 6:
		value = registers["C"] // 2^C (operand refers to register C)
	}
	return value
}

func bxl(registers Register, B, operand int) int {

	return B ^ operand
}

func bst(registers Register, operand int) int {
	value := getValueForOperand(registers, operand)
	return value % 8
}

func jnz(registers Register, A, operand, currentIP int) int {
	if A != 0 {
		return operand // Jump to operand position
	}
	return currentIP + 2 // Move to next instruction
}

func bxc(B, C int) int {
	return B ^ C
}

func out(registers Register, operand, C int) int {
	value := getValueForOperand(registers, operand)

	return value % 8
}

// Helper function to convert int slice to string slice for output
func intSliceToStrSlice(input []int) []string {
	var result []string
	for _, num := range input {
		result = append(result, fmt.Sprintf("%d", num))
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error occured in input.tsxt ", err)
	}
	defer file.Close()

	registers, program := readInput(file)
	// outputPart1 := calculateOutputs(registers, program)
	a := 0
	for n := len(program) - 1; n >= 0; n-- {
		a <<= 3
		registers["A"] = a
		for !slices.Equal(calculateOutputs(registers, program), program[n:]) {
			a++
			registers["A"] = a
		}
		// time.Sleep(2 * time.Second)

		fmt.Println("A", registers["A"], a)
	}

	// fmt.Println("Part1", strings.Join(intSliceToStrSlice(outputPart1), ","))
}
