package main

import "testing"

var input4 = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestDay4_1(t *testing.T) {
	assignments := parseSectionAssignments(input4)
	cnt := CountAssignmentsWhere(assignments, isRedundantAssignment)
	if cnt != 2 {
		t.Fatalf("Got %v redundant assignments, expected 2", cnt)
	}
}

func TestDay4_2(t *testing.T) {
	assignments := parseSectionAssignments(input4)
	cnt := CountAssignmentsWhere(assignments, isOverlappingAssignment)
	if cnt != 4 {
		t.Fatalf("Got %v redundant assignments, expected 4", cnt)
	}
}

func BenchmarkDay4_parse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseSectionAssignments(inputAsString(4))
	}
}

func BenchmarkDay4_part1(b *testing.B) {
	assignments := parseSectionAssignments(inputAsString(4))
	for i := 0; i < b.N; i++ {
		CountAssignmentsWhere(assignments, isRedundantAssignment)
	}
}

func BenchmarkDay4_part2(b *testing.B) {
	assignments := parseSectionAssignments(inputAsString(4))
	for i := 0; i < b.N; i++ {
		CountAssignmentsWhere(assignments, isOverlappingAssignment)
	}
}
