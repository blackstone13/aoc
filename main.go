package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Command Line Flags & Parsing
	var day *int = flag.Int("day", 1, "The Advent Day to execute.")
	var part *int = flag.Int("part", 1, "The Problem Part to execute.")
	flag.Parse()

	registry := NewProblemRegistry()

	selected, ok := registry.GetProblem(*day, *part)
	if !ok {
		fmt.Printf("No problem registered for Day %d - Part %d\n", *day, *part)
		os.Exit(1)
	}

	selected()
}
