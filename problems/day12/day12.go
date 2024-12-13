package day12

import (
	"bufio"
	"fmt"
	"os"
)

func PartOne() {
	filename := "problems/day12/sample.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file '%s': %v", filename, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	farm := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		farm = append(farm, []rune(line))
	}

	for row := 0; row < len(farm); row++ {
		for col := 0; col < len(farm[0]); col++ {

		}
	}

	fmt.Printf("Farm:\n%v\n", farm)
}

func PartTwo() {
	fmt.Println("Not implemented.")
}
