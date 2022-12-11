package main

import (
	"fmt"
	"sort"
	"strconv"
)

type calories []int
type sortedElfTrain []calories

func (et sortedElfTrain) Len() int      { return len(et) }
func (et sortedElfTrain) Swap(i, j int) { et[i], et[j] = et[j], et[i] }

func (cals calories) totalCalories() int {
	totalElfCalories := 0
	for _, cal := range cals {
		totalElfCalories += cal
	}
	return totalElfCalories
}

func (et sortedElfTrain) carryingMostCalories() int {
	return et[0].totalCalories()
}

func (et sortedElfTrain) topThreeCarrying() int {
	if len(et) < 3 {
		panic("Cannot calculate top three")
	}
	return et[0].totalCalories() + et[1].totalCalories() + et[2].totalCalories()
}

func parseElfTrain(input []string) sortedElfTrain {
	cals := calories{}
	elfs := sortedElfTrain{}
	for _, cal := range input {
		if len(cal) == 0 {
			elfs = append(elfs, cals)
			cals = calories{}
		} else {
			if c, err := strconv.Atoi(cal); err == nil {
				cals = append(cals, c)
			} else {
				panic("Failed to parse calories: " + err.Error())
			}
		}
	}
	sort.Slice(elfs, func(i, j int) bool { return elfs[i].totalCalories() > elfs[j].totalCalories() })
	return elfs
}

func day1(input []string) {
	elfs := parseElfTrain(input)
	fmt.Println(elfs.carryingMostCalories())
	fmt.Println(elfs.topThreeCarrying())
}

func init() {
	Solutions[1] = day1
}
