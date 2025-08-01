package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2() {
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
		fmt.Printf("UPDATE: %s\n", up)
		var nums []int
		for _, s := range up {
			num, _ := strconv.Atoi(strings.TrimSpace(s))
			nums = append(nums, num)
		}

		// Check if update was re-ordered
		wasReordered := reorderInvalidUpdate(nums, ruleMap)
		if wasReordered {
			middle := len(nums) / 2
			total += nums[middle]
			fmt.Printf("---> Adding middle number %d, the total is now %d.\n", nums[middle], total)
		}
	}
	fmt.Printf("Total sum of middle page numbers for re-ordered updates: %d\n", total)
}

func reorderInvalidUpdate(nums []int, ruleMap map[string]bool) bool {
	reordered := false
	for {
		swapped := false
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				// Check if there's a rule that says nums[j] should come before nums[i]
				reverseRule := fmt.Sprintf("%d|%d", nums[j], nums[i])
				if ruleMap[reverseRule] {
					fmt.Printf("--> %d before %d breaks rule %d|%d\n", nums[i], nums[j], nums[j], nums[i])
					nums[i], nums[j] = nums[j], nums[i]
					fmt.Printf("--> Swapped %d and %d\n", nums[i], nums[j])
					fmt.Printf("NEW RE-ORDERED UPDATE: %d\n", nums)
					swapped = true
					reordered = true
				}
			}
		}
		if !swapped {
			break
		}
	}
	return reordered
}
