package main

import "fmt"

type CleaningSection struct {
	Start int
	End   int
}

type SectionAssignment struct {
	A CleaningSection
	B CleaningSection
}

func isRedundantAssignment(sa SectionAssignment) bool {
	return (sa.A.Start <= sa.B.Start && sa.A.End >= sa.B.End) ||
		(sa.B.Start <= sa.A.Start && sa.B.End >= sa.A.End)
}

func isOverlappingAssignment(sa SectionAssignment) bool {
	return (sa.A.Start <= sa.B.Start && sa.A.End >= sa.B.Start) ||
		(sa.B.Start <= sa.A.Start && sa.B.End >= sa.A.Start)
}

func CountAssignmentsWhere(sas []SectionAssignment, predicate func(SectionAssignment) bool) int {
	cnt := 0
	for _, sa := range sas {
		if predicate(sa) {
			cnt++
		}
	}
	return cnt
}

func parseSectionAssignments(input []string) []SectionAssignment {
	sas := []SectionAssignment{}
	for _, row := range input {
		var sa SectionAssignment
		_, err := fmt.Sscanf(row, "%d-%d,%d-%d",
			&sa.A.Start, &sa.A.End, &sa.B.Start, &sa.B.End)
		if err != nil {
			panic("Invalid input: " + err.Error())
		}
		sas = append(sas, sa)
	}
	return sas
}

func day4(input []string) {
	assignments := parseSectionAssignments(input)
	fmt.Println(CountAssignmentsWhere(assignments, isRedundantAssignment))
	fmt.Println(CountAssignmentsWhere(assignments, isOverlappingAssignment))
}

func init() {
	Solutions[4] = day4
}
