package day03

import (
	"fmt"
	"os"
	// "bufio"
	"regexp"
	"strconv"
)

func PartOne() {
	filename := "problems/day03/input.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file '%s': %v\n", filename, err)
	}

	// validStatement := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	validStatement := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	statements := validStatement.FindAllStringSubmatch(string(content), -1)
	total := 0

	for _, statement := range statements {
		a, err := strconv.Atoi(statement[1])
		if err != nil {
			fmt.Printf("error parsing 'a' value: %v", err)
			os.Exit(1)
		}

		b, err := strconv.Atoi(statement[2])
		if err != nil {
			fmt.Printf("error parsing 'a' value: %v", err)
			os.Exit(1)
		}

		total += a * b
	}

	fmt.Printf("Total: %d\n", total)
	// fmt.Printf("Debug: %v\n", string(content))
}

// Custom Parser instead of Regex
func PartTwo() {
	filename := "problems/day03/input.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file '%s': %v\n", filename, err)
	}

	input := string(content)
	total := 0
	mulEnabled := true
	length := len(input)

	for i := 0; i < length; {
		if input[i] == 'm' && i+3 < length && input[i:i+4] == "mul(" {
			i += 4

			a, pos, ok := parseNumber(input, i)
			if !ok {
				continue
			}
			i = pos

			if i > length || input[i] != ',' {
				continue
			}
			i++

			b, pos, ok := parseNumber(input, i)
			if !ok {
				continue
			}
			i = pos

			if i >= length || input[i] != ')' {
				continue
			}
			i++

			if mulEnabled {
				total += a * b
			}
			continue
		} else if input[i] == 'd' {
			if i+3 < length && input[i:i+4] == "do()" {
				i += 4
				mulEnabled = true
				continue
			} else if i+6 < length && input[i:i+7] == "don't()" {
				i += 7
				mulEnabled = false
				continue
			}
		}
		i++

	}

	fmt.Printf("Total: %d\n", total)
}

func parseNumber(input string, pos int) (int, int, bool) {
	start := pos
	digits := 0

	for pos < len(input) && digits < 3 && input[pos] >= '0' && input[pos] <= '9' {
		pos++
		digits++
	}

	if digits == 0 {
		return 0, pos, false
	}

	numStr := input[start:pos]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, pos, false
	}

	return num, pos, true
}
