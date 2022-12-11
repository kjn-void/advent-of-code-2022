package main

import (
	"fmt"
	"math/bits"
	"unicode"
)

type bitset uint64

type rucksack struct {
	compartments [2]bitset
}

func makeBitset(desc string) bitset {
	bset := bitset(0)
	for _, content := range desc {
		if unicode.IsLower(content) {
			bset |= (1 << (content - 'a'))
		} else {
			bset |= (1 << (content - 'A' + 26))
		}
	}
	return bset
}

func (b bitset) priority() int {
	return 1 + bits.TrailingZeros64(uint64(b))
}

func (r rucksack) priority() int {
	return (r.compartments[0] & r.compartments[1]).priority()
}

func (r rucksack) items() bitset {
	return r.compartments[0] | r.compartments[1]
}

func sumOfDuplicateItems(rs []rucksack) int {
	sum := 0
	for _, r := range rs {
		sum += r.priority()
	}
	return sum
}

func groupPriority(group []rucksack) int {
	grpItems := group[0].items()
	for _, rcksck := range group[1:] {
		grpItems &= rcksck.items()
	}
	return grpItems.priority()
}

func chunks(rs []rucksack, stride int) chan []rucksack {
	ch := make(chan []rucksack)
	go func() {
		for i := 0; i < len(rs); i += stride {
			chks := []rucksack{}
			for j := i; j < i+stride; j++ {
				chks = append(chks, rs[j])
			}
			ch <- chks
		}
		close(ch)
	}()
	return ch
}

func sumOfGroupItems(rs []rucksack) int {
	sum := 0
	for group := range chunks(rs, 3) {
		sum += groupPriority(group)
	}
	return sum
}

func parseRucksacks(input []string) []rucksack {
	rs := []rucksack{}
	for _, row := range input {
		mid := len(row) / 2
		a := makeBitset(row[:mid])
		b := makeBitset(row[mid:])
		rs = append(rs, rucksack{[2]bitset{a, b}})
	}
	return rs
}

func day3(input []string) {
	rs := parseRucksacks(input)
	fmt.Println(sumOfDuplicateItems(rs))
	fmt.Println(sumOfGroupItems(rs))
}

func init() {
	Solutions[3] = day3
}
