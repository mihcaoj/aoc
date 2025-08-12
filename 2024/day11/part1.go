package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stones []int
	for scanner.Scan() {
		line := scanner.Text()
		for part := range strings.FieldsSeq(line) {
			num, _ := strconv.Atoi(part)
			stones = append(stones, num)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file:", err)
	}

	for range 25 {
		var newStones []int
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else {
				s := strconv.Itoa(stone)
				if len(s)%2 == 0 {
					mid := len(s) / 2
					left, _ := strconv.Atoi(s[:mid])
					right, _ := strconv.Atoi(s[mid:])
					newStones = append(newStones, left, right)
				} else {
					newStones = append(newStones, stone*2024)
				}
			}
		}
		stones = newStones
	}
	fmt.Printf("Number of Stones: %d\n", len(stones))
}
