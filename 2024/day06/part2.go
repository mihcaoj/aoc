package main

import (
	"fmt"
	"log"
)

type State struct {
	X, Y, Dir int
}

func partTwo() {
	grid, err := readGrid()
	if err != nil {
		log.Fatal("Error reading grid:", err)
	}

	// Scan through the grid and collect positions where we can place obstacle
	var candidates []struct{ X, Y int }
	for y, row := range grid {
		for x, c := range row {
			if c == '.' {
				candidates = append(candidates, struct{ X, Y int }{X: x, Y: y})
			}
		}
	}

	// Try placing an obstacle at each candidate position
	loopCount := 0
	for _, candidate := range candidates {
		gridClone := cloneGrid(grid)
		row := []rune(gridClone[candidate.Y])
		row[candidate.X] = '#'
		gridClone[candidate.Y] = string(row)

		if simulation(gridClone) {
			loopCount++
			// fmt.Printf("Adding obstacle at %d,%d creates a loop\n", candidate.X, candidate.Y)
		}
	}
	fmt.Println("Total loop count:", loopCount)
}

func simulation(grid []string) bool {
	guardX, guardY, err := findGuardPosition(grid)
	if err != nil {
		log.Fatal("Error fiding guard position:", err)
	}

	visited := make(map[State]struct{})
	dir := 0 // facing up at first
	state := State{X: guardX, Y: guardY, Dir: dir}
	visited[state] = struct{}{}
	isInGrid := true
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

		// Have we already visited this spot with the same direction?
		state := State{X: newGuardX, Y: newGuardY, Dir: dir}
		if _, ok := visited[state]; ok {
			// we found a loop
			return true
		}
		visited[state] = struct{}{}

		// Update position
		guardX, guardY = newGuardX, newGuardY
	}
	return false
}

func cloneGrid(grid []string) []string {
	newGrid := make([]string, len(grid))
	copy(newGrid, grid)
	return newGrid
}
