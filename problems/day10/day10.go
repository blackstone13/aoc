package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pos struct {
	row, col int
}

func PartOne() {
	// filename := "problems/day10/sample.txt"
	filename := "problems/day10/input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file '%s': %v\n", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	topo := make([][]int, 0)
	trailheads := make([]Pos, 0)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		topo = append(topo, make([]int, 0))
		for col, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Printf("error converting char '%c' to int: %v\n", char, err)
			}
			topo[row] = append(topo[row], num)
			if num == 0 {
				pos := Pos{row: row, col: col}
				trailheads = append(trailheads, pos)
			}
		}
		row++
	}

	totalTrails := 0
	distinctTrails := 0

	for _, th := range trailheads {
		peaks := make(map[Pos]int)
		distinctTrails += dfsTrails(th, topo, peaks)
		totalTrails += len(peaks)
	}

	// fmt.Printf("Trailheads:\n%v\n", trailheads)
	fmt.Printf("Trails: %d\n", totalTrails)
	fmt.Printf("Distint Trails: %d\n", distinctTrails)
}

func dfsTrails(pos Pos, topo [][]int, peaks map[Pos]int) int {
	elev := topo[pos.row][pos.col]
	if elev == 9 {
		peaks[pos] = 1
		return 1
	}

	rows := len(topo)
	cols := len(topo[0])
	count := 0

	// check right
	right := Pos{pos.row, pos.col + 1}
	if right.col < cols && topo[right.row][right.col] == elev+1 {
		count += dfsTrails(right, topo, peaks)
	}

	left := Pos{pos.row, pos.col - 1}
	if left.col >= 0 && topo[left.row][left.col] == elev+1 {
		count += dfsTrails(left, topo, peaks)
	}

	up := Pos{pos.row - 1, pos.col}
	if up.row >= 0 && topo[up.row][up.col] == elev+1 {
		count += dfsTrails(up, topo, peaks)
	}

	down := Pos{pos.row + 1, pos.col}
	if down.row < rows && topo[down.row][down.col] == elev+1 {
		count += dfsTrails(down, topo, peaks)
	}
	return count
}

func PartTwo() {
	fmt.Println("Not Implemented.")
}
