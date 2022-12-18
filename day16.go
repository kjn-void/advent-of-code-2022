package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	START_VALVE         = "AA"
	SOLO_TIME_LIMIT     = 30
	ELEPHANT_TIME_LIMIT = 26
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
}

type ValvePathWithElephant struct {
	ValvePath
	IsMeAtLocation    bool
	MovingTowards     int
	DistanceRemaining int
	E                 [ELEPHANT_TIME_LIMIT]string
	M                 [ELEPHANT_TIME_LIMIT]string
	P                 [ELEPHANT_TIME_LIMIT]int
	D                 [ELEPHANT_TIME_LIMIT]int
}

type ValveQueue []ValveMapPath

type ValveSet uint32

type ValveVisited map[ValveSet]int

func (s ValveSet) isSet(i int) bool {
	return s&(1<<i) != 0
}

func (s ValveSet) getMoveAt() int {
	return int(s >> 24)
}

func (s *ValveSet) set(i int) {
	*s |= 1 << i
}

func (s *ValveSet) moveAt(idx int) {
	*s &= ValveSet((idx << 24) - 1)
	*s |= ValveSet(idx << 24)
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

func FindMaxPressureRelease(valves []*Valve) int {
	numValves := len(valves)
	paths := map[ValveSet]ValvePath{}
	// Held Karp
	for to := 1; to < numValves; to++ {
		s := ValveSet(1)
		s.set(to)
		valve := valves[to]
		distance := valve.Distances[0]
		paths[s] = ValvePath{valve.pressureRelease(distance, SOLO_TIME_LIMIT), distance, to}
	}

	for s := 1; s < numValves-1; s++ {
		sPaths := map[ValveSet]ValvePath{}
		for visited, path := range paths {
			from := valves[path.Location]
			for to := 1; to < numValves; to++ {
				if !visited.isSet(to) {
					newVisited := visited
					newVisited.set(to)
					elapsed := path.Elapsed + from.Distances[to]
					pressure := path.Pressure + valves[to].pressureRelease(elapsed, SOLO_TIME_LIMIT)
					newVp := ValvePath{pressure, elapsed, to}
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

func makePathWithElephant(valves []*Valve, oldPath ValvePathWithElephant, to int, s ValveSet) (ValvePathWithElephant, ValveSet) {
	path := oldPath
	from := oldPath.Location
	distance := valves[from].Distances[to]

	if distance > path.DistanceRemaining {
		path.Elapsed += path.DistanceRemaining
		path.DistanceRemaining = distance - path.DistanceRemaining
		path.IsMeAtLocation = !path.IsMeAtLocation
		path.Location = path.MovingTowards
		path.MovingTowards = to
	} else {
		path.DistanceRemaining -= distance
		path.Elapsed += distance
		path.Location = to
	}

	dP := valves[path.Location].pressureRelease(path.Elapsed, ELEPHANT_TIME_LIMIT)
	path.Pressure += dP
	path.P[path.Elapsed] = path.Pressure
	path.D[path.Elapsed] = dP
	if path.IsMeAtLocation {
		path.M[path.Elapsed] = valves[path.Location].Name
	} else {
		path.E[path.Elapsed] = valves[path.Location].Name
	}

	s.set(path.Location)
	s.moveAt(path.MovingTowards)

	return path, s
}

func FindMaxPressureReleaseWithElephant(valves []*Valve) int {
	numValves := len(valves)
	paths := make([]map[ValveSet]ValvePathWithElephant, numValves)
	for i := range paths {
		paths[i] = map[ValveSet]ValvePathWithElephant{}
	}
	// Held Karp
	start := ValvePathWithElephant{IsMeAtLocation: true}
	for i := 0; i < len(start.E); i++ {
		start.E[i] = ".."
		start.M[i] = ".."
	}
	start.M[0] = "AA"
	start.E[0] = "AA"

	for mTo := 1; mTo < numValves-1; mTo++ {
		for elephantTo := mTo + 1; elephantTo < numValves; elephantTo++ {
			mDist := valves[mTo].Distances[0]
			eDist := valves[elephantTo].Distances[0]
			visited := ValveSet(1)
			to := mTo
			start.IsMeAtLocation = mDist <= eDist
			if start.IsMeAtLocation {
				visited.moveAt(elephantTo)
			} else {
				to = elephantTo
				visited.moveAt(mTo)
			}
			path, newS := makePathWithElephant(valves, start, to, visited)
			paths[0][newS] = path
		}
	}

	for s := 0; s < numValves-1; s++ {
		for visited, path := range paths[s] {
			for to := 1; to < numValves; to++ {
				if !visited.isSet(to) && visited.getMoveAt() != to {
					newPath, newVisited := makePathWithElephant(valves, path, to, visited)
					if vp, found := paths[s+1][newVisited]; !found || vp.Pressure < newPath.Pressure {
						paths[s+1][newVisited] = newPath
					}
				}
			}
		}
	}

	maxPressureRelease := 0
	for _, vp := range paths[len(paths)-1] {
		if maxPressureRelease < vp.Pressure {
			maxPressureRelease = vp.Pressure
			fmt.Println(maxPressureRelease)
			fmt.Println(vp.M)
			fmt.Println(vp.E)
			fmt.Println(vp.D)
			fmt.Println(vp.P)
		}
	}
	return maxPressureRelease
}

func day16(input []string) {
	valves := parseValves(input)
	fmt.Println(FindMaxPressureRelease(valves))
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
