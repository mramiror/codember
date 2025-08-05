package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	StateStart = iota
	StateDigit
	StateChar
)

func isValid(pwd string) bool {
	state := StateStart
	lastDigit := '0'
	lastChar := 'a'

	for _, c := range pwd {
		switch {
		case c >= '0' && c <= '9':
			if state == StateChar {
				return false
			}
			state = StateDigit
			if c < lastDigit {
				return false
			}
			lastDigit = c
		case c >= 'a' && c <= 'z':
			state = StateChar
			if c < lastChar {
				return false
			}
			lastChar = c
		default:
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	valid, invalid := 0, 0

	for scanner.Scan() {
		pwd := strings.TrimSpace(scanner.Text())
		if pwd == "" {
			continue
		}
		if isValid(pwd) {
			valid++
		} else {
			invalid++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading lines: %v", err)
	}
	fmt.Printf("submit %dtrue%dfalse\n", valid, invalid)
}
