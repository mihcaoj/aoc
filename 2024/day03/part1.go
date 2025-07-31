package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	pattern := `mul\(([0-9]{1,3}),([0-9]{1,3})\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Invalid regex pattern:", err)
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	acc := 0
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			product := a * b
			acc += product
		}
	}
	fmt.Printf("Sum of products: %d\n", acc)

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}
}
