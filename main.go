package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solution func(input []string)

// Create 26 days to allow for 1 based indexing
var Solutions []Solution = make([]Solution, 26)

func parseDays(args []string) []int {
	days := make([]int, 0)
	for _, arg := range args {
		day, err := strconv.Atoi(arg)
		if err != nil || day < 1 || day > 25 {
			fmt.Printf("Day argument is not a valid, must be an integer between 1-25, is %v\n", arg)
			os.Exit(1)
		}
		days = append(days, day)
	}
	return days
}

func inputAsString(day int) []string {
	content, err := os.ReadFile(fmt.Sprintf("inputs/day%v.txt", day))
	if err != nil {
		fmt.Printf("Failed to read input for day %v: %v", day, err)
		os.Exit(1)
	}
	input := strings.Split(string(content), "\n")
	if input[len(input)-1] == "" {
		return input[:len(input)-1]
	}
	return input
}

func main() {
	flag.Parse()
	days := parseDays(flag.Args())
	for _, day := range days {
		fmt.Printf("# Day %v\n", day)
		Solutions[day](inputAsString(day))
	}
}
