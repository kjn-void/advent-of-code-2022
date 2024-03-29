package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Calories []int
type SortedElfTrain []Calories

func (et SortedElfTrain) Len() int      { return len(et) }
func (et SortedElfTrain) Swap(i, j int) { et[i], et[j] = et[j], et[i] }

func (cals Calories) totalCalories() int {
	totalElfCalories := 0
	for _, cal := range cals {
		totalElfCalories += cal
	}
	return totalElfCalories
}

func (et SortedElfTrain) CarryingMostCalories() int {
	return et[0].totalCalories()
}

func (et SortedElfTrain) TopThreeCarrying() int {
	if len(et) < 3 {
		panic("Cannot calculate top three")
	}
	return et[0].totalCalories() + et[1].totalCalories() + et[2].totalCalories()
}

func parseElfTrain(input []string) SortedElfTrain {
	cals := Calories{}
	elfs := SortedElfTrain{}
	for i, cal := range input {
		if len(cal) > 0 {
			if c, err := strconv.Atoi(cal); err == nil {
				cals = append(cals, c)
			} else {
				panic("Failed to parse calories: " + err.Error())
			}
		}
		if len(cal) == 0 || i == len(input)-1 {
			elfs = append(elfs, cals)
			cals = Calories{}
		}
	}
	sort.Slice(elfs, func(i, j int) bool { return elfs[i].totalCalories() > elfs[j].totalCalories() })
	return elfs
}

func day1(input []string) {
	elfs := parseElfTrain(input)
	fmt.Println(elfs.CarryingMostCalories())
	fmt.Println(elfs.TopThreeCarrying())
}

func init() {
	Solutions[1] = day1
}
