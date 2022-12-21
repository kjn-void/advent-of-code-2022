package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	RootName = "root"
	MyName   = "humn"
)

type MonkeyName string

type MonkeyMatch map[MonkeyName]func(MonkeyMatch) int64

func RootMonkeyYelling(mm MonkeyMatch) int64 {
	return mm[RootName](mm)
}

func rootMatcher(mm MonkeyMatch, result chan int64, rootArgs string) {
	p := strings.Fields(rootArgs)
	if p[1] != "+" {
		panic("Fix additional operations for root monkey..." + p[1])
	}
	nameA := MonkeyName(p[0])
	nameB := MonkeyName(p[2])
	for {
		a := mm[nameA](mm)
		b := mm[nameB](mm)
		result <- a - b
		if a == b {
			break
		}
	}
}

func isPositive(n int64) bool {
	return n >= 0
}

func MyNumberForRootMonkeyMatch(mm MonkeyMatch, rootArgs string) int64 {
	result := make(chan int64)
	iYell := make(chan int64)
	go rootMatcher(mm, result, rootArgs)
	mm[MyName] = func(MonkeyMatch) int64 { return <-iYell }
	// Find range, xB..xA where the root monkey function change sign
	iYell <- 0
	xZeroSign := isPositive(<-result)
	aSign := xZeroSign
	xA := int64(-1)
	xB := int64(1)
	for growA := true; ; growA = !growA {
		if growA {
			xA *= 2
			iYell <- int64(xA)
		} else {
			xB *= 2
			iYell <- int64(xB)
		}
		diff := <-result
		if isPositive(diff) != xZeroSign {
			if growA {
				aSign = !xZeroSign
				xB = xA / 2
			} else {
				xA = xB / 2
			}
			break
		}
	}
	// Answer is in range xB..xA
	for {
		n := (xB-xA)/2 + xA
		iYell <- n
		diff := <-result
		if diff == 0 {
			return n
		}
		if isPositive(diff) == aSign {
			xA = n + 1
		} else {
			xB = n - 1
		}
	}
}

func day21(input []string) {
	mm, rootArgs := parseMonkeyMatch(input)
	fmt.Println(RootMonkeyYelling(mm))
	fmt.Println(MyNumberForRootMonkeyMatch(mm, rootArgs))
}

func init() {
	Solutions[21] = day21
}

func parseMonkeyMatch(input []string) (MonkeyMatch, string) {
	mm := MonkeyMatch{}
	rootArgs := ""
	for _, monkey := range input {
		var names [2]MonkeyName
		var op string
		parts := strings.Split(monkey, ": ")
		monkeyName := MonkeyName(parts[0])
		if number, err := strconv.Atoi(parts[1]); err == nil {
			mm[monkeyName] = func(mm MonkeyMatch) int64 { return int64(number) }
		} else if _, err := fmt.Sscanf(parts[1], "%s %s %s", &names[0], &op, &names[1]); err == nil {
			switch op {
			case "*":
				mm[monkeyName] = func(mm MonkeyMatch) int64 { return mm[names[0]](mm) * mm[names[1]](mm) }
			case "/":
				mm[monkeyName] = func(mm MonkeyMatch) int64 { return mm[names[0]](mm) / mm[names[1]](mm) }
			case "+":
				mm[monkeyName] = func(mm MonkeyMatch) int64 { return mm[names[0]](mm) + mm[names[1]](mm) }
			case "-":
				mm[monkeyName] = func(mm MonkeyMatch) int64 { return mm[names[0]](mm) - mm[names[1]](mm) }
			}
		} else {
			panic("Failed to parse monkey match line: " + monkey)
		}
		if monkeyName == RootName {
			rootArgs = parts[1]
		}
	}
	return mm, rootArgs
}
