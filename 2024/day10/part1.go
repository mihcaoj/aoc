package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	Row, Col int
}

var directions = []struct{ dx, dy int }{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func partOne() {
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
		visited := make(map[Point]bool)
		trailheadScores += dfs(head.Row, head.Col, 0, grid, visited)
	}
	fmt.Println("Trailhead scores:", trailheadScores)
}

func dfs(row, col, currentHeight int, grid [][]int, visited map[Point]bool) int {
	if currentHeight == 9 {
		point := Point{row, col}
		if visited[point] {
			return 0
		}
		visited[point] = true
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
			count += dfs(newRow, newCol, currentHeight+1, grid, visited)
		}
	}
	return count
}

func isInBounds(row, col int, grid [][]int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}
