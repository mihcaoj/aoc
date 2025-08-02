package main

import (
	"fmt"
	"log"
)

func partOne() {
	safeReports, err := processFile("input.txt", isValidReport)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total safe reports: %d\n", safeReports)
}
