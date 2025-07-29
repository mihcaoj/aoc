package main

import (
	"fmt"
	"log"
)

func Part1() {
	safeReports, err := processFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total safe reports: %d\n", safeReports)
}
