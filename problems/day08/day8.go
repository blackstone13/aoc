package day08

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	// Symbol   rune
	Row, Col int
}

func PartOne() {
	filename := "problems/day08/input.txt"
	// filename := "problems/day08/sample.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file '%s': %v", filename, err)
	}
	defer file.Close()

	atMap := make(map[rune][]Node)
	antinodeMap := make(map[Node]int)

	scanner := bufio.NewScanner(file)
	scanRow := 0
	scanCol := 0

	for scanner.Scan() {
		line := scanner.Text()

		if scanCol == 0 {
			scanCol = len(line)
		}

		for i, char := range line {
			if char != '.' {
				antenna := Node{Row: scanRow, Col: i}
				atMap[char] = append(atMap[char], antenna)
			}
		}
		scanRow++
	}

	for at, locations := range atMap {
		fmt.Printf("%c - %v\n", at, locations)
		numLocations := len(locations)
		if numLocations < 2 {
			continue
		}

		for i := 0; i < numLocations-1; i++ {
			for j := i + 1; j < numLocations; j++ {
				loc1 := locations[i]
				loc2 := locations[j]
				dy := loc1.Row - loc2.Row
				dx := loc1.Col - loc2.Col

				antinodeMap[loc1] = 1
				antinodeMap[loc2] = 1

				fmt.Printf("%v <- (%d, %d) -> %v\n", loc1, dy, dx, loc2)
				anti1y := loc1.Row + dy
				anti1x := loc1.Col + dx
				fmt.Printf("First Antinode: (%d, %d) - ", anti1y, anti1x)
				for anti1y >= 0 && anti1y < scanRow && anti1x >= 0 && anti1x < scanCol {
					anti1 := Node{anti1y, anti1x}
					antinodeMap[anti1] = 1

					anti1y += dy
					anti1x += dx
				}

				anti2y := loc2.Row - dy
				anti2x := loc2.Col - dx
				fmt.Printf("Second Antinode: (%d, %d) - ", anti2y, anti2x)
				for anti2y >= 0 && anti2y < scanRow && anti2x >= 0 && anti2x < scanCol {
					anti2 := Node{anti2y, anti2x}
					antinodeMap[anti2] = 1

					anti2y -= dy
					anti2x -= dx

				}
			}
		}
	}

	fmt.Printf("Antinodes Found: %d\n", len(antinodeMap))
}
