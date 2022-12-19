package main

import (
	"fmt"
	"math/bits"
	"unicode"
)

type BitSet uint64

type Rucksack struct {
	Compartments [2]BitSet
}

func makeBitset(desc string) BitSet {
	bset := BitSet(0)
	for _, content := range desc {
		if unicode.IsLower(content) {
			bset |= (1 << (content - 'a'))
		} else {
			bset |= (1 << (content - 'A' + 26))
		}
	}
	return bset
}

func (b BitSet) priority() int {
	return 1 + bits.TrailingZeros64(uint64(b))
}

func (r Rucksack) priority() int {
	return (r.Compartments[0] & r.Compartments[1]).priority()
}

func (r Rucksack) items() BitSet {
	return r.Compartments[0] | r.Compartments[1]
}

func SumOfDuplicateItems(rs []Rucksack) int {
	sum := 0
	for _, r := range rs {
		sum += r.priority()
	}
	return sum
}

func groupPriority(group []Rucksack) int {
	grpItems := group[0].items()
	for _, rcksck := range group[1:] {
		grpItems &= rcksck.items()
	}
	return grpItems.priority()
}

func chunks(rs []Rucksack, stride int) chan []Rucksack {
	ch := make(chan []Rucksack)
	go func() {
		for i := 0; i < len(rs); i += stride {
			chks := []Rucksack{}
			for j := i; j < i+stride; j++ {
				chks = append(chks, rs[j])
			}
			ch <- chks
		}
		close(ch)
	}()
	return ch
}

func SumOfGroupItems(rs []Rucksack) int {
	sum := 0
	for group := range chunks(rs, 3) {
		sum += groupPriority(group)
	}
	return sum
}

func parseRucksacks(input []string) []Rucksack {
	rs := []Rucksack{}
	for _, row := range input {
		mid := len(row) / 2
		a := makeBitset(row[:mid])
		b := makeBitset(row[mid:])
		rs = append(rs, Rucksack{[2]BitSet{a, b}})
	}
	return rs
}

func day3(input []string) {
	rs := parseRucksacks(input)
	fmt.Println(SumOfDuplicateItems(rs))
	fmt.Println(SumOfGroupItems(rs))
}

func init() {
	Solutions[3] = day3
}
