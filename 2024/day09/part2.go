package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func partTwo() {
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

	maxFileId := fileId - 1
	for currentFileId := maxFileId; currentFileId >= 0; currentFileId-- {
		fileStart := -1
		fileSize := 0

		for i, block := range blocks {
			if block == currentFileId {
				if fileStart == -1 {
					fileStart = i
				}
				fileSize++
			}
		}

		// Find the leftmost free space that can fit this file
		freeStart := findLeftmostFreeSpace(blocks, fileSize, fileStart)

		if freeStart != -1 {
			// clear the original position
			for i := fileStart; i < fileStart+fileSize; i++ {
				blocks[i] = -1
			}
			// place it in the new position
			for i := freeStart; i < freeStart+fileSize; i++ {
				blocks[i] = currentFileId
			}
		}
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

func findLeftmostFreeSpace(blocks []int, size int, beforeIndex int) int {
	for i := 0; i <= beforeIndex-size; i++ {
		if blocks[i] == -1 {
			// found start of potential free space
			contiguousSize := 0
			for j := i; j < len(blocks) && blocks[j] == -1; j++ {
				contiguousSize++
			}
			if contiguousSize >= size {
				return i
			}
			i += contiguousSize - 1
		}
	}
	return -1
}
