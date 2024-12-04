package day01

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInputDayOne() ([]int, []int, error) {

	const inputFilename = "problems/day01/input.txt"
	file, err := os.Open(inputFilename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Fields(line)
		if len(pairs) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		leftNum, err := strconv.Atoi(pairs[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse left integer: %v", err)
		}

		rightNum, err := strconv.Atoi(pairs[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse right integer: %v", err)
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	return left, right, nil
}

func PartOne() {
	left, right, err := parseInputDayOne()
	if err != nil {
		fmt.Printf("error parsing day one inputs: %v\n", err)
		return
	}

	slices.Sort(left)
	slices.Sort(right)

	var totalDistance int = 0

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	fmt.Println("Total Distance: ", totalDistance)
}

func PartTwo() {
	left, right, err := parseInputDayOne()
	if err != nil {
		fmt.Printf("error parsing day one inputs: %v\n", err)
		return
	}

	countMap := make(map[int]int)

	for _, v := range right {
		count, ok := countMap[v]
		if !ok {
			countMap[v] = 1
		} else {
			countMap[v] = count + 1
		}
	}

	similarityScore := 0

	for _, v := range left {
		count, ok := countMap[v]
		if ok {
			similarityScore += (count * v)
		}
	}

	fmt.Println("Similarity Score: ", similarityScore)
}
