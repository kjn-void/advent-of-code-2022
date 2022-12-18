package main

import (
	"fmt"
	"math"
)

type SensorPos struct {
	X, Y int
}

type Sensor struct {
	Pos    SensorPos
	Radius int
}

type SensorMap struct {
	Sensors []Sensor
	Bacons  map[SensorPos]bool
	MinX    int
	MaxX    int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (pt SensorPos) manhattanDistance(otherPt SensorPos) int {
	return abs(pt.X-otherPt.X) + abs(pt.Y-otherPt.Y)
}

func (sensorMap SensorMap) isCovered(pos SensorPos, ignoreBacons bool) bool {
	for _, sensor := range sensorMap.Sensors {
		if sensor.Pos.manhattanDistance(pos) <= sensor.Radius {
			return ignoreBacons || !sensorMap.Bacons[pos]
		}
	}
	return false
}

func (sensorMap SensorMap) CoveredPositions(row int) int {
	cnt := 0
	for x := sensorMap.MinX; x <= sensorMap.MaxX; x++ {
		if sensorMap.isCovered(SensorPos{x, row}, false) {
			cnt++
		}
	}
	return cnt
}

func (sensorMap SensorMap) checkSensorPerimeter(sensor *Sensor, missingPosCh chan SensorPos, maxCoord int) {
	pos := SensorPos{sensor.Pos.X, sensor.Pos.Y + sensor.Radius + 1}
	for i, d := range [4]SensorPos{{1, -1}, {-1, -1}, {-1, 1}, {1, 1}} {
		for {
			if pos.X >= 0 && pos.X <= maxCoord && pos.Y >= 0 && pos.Y <= maxCoord && !sensorMap.isCovered(pos, true) {
				missingPosCh <- pos
				return
			}
			pos.X += d.X
			pos.Y += d.Y
			if ((i == 0 || i == 2) && pos.Y == sensor.Pos.Y) ||
				((i == 1 || i == 3) && pos.X == sensor.Pos.X) {
				break
			}
		}
	}
}

func (sensorMap SensorMap) TuningFrequency(maxCoord int) uint64 {
	numSensors := len(sensorMap.Sensors)
	doneCh := make(chan bool)
	missingPosCh := make(chan SensorPos, numSensors)

	for i := 0; i < numSensors; i++ {
		go func(sensorId int) {
			sensorMap.checkSensorPerimeter(&sensorMap.Sensors[sensorId], missingPosCh, maxCoord)
			doneCh <- true
		}(i)
	}

	for i := 0; i < numSensors; i++ {
		<-doneCh
	}

	missingPos := <-missingPosCh
	return 4000000*uint64(missingPos.X) + uint64(missingPos.Y)
}

func day15(input []string) {
	sensorMap := parseSensorMap(input)
	fmt.Println(sensorMap.CoveredPositions(2000000))
	fmt.Println(sensorMap.TuningFrequency(4000000))
}

func init() {
	Solutions[15] = day15
}

func parseSensorMap(input []string) SensorMap {
	sensorMap := SensorMap{[]Sensor{}, map[SensorPos]bool{}, math.MaxInt, math.MinInt}
	for _, desc := range input {
		bacon := SensorPos{}
		sensor := Sensor{}
		if _, err := fmt.Sscanf(desc, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.Pos.X, &sensor.Pos.Y, &bacon.X, &bacon.Y); err != nil {
			panic("Failed to parse sensor: " + err.Error())
		}
		sensor.Radius = sensor.Pos.manhattanDistance(bacon)
		sensorMap.Sensors = append(sensorMap.Sensors, sensor)
		sensorMap.Bacons[bacon] = true

		if sensorMap.MinX > sensor.Pos.X-sensor.Radius {
			sensorMap.MinX = sensor.Pos.X - sensor.Radius
		}
		if sensorMap.MaxX < sensor.Pos.X+sensor.Radius {
			sensorMap.MaxX = sensor.Pos.X + sensor.Radius
		}
	}
	return sensorMap
}
