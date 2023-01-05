package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


type Blueprint struct {
	num       int
	ore_cost []int
	clay_cost []int
	obsidian_cost []int
	geode_cost     []int
}


type State struct {
  ore, clay, obsidian, geode, r1, r2, r3, r4, t int 
}

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func Min(nums ...int) int {
  min := nums[0]
  for _, num := range nums {
    if num < min {
      min = num
    }
  }
  return min
}

func Max(nums ...int) int {
  max := nums[0]
  for _, num := range nums {
    if num > max {
      max = num
    }
  }
  return max
}

func convStr(lst []int) []string {
  var res []string
  for _, el := range lst {
    res = append(res, strconv.Itoa(el))
  }
  return res
}

func hashVal(state State) string {
  r1 := state.r1
  r2 := state.r2
  r3 := state.r3
  r4 := state.r4
  t := state.t
  o := state.ore
  c := state.clay
  ob := state.obsidian
  g := state.geode

  return strings.Join(convStr([]int{o, c, ob, g, r1, r2, r3, r4, t}), "_")
}


func parseLine(input string) blueprint {
	r, _ := regexp.Compile(
		"^Blueprint ([0-9]+): Each ore robot costs ([0-9]+) ore. Each clay robot costs ([0-9]+) ore. Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay. Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian.$",
	)
	matches := r.FindStringSubmatch(input)
	return blueprint{convInt(matches[1]), convInt(matches[2]), convInt(matches[3]), convInt(matches[4]), convInt(matches[5]), convInt(matches[6]), convInt(matches[7])}
}

func part1() int {
	dat, _ := os.ReadFile("./day19-input")
	blueprints := parseInput(string(dat))

	sum := 0
	for _, bp := range blueprints {
		st := newState(bp)
		geodesMade := st.calcMostGeodes(0, map[string]int{}, 24, 24)
		sum += bp.id * geodesMade
	}

	// total quality of all blueprints, quality = id * (# geodes in 24 min)
	return sum
}


type blueprint struct {
	id                                        int
	oreForOreRobot                            int
	oreForClayRobot                           int
	oreForObsidianRobot, clayForObsidianRobot int
	oreForGeodeRobot, obsidianForGeodeRobot   int
}

type state struct {
	blueprint
	ore, clay, obsidian, geode                         int
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
}

func newState(blueprint blueprint) state {
	return state{
		blueprint: blueprint,
		oreRobots: 1,
	}
}

func (s *state) farm() {
	s.ore += s.oreRobots
	s.clay += s.clayRobots
	s.obsidian += s.obsidianRobots
	s.geode += s.geodeRobots
}

func (s *state) hash(time int) string {
	return fmt.Sprint(time, s.ore, s.clay, s.obsidian,
		s.geode, s.oreRobots, s.clayRobots, s.obsidianRobots, s.geodeRobots)
}

func (s state) copy() state {
	return s
}


func (s *state) calcMostGeodes(time int, memo map[string]int, totalTime int, earliestGeode int) int {
	if time == totalTime {
		return s.geode
	}

	h := s.hash(time)
	if v, ok := memo[h]; ok {
		return v
	}

	if s.geode == 0 && time > earliestGeode {
		return 0
	}

	mostGeodes := s.geode

	if s.ore >= s.oreForGeodeRobot &&
		s.obsidian >= s.obsidianForGeodeRobot {
		cp := s.copy()

		cp.farm()

		cp.ore -= cp.oreForGeodeRobot
		cp.obsidian -= cp.obsidianForGeodeRobot
		cp.geodeRobots++
		if cp.geodeRobots == 1 {
			earliestGeode = Min(earliestGeode, time+1)
		}
		mostGeodes = Max(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))

		memo[h] = mostGeodes
		return mostGeodes
	}

	if time <= totalTime-16 &&
		s.oreRobots < s.oreForObsidianRobot*2 &&
		s.ore >= s.oreForOreRobot {
		cp := s.copy()
		cp.ore -= cp.oreForOreRobot

		cp.farm()

		cp.oreRobots++
		mostGeodes = Max(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))
	}
	if time <= totalTime-8 &&
		s.clayRobots < s.clayForObsidianRobot &&
		s.ore >= s.oreForClayRobot {
		cp := s.copy()
		cp.ore -= cp.oreForClayRobot

		cp.farm()

		cp.clayRobots++
		mostGeodes = Max(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))
	}
	if time <= totalTime-4 &&
		s.obsidianRobots < s.obsidianForGeodeRobot &&
		s.ore >= s.oreForObsidianRobot && s.clay >= s.clayForObsidianRobot {

		cp := s.copy()
		cp.ore -= cp.oreForObsidianRobot
		cp.clay -= cp.clayForObsidianRobot
		cp.farm()

		cp.obsidianRobots++
		mostGeodes = Max(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))
	}

	cp := s.copy()
	cp.ore += cp.oreRobots
	cp.clay += cp.clayRobots
	cp.obsidian += cp.obsidianRobots
	cp.geode += cp.geodeRobots
	mostGeodes = Max(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))

	memo[h] = mostGeodes
	return mostGeodes
}

func parseInput(input string) (ans []blueprint) {
	for _, line := range strings.Split(input, "\n") {
    if len(line) == 0 {
      continue
    }
		bp := parseLine(line)
    ans = append(ans, bp)
	}
	return ans
}
