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
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var leftList []int
	var rightList []int
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		leftVal, _ := strconv.Atoi(columns[0])
		rightVal, _ := strconv.Atoi(columns[1])

		leftList = append(leftList, leftVal)
		rightList = append(rightList, rightVal)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning text:", err)
	}

	counts := make(map[int]int)
	for _, val := range rightList {
		counts[val]++
	}

	similarityScore := 0
	for i := range leftList {
		leftNum := leftList[i]
		if count, exists := counts[leftNum]; exists {
			similarityScore += leftNum * count
		}
	}
	fmt.Printf("Similarity score: %d\n", similarityScore)
}
