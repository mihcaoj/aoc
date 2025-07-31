package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MinDiff = 1
	MaxDiff = 3
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// parseNumbers converts space-separated string into a slice of integers
func parseNumbers(line string) ([]int, error) {
	parts := strings.Split(line, " ")
	numbers := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			continue
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

// isValidReport checks if a sequence of numbers follows the safety rules
func isValidReport(numbers []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]

		// Check if all increasing or all decreasing
		if numbers[i] < numbers[i+1] {
			isDecreasing = false
		} else if numbers[i] > numbers[i+1] {
			isIncreasing = false
		}

		// Check if adjacent numbers differ by at least 1 and at most 3
		absDiff := abs(diff)
		if absDiff < MinDiff || absDiff > MaxDiff {
			return false
		}
	}

	return isIncreasing || isDecreasing
}

// isValidWithDampener checks if the report is valid, or can be made valid by removing a single level
func isValidWithDampener(numbers []int) bool {
	if isValidReport(numbers) {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		// create a new slice by copying the elements from numbers up to index i (excluding i)
		shortened := append([]int{}, numbers[:i]...)
		// append all elements from numbers starting from index i+1 to the end
		shortened = append(shortened, numbers[i+1:]...)

		if isValidReport(shortened) {
			return true
		}
	}

	return false
}

// processFile reads and processes the input file and returns the count of safe reports
func processFile(filename string, validationFunc func([]int) bool) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0

	for scanner.Scan() {
		line := scanner.Text()

		numbers, err := parseNumbers(line)
		if err != nil {
			fmt.Printf("Skipping line due to parse error: %v\n", err)
			continue
		}

		if validationFunc(numbers) {
			safeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		return safeReports, fmt.Errorf("Error reading file: %w", err)
	}

	return safeReports, nil
}
