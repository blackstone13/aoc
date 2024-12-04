package day04

import (
	"bufio"
	"fmt"
	"os"
)

func PartOne() {
	filename := "problems/day04/input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file '%s': %v", filename, err)
	}

	scanner := bufio.NewScanner(file)

	crossword := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		crossword = append(crossword, line)
	}

	count := 0

	for i := 0; i < len(crossword); i++ {
		// for i := 4; i < 5; i++ {
		for j := 0; j < len(crossword[0]); j++ {
			if crossword[i][j] == 'X' {
				// Search Forward
				if search(crossword, j, i, 1, 0) {
					count++
				}
				// Search Backwards
				if search(crossword, j, i, -1, 0) {
					count++
				}
				// Search Down
				if search(crossword, j, i, 0, 1) {
					count++
				}
				// Search Up
				if search(crossword, j, i, 0, -1) {
					count++
				}
				// Search Forward & Down
				if search(crossword, j, i, 1, 1) {
					count++
				}
				// Search Forward & Up
				if search(crossword, j, i, 1, -1) {
					count++
				}
				// Search Backwards & Down
				if search(crossword, j, i, -1, 1) {
					count++
				}
				// Search Backwards & Up
				if search(crossword, j, i, -1, -1) {
					count++
				}
			}
		}
	}
	fmt.Printf("Total XMAS Count: %d\n", count)
}

func search(crossword []string, posX int, posY int, dX int, dY int) bool {
	word := "XMAS"
	rows := len(crossword)
	cols := len(crossword[0])

	for wi := 1; wi < len(word); wi++ {
		newX := posX + (dX * wi)
		newY := posY + (dY * wi)
		if newX < 0 || newX >= cols || newY < 0 || newY >= rows {
			return false
		}

		if crossword[newY][newX] != word[wi] {
			return false
		}

	}
	return true

}

func PartTwo() {

	filename := "problems/day04/input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file '%s': %v", filename, err)
	}

	scanner := bufio.NewScanner(file)

	crossword := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		crossword = append(crossword, line)
	}

	count := 0

	for i := 0; i < len(crossword); i++ {
		for j := 0; j < len(crossword[0]); j++ {
			if crossword[i][j] == 'A' && i+1 < len(crossword) && j+1 < len(crossword[0]) && i-1 >= 0 && j-1 >= 0 {
				// X-MAS Possible - Is centered around 'A' and valid bounds.
				d1 := string(crossword[i-1][j-1]) + "A" + string(crossword[i+1][j+1])
				d2 := string(crossword[i+1][j-1]) + "A" + string(crossword[i-1][j+1])

				if (d1 == "MAS" || d1 == "SAM") && (d2 == "MAS" || d2 == "SAM") {
					count++
				}
			}
		}
	}

	fmt.Printf("Total X-MAS Found: %d\n", count)
}
