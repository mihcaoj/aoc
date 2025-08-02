package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var directions = []struct{ dx, dy int }{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func readGrid() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	var grid []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file:", err)
	}
	return grid, nil
}

func findGuardPosition(grid []string) (int, int, error) {
	for y, line := range grid {
		for x, c := range line {
			if c == '^' {
				return x, y, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("Guard not found")
}

func computeNextPosition(x, y, dir int) (int, int) {
	return x + directions[dir].dx, y + directions[dir].dy
}

func isInBounds(x, y int, grid []string) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y])
}

func isObstacle(x, y int, grid []string) bool {
	return grid[y][x] == '#'
}
