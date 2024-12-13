package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	Result  int
	Numbers []int
}

func PartOne() {
	// filename := "problems/day07/sample.txt"
	filename := "problems/day07/input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file '%s': %v\n", filename, err)
		os.Exit(1)
	}

	equations := []Equation{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		eqParts := strings.Split(line, ":")
		resStr := eqParts[0]
		numStrs := strings.Fields(eqParts[1])

		res, err := strconv.Atoi(resStr)
		if err != nil {
			fmt.Printf("error converting result '%s' to int. %v", resStr, err)
			os.Exit(1)
		}

		nums := []int{}
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("error converting num '%s' to int. %v", numStr, err)
				os.Exit(1)
			}
			nums = append(nums, num)
		}
		eq := Equation{Result: res, Numbers: nums}
		equations = append(equations, eq)
	}

	solvableTotals := 0
	for _, eq := range equations {
		if solvableDFS(eq.Result, eq.Numbers) {
			solvableTotals += eq.Result
		}
	}

	fmt.Printf("Solvable Totals: %d\n", solvableTotals)
}

func solvableDFS(result int, nums []int) bool {
	if len(nums) == 1 {
		return result == nums[0]
	}

	added := nums[0] + nums[1]
	addedNums := append([]int{added}, nums[2:]...)
	if solvableDFS(result, addedNums) {
		return true
	}

	multed := nums[0] * nums[1]
	multedNums := append([]int{multed}, nums[2:]...)
	if solvableDFS(result, multedNums) {
		return true
	}

	numStr0 := strconv.Itoa(nums[0])
	numStr1 := strconv.Itoa(nums[1])
	comboNumStr := numStr0 + numStr1
	combo, err := strconv.Atoi(comboNumStr)
	if err != nil {
		fmt.Printf("error converting combo str '%s' to int: %v", comboNumStr, err)
		os.Exit(1)
	}
	comboNums := append([]int{combo}, nums[2:]...)
	if solvableDFS(result, comboNums) {
		return true
	}

	return false

}

func PartTwo() {
	fmt.Println("Not Implemented")
}
