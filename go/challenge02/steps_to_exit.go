package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func stepsToExit(nums []int) bool {
	pos := 0
	steps := 0
	n := len(nums)

}

func main() {
	file, err := os.Open("trace.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total_steps, steps_last_line := 0, 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		nums := strings.Split(line, "")

		steps := stepsToExit(nums)

		total_steps += steps
		steps_last_line = steps
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading lines: %v", err)
	}
	fmt.Printf("submit %d-%d\n", total_steps, steps_last_line)
}
