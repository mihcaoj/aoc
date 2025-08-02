package main

import (
	"fmt"
	"log"
)

func partTwo() {
	safeReports, err := processFile("input.txt", isValidWithDampener)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total safe reports: %d\n", safeReports)
}
