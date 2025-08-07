package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func partTwo() {
	file, err := os.Open("input.txt")
	for err != nil {
		log.Fatal("Error reading file:", err)
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

	// Find the positions of the antennas and store them in a map
	antennas := make(map[rune][]Position)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			char := rune(grid[row][col])
			if char != '.' {
				antennas[char] = append(antennas[char], Position{Row: row, Col: col})
			}
		}
	}

	// Find the positions of the antinodes and store them in a map
	antinodes := make(map[Position]bool)
	for _, positions := range antennas {
		if len(positions) >= 2 {
			for _, pos := range positions {
				antinodes[pos] = true
			}
		}
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				y := positions[i]
				x := positions[j]

				deltaRow := x.Row - y.Row
				deltaCol := x.Col - y.Col

				currentPos := x
				for {
					nextPos := Position{
						Row: currentPos.Row + deltaRow,
						Col: currentPos.Col + deltaCol,
					}
					if !isInBounds(nextPos, grid) {
						break
					}
					antinodes[nextPos] = true
					// fmt.Printf("Found antinode at (%d,%d)\n", nextPos.Row, nextPos.Col)
					currentPos = nextPos
				}

				currentPos = y
				for {
					nextPos := Position{
						Row: currentPos.Row - deltaRow,
						Col: currentPos.Col - deltaCol,
					}
					if !isInBounds(nextPos, grid) {
						break
					}
					antinodes[nextPos] = true
					// fmt.Printf("Found antinode at (%d,%d)\n", nextPos.Row, nextPos.Col)
					currentPos = nextPos
				}
			}
		}
	}
	fmt.Printf("Found a total of %d antinodes\n", len(antinodes))
}
