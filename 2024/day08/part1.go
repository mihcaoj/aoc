package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	Row, Col int
}

func partOne() {
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
		if len(positions) < 2 {
			continue
		}
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				y := positions[i]
				x := positions[j]

				deltaRow := x.Row - y.Row
				deltaCol := x.Col - y.Col

				antinode1 := Position{
					Row: x.Row + deltaRow,
					Col: x.Col + deltaCol,
				}

				antinode2 := Position{
					Row: y.Row - deltaRow,
					Col: y.Col - deltaCol,
				}

				if isInBounds(antinode1, grid) {
					antinodes[antinode1] = true
					// fmt.Printf("Found antinode at (%d,%d)\n", antinode1.Row, antinode1.Col)
				}

				if isInBounds(antinode2, grid) {
					antinodes[antinode2] = true
					// fmt.Printf("Found antinode at (%d,%d)\n", antinode2.Row, antinode2.Col)
				}
			}
		}
	}
	fmt.Printf("Found a total of %d antinodes\n", len(antinodes))
}

func isInBounds(pos Position, grid []string) bool {
	return pos.Row >= 0 && pos.Row < len(grid) && pos.Col >= 0 && pos.Col < len(grid[pos.Row])
}
