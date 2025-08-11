package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, ch := range line {
			row[i] = int(ch - '0') // convert rune to digit
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning text:", err)
	}

	trailheads := []Point{}
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 0 {
				trailheads = append(trailheads, Point{row, col})
			}
		}
	}

	trailheadScores := 0
	for _, head := range trailheads {
		trailheadScores += dfs2(head.Row, head.Col, 0, grid)
	}
	fmt.Println("Trailhead scores:", trailheadScores)
}

func dfs2(row, col, currentHeight int, grid [][]int) int {
	if currentHeight == 9 {
		return 1
	}

	count := 0
	for _, direction := range directions {
		newRow := row + direction.dy
		newCol := col + direction.dx

		if !isInBounds(newRow, newCol, grid) {
			continue
		}

		if grid[newRow][newCol] == currentHeight+1 {
			count += dfs2(newRow, newCol, currentHeight+1, grid)
		}
	}
	return count
}
