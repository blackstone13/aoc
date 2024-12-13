package day09

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartOne() {
	// filename := "problems/day09/sample.txt"
	filename := "problems/day09/input.txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file '%s': %\n", filename, err)
	}

	input := strings.TrimSpace(string(file))
	disk := []int{}

	for i, char := range input {
		fmt.Printf("num char '%c'\n", char)
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Printf("error converting char '%c' to int: %v\n", char, err)
		}

		for block := 0; block < num; block++ {
			if i%2 == 0 {
				disk = append(disk, i/2)
			} else {
				disk = append(disk, -1) // -1 represents 'free' space
			}
		}
	}

	// fmt.Printf("Disk:\n%v\n", disk)
	start := 0
	end := len(disk) - 1

	for start < end-1 {
		// fmt.Printf("start=%d end=%d\n", start, end)
		for disk[start] != -1 {
			start++
		}
		// fmt.Printf("advanced start=%d\n", start)

		disk[start] = disk[end]
		disk[end] = -1
		// fmt.Printf("disk: %v\n", disk)

		end--
	}

	checksum := 0

	for i, data := range disk {
		if data == -1 {
			continue
		}

		checksum += i * data
	}

	// fmt.Printf("Defraged Disk:\n%v\n", disk)
	fmt.Printf("Defragged Checksum: %d\n", checksum)
}

func PartTwo() {
	// filename := "problems/day09/sample.txt"
	filename := "problems/day09/input.txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file '%s': %\n", filename, err)
	}

	input := strings.TrimSpace(string(file))
	disk := []int{}

	for i, char := range input {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Printf("error converting char '%c' to int: %v\n", char, err)
		}

		for block := 0; block < num; block++ {
			if i%2 == 0 {
				disk = append(disk, i/2)
			} else {
				disk = append(disk, -1) // -1 represents 'free' space
			}
		}
	}

	end := len(disk) - 1

	// Process blocks from the end towards the front
	for blockStart := end; blockStart > 0; {
		blockEnd := blockStart

		// Move blockEnd backwards over trailing free spaces
		for blockEnd >= 0 && disk[blockEnd] == -1 {
			blockEnd--
			blockStart--
		}

		// If we've run out of valid blocks
		if blockEnd < 0 || blockStart <= 0 {
			blockStart--
			continue
		}

		// Move blockStart to the beginning of the current file block
		// Make sure blockStart is not decremented below 0
		for blockStart > 0 && blockStart-1 >= 0 && disk[blockStart-1] == disk[blockEnd] {
			blockStart--
		}

		if blockStart < 0 {
			blockStart = 0
		}

		blockLen := blockEnd - blockStart + 1

		// Attempt to find a free space to the left where this block can fit
		freeStart := 0
		for freeStart < blockStart {
			// Advance freeStart until we hit a free space
			for freeStart < blockStart && freeStart < len(disk) && disk[freeStart] != -1 {
				freeStart++
			}
			if freeStart >= blockStart {
				// No free space found to the left
				break
			}

			// Expand freeEnd to cover a continuous run of free space
			freeEnd := freeStart
			for freeEnd+1 < blockStart && freeEnd+1 < len(disk) && disk[freeEnd+1] == -1 {
				freeEnd++
			}

			freeLen := freeEnd - freeStart + 1

			if freeLen >= blockLen {
				// Move the block into this free segment
				for i := 0; i < blockLen; i++ {
					disk[freeStart+i] = disk[blockStart+i]
					disk[blockStart+i] = -1
				}
				break // Stop looking for free segments once moved
			} else {
				// This free segment is too small, move on to the next one
				freeStart = freeEnd + 1
			}
		}

		// Move to the previous block
		blockStart--
	}

	checksum := 0

	for i, data := range disk {
		if data == -1 {
			continue
		}

		checksum += i * data
	}

	// fmt.Printf("Defraged Disk:\n%v\n", disk)
	fmt.Printf("Defragged Checksum: %d\n", checksum)
}
