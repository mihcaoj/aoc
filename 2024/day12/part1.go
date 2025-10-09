package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var directions = []struct{ dx, dy int }{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file:", err)
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	total := 0
	for i := range rows {
		for j := range cols {
			if !visited[i][j] {
				area, perimeter := dfs(grid, visited, i, j)
				total += area * perimeter
			}
		}
	}
	fmt.Println("Total:", total)
}

func dfs(grid [][]rune, visited [][]bool, i, j int) (int, int) {
	stack := [][2]int{{i, j}}
	visited[i][j] = true
	letter := grid[i][j]

	area := 0
	perimeter := 0
	rows, cols := len(grid), len(grid[0])

	for len(stack) > 0 {
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		x, y := cell[0], cell[1]
		area++

		for _, d := range directions {
			nx, ny := x+d.dx, y+d.dy
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
				perimeter++
				continue
			}
			if grid[nx][ny] != letter {
				perimeter++
				continue
			}
			if !visited[nx][ny] {
				visited[nx][ny] = true
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}

	return area, perimeter
}
