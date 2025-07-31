package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	var data [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		data = append(data, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file:", err)
	}

	findMASCross(data)
}

func findMASCross(data [][]string) {
	count := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "A" {
				if diagonalHasMAndS(data, i-1, j-1, i+1, j+1) && diagonalHasMAndS(data, i-1, j+1, i+1, j-1) {
					count++
				}
			}
		}
	}
	fmt.Printf("Found MAS Cross a total of %d times!\n", count)
}

func diagonalHasMAndS(data [][]string, i1, j1, i2, j2 int) bool {
	if inBounds(data, i1, j1) && inBounds(data, i2, j2) {
		c1 := data[i1][j1]
		c2 := data[i2][j2]
		return (c1 == "M" && c2 == "S") || (c1 == "S" && c2 == "M")
	}
	return false
}

func inBounds(data [][]string, i, j int) bool {
	return i >= 0 && i < len(data) && j >= 0 && j < len(data[i])
}
