package main

import (
	"fmt"
	"log"
)

type Position struct {
	X, Y int
}

func partOne() {
	grid, err := readGrid()
	if err != nil {
		log.Fatal("Error reading grid:", err)
	}

	guardX, guardY, err := findGuardPosition(grid)
	if err != nil {
		log.Fatal("Error fiding guard position:", err)
	}

	visited := make(map[Position]struct{})

	// Mark current position of the guard as visited
	pos := Position{X: guardX, Y: guardY}
	visited[pos] = struct{}{}

	isInGrid := true
	dir := 0 // facing up at first
	for isInGrid {
		// Compute next position of the guard
		newGuardX, newGuardY := computeNextPosition(guardX, guardY, dir)

		// Check bounds
		if !isInBounds(newGuardX, newGuardY, grid) {
			isInGrid = false
			break
		}

		// Check collisions
		if isObstacle(newGuardX, newGuardY, grid) {
			// We hit an obstacle - turn right
			dir = (dir + 1) % len(directions)
			continue
		}

		// var guardChar rune
		// switch dir {
		// case 0:
		// 	guardChar = '^'
		// case 1:
		// 	guardChar = '>'
		// case 2:
		// 	guardChar = 'v'
		// case 3:
		// 	guardChar = '<'
		// }

		// Mark old pos with '.'
		// oldRow := []rune(grid[guardY])
		// oldRow[guardX] = '.'
		// grid[guardY] = string(oldRow)

		// // Mark new pos with dir char
		// newRow := []rune(grid[newGuardY])
		// newRow[newGuardX] = guardChar
		// grid[newGuardY] = string(newRow)

		// time.Sleep(150 * time.Millisecond)
		// fmt.Print("\033[H\033[2J") // clear screen
		// printGrid(grid)

		// Update position
		guardX, guardY = newGuardX, newGuardY

		// Mark position in grid as visited
		pos := Position{X: newGuardX, Y: newGuardY}
		if _, ok := visited[pos]; !ok {
			visited[pos] = struct{}{}
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
