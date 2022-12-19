// 1482

package main

import (
	"fmt"
	"runtime"
)

const (
	Ore RobotId = iota
	Clay
	Obsidian
	Geode
	RobotCnt
	Nothing = 255
)

const (
	PRODUCTION_TIME_LIMIT_PART1 = 24
	PRODUCTION_TIME_LIMIT_PART2 = 32
)

type RobotId int

type Blueprint struct {
	Id     int
	Robots [RobotCnt]Resources
	Peak   Resources
}

type Resources struct {
	Ore      uint8
	Clay     uint8
	Obsidian uint8
}

type Production struct {
	*Blueprint
	NumRobots     [RobotCnt]uint8
	NumOpenGeodes uint8
	Building      RobotId
	// Robot(s) that could have been built earlier but wasn't, no point in
	// building that type of robot before building something else.
	// This field is reset when the next robot is built.
	InhibitRobots uint8
	Inventory     Resources
}

func (production *Production) canBuild(requirements Resources) bool {
	return requirements.Ore <= production.Inventory.Ore &&
		requirements.Clay <= production.Inventory.Clay &&
		requirements.Obsidian <= production.Inventory.Obsidian
}

func (production *Production) shouldBuild(robot RobotId) bool {
	switch robot {
	case Ore:
		return production.InhibitRobots&(1<<Ore) == 0 &&
			production.Peak.Ore >= production.NumRobots[Ore]+1
	case Clay:
		return production.InhibitRobots&(1<<Clay) == 0 &&
			production.Peak.Clay >= production.NumRobots[Clay]+1
	case Obsidian:
		return production.InhibitRobots&(1<<Obsidian) == 0 &&
			production.Peak.Obsidian >= production.NumRobots[Obsidian]+1
	}
	return true
}

func (production *Production) mustBuild() bool {
	return production.canBuild(production.Robots[Ore]) &&
		production.canBuild(production.Robots[Clay]) &&
		production.canBuild(production.Robots[Obsidian]) &&
		production.canBuild(production.Robots[Geode])
}

func (production *Production) startBuilds() []Production {
	prods := []Production{}
	if !production.mustBuild() {
		prod := *production
		for robotId := RobotId(0); robotId < RobotCnt; robotId++ {
			if production.canBuild(production.Blueprint.Robots[robotId]) {
				prod.InhibitRobots |= 1 << robotId
			}
		}
		prods = append(prods, prod)
	}
	for i, robot := range production.Blueprint.Robots {
		robotId := RobotId(i)
		if production.canBuild(robot) && production.shouldBuild(robotId) {
			prod := *production
			prod.InhibitRobots = 0
			prod.Building = robotId
			prod.Inventory.Ore -= robot.Ore
			prod.Inventory.Clay -= robot.Clay
			prod.Inventory.Obsidian -= robot.Obsidian
			prods = append(prods, prod)
		}
	}
	return prods
}

func (production *Production) finishBuild() {
	if production.Building != Nothing {
		production.NumRobots[production.Building]++
		production.Building = Nothing
	}
}

func (production *Production) collect() {
	production.Inventory.Ore += production.NumRobots[Ore]
	production.Inventory.Clay += production.NumRobots[Clay]
	production.Inventory.Obsidian += production.NumRobots[Obsidian]
	production.NumOpenGeodes += production.NumRobots[Geode]
}

func (production *Production) tick(nextSteps *[]Production) {
	for _, prod := range production.startBuilds() {
		prod.collect()
		prod.finishBuild()
		*nextSteps = append(*nextSteps, prod)
	}
}

func doRound(prods []Production) []Production {
	nxt := make([]Production, 0, len(prods))
	for _, prod := range prods {
		if prod.shouldContinue() {
			prod.tick(&nxt)
		}
	}
	return nxt
}

func (production *Production) shouldContinue() bool {
	return production.Inventory.Ore <= production.Blueprint.Peak.Ore*4 &&
		production.Inventory.Clay <= production.Blueprint.Peak.Clay*4
}

func (blueprint Blueprint) findMaxOpenGeodes(timeLimit int) int {
	prods := []Production{{Blueprint: &blueprint, Building: Nothing, NumRobots: [4]uint8{1}}}
	time := 1
	for ; time <= timeLimit; time++ {
		prods = doRound(prods)
		if len(prods) > 8*runtime.NumCPU() {
			break
		}
	}
	maxOpenGeodes := uint8(0)
	workers := len(prods)
	maxOpenGeodesCh := make(chan uint8, runtime.NumCPU())
	// Fan-out on multiple CPU-cores
	for _, prod := range prods {
		go func(p Production) {
			ps := []Production{p}
			for t := time; t < timeLimit; t++ {
				ps = doRound(ps)
			}
			maxOG := uint8(0)
			for _, p := range ps {
				if maxOG < p.NumOpenGeodes {
					maxOG = p.NumOpenGeodes
				}
			}
			maxOpenGeodesCh <- maxOG
		}(prod)
	}
	for i := 0; i < workers; i++ {
		openGeodes := <-maxOpenGeodesCh
		if maxOpenGeodes < openGeodes {
			maxOpenGeodes = openGeodes
		}
	}

	return int(maxOpenGeodes)
}

func SumAllQualityLevels(blueprints []Blueprint) int {
	sum := 0
	for _, blueprint := range blueprints {
		sum += blueprint.Id * blueprint.findMaxOpenGeodes(PRODUCTION_TIME_LIMIT_PART1)
	}
	return sum
}

func FindOpenGeodesProduct(blueprints []Blueprint) int {
	product := 1
	for _, blueprint := range blueprints[:3] {
		product *= blueprint.findMaxOpenGeodes(PRODUCTION_TIME_LIMIT_PART2)
	}
	return product
}

func day19(input []string) {
	blueprints := parseBlueprints(input)
	fmt.Println(SumAllQualityLevels(blueprints))
	fmt.Println(FindOpenGeodesProduct(blueprints[:3]))
}

func init() {
	Solutions[19] = day19
}

func parseBlueprints(input []string) []Blueprint {
	blueprints := []Blueprint{}
	for _, desc := range input {
		blueprint := Blueprint{}
		if _, err := fmt.Sscanf(desc,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&blueprint.Id,
			&blueprint.Robots[Ore].Ore,
			&blueprint.Robots[Clay].Ore,
			&blueprint.Robots[Obsidian].Ore, &blueprint.Robots[Obsidian].Clay,
			&blueprint.Robots[Geode].Ore, &blueprint.Robots[Geode].Obsidian); err != nil {
			panic("Failed to parse blueprint: " + err.Error())
		}
		blueprint.Peak.Ore = maxResource(blueprint, func(r Resources) uint8 { return r.Ore })
		blueprint.Peak.Clay = maxResource(blueprint, func(r Resources) uint8 { return r.Clay })
		blueprint.Peak.Obsidian = maxResource(blueprint, func(r Resources) uint8 { return r.Obsidian })
		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}

func maxResource(blueprint Blueprint, getResource func(Resources) uint8) uint8 {
	max := uint8(0)
	for _, robot := range blueprint.Robots {
		prod := getResource(robot)
		if max < prod {
			max = prod
		}
	}
	return max
}
