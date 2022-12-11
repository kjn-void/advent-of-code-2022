package main

import "fmt"

type cleaningSection struct {
	start int
	end   int
}

type sectionAssignment struct {
	a cleaningSection
	b cleaningSection
}

func isRedundantAssignment(sa sectionAssignment) bool {
	return (sa.a.start <= sa.b.start && sa.a.end >= sa.b.end) ||
		(sa.b.start <= sa.a.start && sa.b.end >= sa.a.end)
}

func isOverlappingAssignment(sa sectionAssignment) bool {
	return (sa.a.start <= sa.b.start && sa.a.end >= sa.b.start) ||
		(sa.b.start <= sa.a.start && sa.b.end >= sa.a.start)
}

func countAssignmentsWhere(sas []sectionAssignment, predicate func(sectionAssignment) bool) int {
	cnt := 0
	for _, sa := range sas {
		if predicate(sa) {
			cnt++
		}
	}
	return cnt
}

func parseSectionAssignments(input []string) []sectionAssignment {
	sas := []sectionAssignment{}
	for _, row := range input {
		var sa sectionAssignment
		_, err := fmt.Sscanf(row, "%d-%d,%d-%d",
			&sa.a.start, &sa.a.end, &sa.b.start, &sa.b.end)
		if err != nil {
			panic("Invalid input: " + err.Error())
		}
		sas = append(sas, sa)
	}
	return sas
}

func day4(input []string) {
	assignments := parseSectionAssignments(input)
	fmt.Println(countAssignmentsWhere(assignments, isRedundantAssignment))
	fmt.Println(countAssignmentsWhere(assignments, isOverlappingAssignment))
}

func init() {
	Solutions[4] = day4
}
