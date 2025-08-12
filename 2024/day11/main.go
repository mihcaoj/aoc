package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Running Part 1...")
	start := time.Now()
	partOne()
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)

	// fmt.Println("\nRunning Part 2...")
	// start = time.Now()
	// partTwo()
	// elapsed = time.Since(start)
	// fmt.Println("Execution time:", elapsed)
}
