package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func partOne() {
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

	word := "XMAS"
	findWord(data, word)
}

func findWord(data [][]string, word string) {
	directions := []struct{ dx, dy int }{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	count := 0
	for i, row := range data {
		for j := range row {
			for _, d := range directions {
				if matches(data, word, i, j, d.dx, d.dy) {
					count++
				}
			}
		}
	}
	fmt.Printf("Found %s a total of %d times!\n", word, count)
}

func matches(data [][]string, word string, startI, startJ, dx, dy int) bool {
	for k := 0; k < len(word); k++ {
		newI := startI + k*dx
		newJ := startJ + k*dy

		// Check bounds
		if newI < 0 || newI >= len(data) || newJ < 0 || newJ >= len(data[newI]) {
			return false
		}
		// Check if grid char matches k-th char of the word
		if data[newI][newJ] != string(word[k]) {
			return false
		}
	}
	return true
}
