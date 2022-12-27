package main

import (
	"fmt"
	"math/bits"
	"runtime"
	"sort"
	"strings"
)

const (
	START_VALVE         = "AA"
	SOLO_TIME_LIMIT     = 30
	ELEPHANT_TIME_LIMIT = 26
	S_OFFSET            = 8
)

type Valve struct {
	Name        string
	Flow        int
	Connections []string
	Distances   []int
}

type ValveMapPath struct {
	*Valve
	Distance int
}

type ValvePath struct {
	Pressure int
	Elapsed  int
	Location int
	P        []string
	T        []int
}

type ValveQueue []ValveMapPath

type ValveSet uint32

type ValveSetPair struct {
	A ValveSet
	B ValveSet
}

func (s ValveSet) isSet(i int) bool {
	return s&(1<<(i+S_OFFSET)) != 0
}

func (s *ValveSet) set(i int) {
	*s |= 1 << (i + S_OFFSET)
}

func (s *ValveSet) at(i int) {
	*s &= ^((ValveSet(1) << S_OFFSET) - 1)
	*s |= ValveSet(i)
	s.set(i)
}

func (q *ValveQueue) dequeue() ValveMapPath {
	vp := (*q)[0]
	*q = (*q)[1:]
	return vp
}

func (q *ValveQueue) enqueue(vp ValveMapPath) {
	*q = append(*q, vp)
}

func (v Valve) pressureRelease(elapsed, limit int) int {
	if elapsed > limit {
		return 0
	}
	return v.Flow * (limit - elapsed)
}

func FindMaxPressureRelease(valves []*Valve, limit int, visited ValveSet) int {
	numValves := len(valves)
	numSets := numValves - bits.OnesCount32(uint32(visited)) - 1
	paths := map[ValveSet]ValvePath{}
	// Held Karp
	for to := 1; to < numValves; to++ {
		if visited.isSet(to) {
			continue
		}
		s := visited
		s.at(to)
		valve := valves[to]
		distance := valve.Distances[0]
		paths[s] = ValvePath{valve.pressureRelease(distance, limit), distance, to, []string{valves[to].Name}, []int{distance}}
	}

	for s := 1; s < numSets; s++ {
		sPaths := map[ValveSet]ValvePath{}
		for visited, path := range paths {
			from := valves[path.Location]
			for to := 1; to < numValves; to++ {
				if !visited.isSet(to) {
					newVisited := visited
					newVisited.at(to)
					elapsed := path.Elapsed + from.Distances[to]
					pressure := path.Pressure + valves[to].pressureRelease(elapsed, limit)
					newVp := ValvePath{pressure, elapsed, to, []string{}, []int{}}
					newVp.P = append(append([]string{}, path.P...), valves[to].Name)
					newVp.T = append(append([]int{}, path.T...), elapsed)
					if vp, found := sPaths[newVisited]; !found || vp.Pressure < newVp.Pressure {
						sPaths[newVisited] = newVp
					}
				}
			}
		}
		paths = sPaths
	}

	maxPressureRelease := 0
	for _, vp := range paths {
		if maxPressureRelease < vp.Pressure {
			maxPressureRelease = vp.Pressure
		}
	}
	return maxPressureRelease
}

func FindMaxPressureReleaseSolo(valves []*Valve) int {
	return FindMaxPressureRelease(valves, SOLO_TIME_LIMIT, ValveSet(0))
}

func combIt(n, p, l int, ch chan ValveSetPair, s ValveSet) {
	if s == ValveSet(0) {
		defer close(ch)
	}
	if p == 0 {
		p := ValveSetPair{}
		for i := 1; i < n; i++ {
			if s.isSet(i) {
				p.A.set(i)
			} else {
				p.B.set(i)
			}
		}
		ch <- p
		return
	}
	for i := l; i < n; i++ {
		newS := s
		newS.set(i)
		combIt(n, p-1, i+1, ch, newS)
	}
}

func comb(n, p int) chan ValveSetPair {
	ch := make(chan ValveSetPair)
	go combIt(n, p, 1, ch, ValveSet(0))
	return ch
}

func FindMaxPressureReleaseWithElephant(valves []*Valve) int {
	ch := make(chan int, runtime.NumCPU())
	for i := 0; i <= (len(valves)-1)/2; i++ {
		go func(ii int) {
			mpr := 0
			for p := range comb(len(valves), ii) {
				mp1 := FindMaxPressureRelease(valves, ELEPHANT_TIME_LIMIT, p.A)
				mp2 := FindMaxPressureRelease(valves, ELEPHANT_TIME_LIMIT, p.B)
				totalPressureRelease := mp1 + mp2
				if mpr < totalPressureRelease {
					mpr = totalPressureRelease
				}
			}
			ch <- mpr
		}(i)
	}
	maxPressureRelease := 0
	for i := 0; i <= (len(valves)-1)/2; i++ {
		mpr := <-ch
		if maxPressureRelease < mpr {
			maxPressureRelease = mpr
		}
	}
	return maxPressureRelease
}

func day16(input []string) {
	valves := parseValves(input)
	fmt.Println(FindMaxPressureReleaseSolo(valves))
	fmt.Println(FindMaxPressureReleaseWithElephant(valves))
}

func init() {
	Solutions[16] = day16
}

func parseValves(input []string) []*Valve {
	allValves := map[string]*Valve{}
	for _, desc := range input {
		valve := makeValve(desc)
		allValves[valve.Name] = &valve
	}
	valves := []*Valve{allValves[START_VALVE]}
	for _, valve := range allValves {
		if valve.Flow > 0 {
			valves = append(valves, valve)
		}
	}
	sort.Slice(valves, func(i, j int) bool { return valves[i].Name < valves[j].Name })
	for i := range valves {
		setDistances(valves, i, allValves)
	}
	return valves
}

func makeValve(desc string) Valve {
	valve := Valve{Connections: []string{}}
	parts := strings.Split(desc, "; ")
	if _, err := fmt.Sscanf(parts[0], "Valve %s has flow rate=%d", &valve.Name, &valve.Flow); err != nil {
		panic("Failed to parse valve " + err.Error())
	}
	parts = strings.Fields(parts[1])
	for _, neighbor := range parts[4:] {
		if neighbor[len(neighbor)-1] == ',' {
			neighbor = neighbor[:len(neighbor)-1]
		}
		valve.Connections = append(valve.Connections, neighbor)
	}
	return valve
}

func setDistances(valves []*Valve, i int, allValves map[string]*Valve) {
	valve := valves[i]
	valve.Distances = make([]int, len(valves))
	visited := map[string]bool{valve.Name: true}
	queue := ValveQueue{ValveMapPath{valve, 0}}
	for len(queue) > 0 {
		path := queue.dequeue()
		for _, conn := range path.Connections {
			if !visited[conn] {
				next := ValveMapPath{allValves[conn], path.Distance + 1}
				queue.enqueue(next)
				visited[conn] = true
				if next.Flow > 0 || next.Name == START_VALVE {
					// +1 as only case where the valve is open is considered in Held Karp
					valve.Distances[valveIndex(valves, next.Name)] = next.Distance + 1
				}
			}
		}
	}
}

func valveIndex(valves []*Valve, name string) int {
	for i, valve := range valves {
		if valve.Name == name {
			return i
		}
	}
	panic("No valve named " + name)
}
