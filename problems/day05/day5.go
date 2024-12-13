package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartOne() {
	filename := "problems/day05/input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file '%s': %v\n", filename, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	parseRules := true
	rules := make(map[int][]int)
	updates := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parseRules = false
			continue
		}

		if parseRules {
			parts := strings.Split(line, "|")
			before, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("error parsing before rule: %v", err)
				os.Exit(1)
			}

			after, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("error parsing before rule: %v", err)
				os.Exit(1)
			}

			rules[before] = append(rules[before], after)

		} else {
			update := []int{}
			pageStrs := strings.Split(line, ",")
			for _, pageStr := range pageStrs {
				page, err := strconv.Atoi(pageStr)
				if err != nil {
					fmt.Printf("error parsing page: %v", err)
				}
				update = append(update, page)
			}
			updates = append(updates, update)
		}
	}

	ans := 0

	for _, update := range updates {
		// helper reverse index
		indexMap := make(map[int]int)
		for i, page := range update {
			indexMap[page] = i
		}

		isCorrect := true

		for _, currPage := range update {
			currPageRules := rules[currPage]
			currPageIdx := indexMap[currPage]
			for _, pageRule := range currPageRules {
				rulePageIdx, ok := indexMap[pageRule]
				if ok && currPageIdx >= rulePageIdx {
					// page from rule exists in update, but is out of order
					isCorrect = false
				}
			}
		}

		if isCorrect {
			updateMidIdx := len(update) / 2
			ans += update[updateMidIdx]
		} else {

		}
	}

	fmt.Printf("Answer: %d\n", ans)
}

func PartTwo() {
	fmt.Println("Not implemented")
}
