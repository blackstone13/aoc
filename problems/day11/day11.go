package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	value int
	next  *Stone
}

func PartOne() {
	// filename := "problems/day11/sample.txt"
	filename := "problems/day11/input.txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file '%s': %v\n", filename, err)
		os.Exit(1)
	}

	stoneStrs := strings.Fields(string(file))
	var head *Stone
	var curr *Stone

	for _, stoneStr := range stoneStrs {
		value, err := strconv.Atoi(stoneStr)
		if err != nil {
			fmt.Printf("error converting '%s' to int: %v\n", stoneStr, err)
			os.Exit(1)
		}

		stone := Stone{value, nil}
		if head == nil {
			head = &stone
			curr = &stone
		} else {
			curr.next = &stone
			curr = &stone
		}

	}

	curr = head

	for i := 0; i < 75; i++ {
		fmt.Printf("Iteration: %d\n", i)
		for curr != nil {
			if curr.value == 0 {
				curr.value = 1
				curr = curr.next
				continue
			}

			numStr := strconv.Itoa(curr.value)
			if len(numStr)%2 == 0 {
				half := len(numStr) / 2
				numStrL := numStr[:half]
				numStrR := numStr[half:]

				numL, err := strconv.Atoi(numStrL)
				if err != nil {
					fmt.Printf("error converting numL '%s' to int: %v\n", numStrL, numL)
					os.Exit(1)
				}

				numR, err := strconv.Atoi(numStrR)
				if err != nil {
					fmt.Printf("error converting numL '%s' to int: %v\n", numStrR, numR)
				}

				curr.value = numL
				stoneR := Stone{value: numR}
				stoneR.next = curr.next
				curr.next = &stoneR
				curr = stoneR.next
				continue
			}

			curr.value *= 2024
			curr = curr.next
		}

		curr = head
		// for curr != nil {
		// 	fmt.Printf("(%d) -> ", curr.value)
		// 	curr = curr.next
		// }
		// fmt.Printf("\n")
		// curr = head
	}

	count := 0
	curr = head
	for curr != nil {
		count++
		curr = curr.next
	}
	fmt.Printf("Stones: %d\n", count)
}

func PartTwo() {
	fmt.Println("Not implemented.")
}
