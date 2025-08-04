package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partTwo() {
	file, err := os.Open("input.txt")
	for err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ":")

		targetNum, _ := strconv.Atoi(strings.TrimSpace(numbers[0]))
		operandsStr := strings.TrimSpace(numbers[1])
		operandStrings := strings.Fields(operandsStr) // splits on whitespace

		var operands []int
		for _, opStr := range operandStrings {
			op, _ := strconv.Atoi(opStr)
			operands = append(operands, op)
		}

		if canReachTargetNum(operands, targetNum) {
			result += targetNum
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file:", err)
	}
	fmt.Printf("Result: %d\n", result)
}

// canReachTargetNum checks if we can reach the target using +, *, and ||
func canReachTargetNum(operands []int, target int) bool {
	if len(operands) == 0 {
		return false
	}
	if len(operands) == 1 {
		return operands[0] == target
	}
	return tryOperationsReloaded(operands[0], operands[1:], target)
}

// tryOperationsReloaded recursively tries all combinations of +, *, and || operations
func tryOperationsReloaded(currentValue int, remainingOperands []int, target int) bool {
	if len(remainingOperands) == 0 {
		return currentValue == target
	}
	if currentValue > target {
		return false
	}

	nextOperand := remainingOperands[0]
	remaining := remainingOperands[1:]

	if tryOperationsReloaded(currentValue+nextOperand, remaining, target) {
		return true
	}
	if tryOperationsReloaded(currentValue*nextOperand, remaining, target) {
		return true
	}

	concat := concatenateInts(currentValue, nextOperand)
	if tryOperationsReloaded(concat, remaining, target) {
		return true
	}
	return false
}

func concatenateInts(a, b int) int {
	result, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return result
}
