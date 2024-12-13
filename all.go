package main

import (
	"aoc/problems/day01"
	"aoc/problems/day02"
	"aoc/problems/day03"
	"aoc/problems/day04"
	"aoc/problems/day05"
	"aoc/problems/day06"
	"aoc/problems/day07"
	"aoc/problems/day08"
	"aoc/problems/day09"
	"aoc/problems/day10"
	"aoc/problems/day11"
	"aoc/problems/day12"
)

func (r ProblemRegistry) RegisterAll() {
	r.RegisterProblem(1, 1, day01.PartOne)
	r.RegisterProblem(1, 2, day01.PartTwo)
	r.RegisterProblem(2, 1, day02.PartOne)
	r.RegisterProblem(2, 2, day02.PartTwo)
	r.RegisterProblem(3, 1, day03.PartOne)
	r.RegisterProblem(3, 2, day03.PartTwo)
	r.RegisterProblem(4, 1, day04.PartOne)
	r.RegisterProblem(4, 2, day04.PartTwo)
	r.RegisterProblem(5, 1, day05.PartOne)
	r.RegisterProblem(5, 2, day05.PartTwo)
	r.RegisterProblem(6, 1, day06.PartOne)
	r.RegisterProblem(6, 2, day06.PartTwo)
	r.RegisterProblem(7, 1, day07.PartOne)
	r.RegisterProblem(7, 2, day07.PartTwo)
	r.RegisterProblem(8, 1, day08.PartOne)
	r.RegisterProblem(9, 1, day09.PartOne)
	r.RegisterProblem(9, 2, day09.PartTwo)
	r.RegisterProblem(10, 1, day10.PartOne)
	r.RegisterProblem(10, 2, day10.PartTwo)
	r.RegisterProblem(11, 1, day11.PartOne)
	r.RegisterProblem(11, 2, day11.PartTwo)
	r.RegisterProblem(12, 1, day12.PartOne)
	r.RegisterProblem(12, 2, day12.PartTwo)
}
