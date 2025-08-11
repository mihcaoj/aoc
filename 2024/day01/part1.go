package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne() {
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

	sort.Ints(leftList)
	sort.Ints(rightList)

	distances := 0
	for i := range leftList {
		left := leftList[i]
		right := rightList[i]
		diff := abs(right - left)
		distances += diff
	}
	fmt.Printf("Total distance: %d\n", distances)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
