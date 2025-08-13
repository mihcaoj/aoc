package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type key struct {
	stone int
	steps int
}

var memo = make(map[key]int)

func partTwo() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	var stones []int
	for part := range strings.FieldsSeq(string(file)) {
		num, _ := strconv.Atoi(part)
		stones = append(stones, num)
	}

	blinks := 75
	sum := 0
	for _, stone := range stones {
		sum += count(stone, blinks)
	}
	fmt.Printf("Number of Stones: %d\n", sum)
}

func count(stone, steps int) int {
	k := key{stone, steps}
	if val, ok := memo[k]; ok {
		return val
	}
	if steps == 0 {
		return 1
	}
	if stone == 0 {
		result := count(1, steps-1)
		memo[k] = result
		return result
	}

	s := strconv.Itoa(stone)
	length := len(s)

	var result int
	if length%2 == 0 {
		left, _ := strconv.Atoi(s[:length/2])
		right, _ := strconv.Atoi(s[length/2:])
		result = count(left, steps-1) + count(right, steps-1)
	} else {
		result = count(stone*2024, steps-1)
	}

	memo[k] = result
	return result
}
