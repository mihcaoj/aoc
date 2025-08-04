package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
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

		if canReachTarget(operands, targetNum) {
			result += targetNum
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file:", err)
	}
	fmt.Printf("Result: %d\n", result)
}

// canReachTarget checks if we can reach the target using + and *
func canReachTarget(operands []int, target int) bool {
	if len(operands) == 0 {
		return false
	}
	if len(operands) == 1 {
		return operands[0] == target
	}
	return tryOperations(operands[0], operands[1:], target)
}

// tryOperations recursively tries all combinations of + and * operations
func tryOperations(currentValue int, remainingOperands []int, target int) bool {
	if len(remainingOperands) == 0 {
		return currentValue == target
	}
	if currentValue > target {
		return false
	}

	nextOperand := remainingOperands[0]
	remaining := remainingOperands[1:]

	if tryOperations(currentValue+nextOperand, remaining, target) {
		return true
	}
	if tryOperations(currentValue*nextOperand, remaining, target) {
		return true
	}
	return false
}
