package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	parts := strings.Split(string(file), "\n\n")
	rules := strings.Split(strings.TrimSpace(parts[0]), "\n")
	updates := strings.Split(strings.TrimSpace(parts[1]), "\n")

	ruleMap := make(map[string]bool)
	for _, rule := range rules {
		ruleMap[rule] = true
	}

	// Process each update
	total := 0
	for _, update := range updates {
		up := strings.Split(update, ",")
		var nums []int
		for _, s := range up {
			num, _ := strconv.Atoi(strings.TrimSpace(s))
			nums = append(nums, num)
		}

		// Check if valid and sum the middle numbers
		isValid := validateUpdate(nums, ruleMap)
		if isValid {
			middle := len(nums) / 2
			total += nums[middle]
		}
	}
	fmt.Printf("Total sum of middle page numbers: %d\n", total)
}

func validateUpdate(nums []int, ruleMap map[string]bool) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// Check if there's a rule that says nums[j] should come before nums[i]
			reverseRule := fmt.Sprintf("%d|%d", nums[j], nums[i])
			if ruleMap[reverseRule] {
				return false
			}
		}
	}
	return true
}
