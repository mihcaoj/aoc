package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var inputString string
	for scanner.Scan() {
		line := scanner.Text()
		inputString += line
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning text:", err)
	}

	// Convert input string to digits
	digits := make([]int, len(inputString))
	for i, char := range inputString {
		digits[i] = int(char - '0')
	}

	// Decompress
	var blocks []int
	fileId := 0
	for i, digit := range digits {
		if i%2 == 0 {
			// found file
			for j := 0; j < digit; j++ {
				blocks = append(blocks, fileId)
			}
			fileId++
		} else {
			// found free space
			for j := 0; j < digit; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	// Defragment
	for {
		leftFree := -1
		for i, block := range blocks {
			if block == -1 {
				leftFree = i
				break
			}
		}
		if leftFree == -1 {
			break
		}

		rightFile := -1
		for i := len(blocks) - 1; i > leftFree; i-- {
			if blocks[i] != -1 {
				rightFile = i
				break
			}
		}
		if rightFile == -1 {
			break
		}

		blocks[leftFree], blocks[rightFile] = blocks[rightFile], blocks[leftFree]
	}

	// Checksum
	checksum := 0
	for i, block := range blocks {
		if block != -1 {
			checksum += i * block
		}
	}
	fmt.Println("Filesystem checksum:", checksum)
}
