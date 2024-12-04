package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}

		if diff > 0 != increasing {
			return false
		}
	}
	return true
}

func PartOne() {

	inputFilename := "problems/day02/input.txt"
	file, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("error opening file %s: %v\n", inputFilename, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	safeCount := 0

	for scanner.Scan() {
		report := scanner.Text()
		levelStrs := strings.Fields(report)
		levels := make([]int, len(levelStrs))

		for i, levelStr := range levelStrs {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				fmt.Printf("failed to level as integer: %v\n", err)
				os.Exit(1)
			}
			levels[i] = level
		}

		if checkSafe(levels) {
			safeCount++
		}
	}

	fmt.Printf("Safe Reports: %d\n", safeCount)
}

func validTransition(a, b int) bool {
	diff := b - a
	return diff >= -3 && diff <= 3 && diff != 0
}

func checkSafeWithTolerance(levels []int) bool {
	if checkSafe(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		modified := append([]int{}, levels[:i]...)
		modified = append(modified, levels[i+1:]...)

		if checkSafe(modified) {
			return true
		}
	}

	return false
}

func PartTwo() {

	inputFilename := "problems/day02/input.txt"
	file, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("error opening file %s: %v\n", inputFilename, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	safeCount := 0

	for scanner.Scan() {
		report := scanner.Text()
		levelStrs := strings.Fields(report)
		levels := make([]int, len(levelStrs))

		for i, levelStr := range levelStrs {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				fmt.Printf("failed to parse level as integer: %v\n", err)
				os.Exit(1)
			}
			levels[i] = level
		}

		if checkSafeWithTolerance(levels) {
			safeCount++
		}
	}

	fmt.Printf("Safe Reports with Tolerance: %d\n", safeCount)
}
