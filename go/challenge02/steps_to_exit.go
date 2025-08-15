package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stepsToExit(nums []int) int {
	pos, steps := 0, 0
	for pos >= 0 && pos < len(nums) {
		jump := nums[pos]
		nums[pos]++
		pos += jump
		steps++
	}
	return steps
}

func parseLine(line string) ([]int, error) {
	parts := strings.Fields(line)
	nums := make([]int, len(parts))
	for i, p := range parts {
		v, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("parse %q: %w", p, err)
		}
		nums[i] = v
	}
	return nums, nil
}

func main() {
	file, err := os.Open("trace.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSteps, stepsLastLine := 0, 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		nums, err := parseLine(line)
		if err != nil {
			log.Fatalf("Error parsing line: %v", err)
		}

		steps := stepsToExit(nums)

		totalSteps += steps
		stepsLastLine = steps
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading lines: %v", err)
	}
	fmt.Printf("submit %d-%d\n", totalSteps, stepsLastLine)
}
