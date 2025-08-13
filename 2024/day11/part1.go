package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	var stones []int
	for part := range strings.FieldsSeq(string(file)) {
		num, _ := strconv.Atoi(part)
		stones = append(stones, num)
	}

	for range 25 {
		var newStones []int
		for _, stone := range stones {
			s := strconv.Itoa(stone)
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(s)%2 == 0 {
				mid := len(s) / 2
				left, _ := strconv.Atoi(s[:mid])
				right, _ := strconv.Atoi(s[mid:])
				newStones = append(newStones, left, right)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	fmt.Printf("Number of Stones: %d\n", len(stones))
}
