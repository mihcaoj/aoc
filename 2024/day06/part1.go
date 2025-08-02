package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	X, Y int
}

var directions = []struct{ dx, dy int }{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func Part1() {
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

	var guardX, guardY int
	found := false
	for y, line := range grid {
		for x, c := range line {
			if c == '^' {
				guardX, guardY = x, y
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	// fmt.Printf("Found guard at %d,%d\n", guardX, guardY)
	visited := make(map[Position]struct{})

	// Mark current position of the guard as visited
	pos := Position{X: guardX, Y: guardY}
	visited[pos] = struct{}{}

	isInGrid := true
	dir := 0 // facing up at first
	for isInGrid {
		// Compute next position of the guard
		newGuardX, newGuardY := guardX, guardY
		newGuardX += directions[dir].dx
		newGuardY += directions[dir].dy

		// Check bounds
		if newGuardY < 0 || newGuardY >= len(grid) || newGuardX < 0 || newGuardX >= len(grid) {
			// Guard left
			isInGrid = false
			break
		}

		// Check collisions
		if grid[newGuardY][newGuardX] == '#' {
			// We hit an obstacle - turn right
			dir = (dir + 1) % len(directions)
			continue
		}

		// fmt.Printf("x,y = %d,%d\n", newGuardX, newGuardY)

		var guardChar rune
		switch dir {
		case 0:
			guardChar = '^'
		case 1:
			guardChar = '>'
		case 2:
			guardChar = 'v'
		case 3:
			guardChar = '<'
		}

		// Mark old pos with '.'
		oldRow := []rune(grid[guardY])
		oldRow[guardX] = '.'
		grid[guardY] = string(oldRow)

		// Mark new pos with dir char
		newRow := []rune(grid[newGuardY])
		newRow[newGuardX] = guardChar
		grid[newGuardY] = string(newRow)

		// time.Sleep(150 * time.Millisecond)
		// fmt.Print("\033[H\033[2J") // clear screen
		// printGrid(grid)

		// Update position
		guardX, guardY = newGuardX, newGuardY

		// Mark position in grid as visited
		pos := Position{X: newGuardX, Y: newGuardY}
		if _, ok := visited[pos]; !ok {
			visited[pos] = struct{}{}
			// fmt.Printf("%v\n", visited)
		}
	}
	distinctPositions := len(visited)
	fmt.Println("Distinct positions visited:", distinctPositions)
}

func printGrid(grid []string) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}
