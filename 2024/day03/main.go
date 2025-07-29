package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
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
		//fmt.Printf("Input: %s\n", line)

		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			//fmt.Printf("Match: %s\n", match[0])

			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			product := a * b
			//fmt.Printf("%d * %d = %d\n", a, b, product)
			acc += product
		}
	}
	fmt.Printf("Sum of products: %d\n", acc)

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}
}
