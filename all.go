package main

import (
	"aoc/problems/day01"
	"aoc/problems/day02"
	"aoc/problems/day03"
	"aoc/problems/day04"
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
}
