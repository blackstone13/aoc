package day06

import (
	"bufio"
	"fmt"
	"os"
)

type Guard struct {
	row, col, dRow, dCol int
}

func PartOne() {
	// filename := "problems/day06/input.txt"
	filename := "problems/day06/sample.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s: %v", filename, err)
	}
	defer file.Close()

	room := [][]rune{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		room = append(room, []rune(scanner.Text()))
	}

	numRows := len(room)
	numCols := len(room[0])
	spacesVisited := 0
	possibleLoops := 0
	// up - right - down - left
	dirs := []struct{ dRow, dCol int }{
		{dRow: -1, dCol: 0},
		{dRow: 0, dCol: 1},
		{dRow: 1, dCol: 0},
		{dRow: 0, dCol: -1},
	}
	dirIdx := 0
	guardDir := dirs[dirIdx]
	guardRow := 0
	guardCol := 0

	// find guard starting position
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if room[row][col] == '^' {
				// create guard facing up
				guardRow = row
				guardCol = col
			}
		}
	}

	fmt.Printf("Guard at row=%d col=%d\n", guardRow, guardCol)

	// while guard has not left area
	nextRow := guardRow + guardDir.dRow
	nextCol := guardCol + guardDir.dCol

	for nextRow >= 0 && nextRow < numRows && nextCol >= 0 && nextCol < numCols {
		fmt.Printf("Next step is row=%d col=%d\n", nextRow, nextCol)

		if room[nextRow][nextCol] == '#' {
			fmt.Printf("Next step is a wall turning\n")
			//change direction
			dirIdx = (dirIdx + 1) % 4
			guardDir = dirs[dirIdx]

			nextRow = guardRow + guardDir.dRow
			nextCol = guardCol + guardDir.dCol
			fmt.Printf("Direction is now dRow=%d dCol=%d\n", guardDir.dRow, guardDir.dCol)
			fmt.Printf("New next step is row=%d col=%d\n", nextRow, nextCol)
			continue
		}

		guardRow = nextRow
		guardCol = nextCol
		nextRow = guardRow + guardDir.dRow
		nextCol = guardCol + guardDir.dCol
		if room[guardRow][guardCol] != 'X' {
			room[guardRow][guardCol] = 'X'
			spacesVisited++
		} else {
			//visting a space a second time == loop possible
			possibleLoops++
		}
		fmt.Printf("Took next step\n")
	}

	fmt.Printf("Guard visited %d spaces.\n", spacesVisited)
	fmt.Printf("Possible loops: %d\n", possibleLoops)

}

func PartTwo() {
	fmt.Println("Not implemented")
}
