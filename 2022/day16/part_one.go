package main

import (
	_ "embed"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func parseLine(s string) Valve {
	r, _ := regexp.Compile("^Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? (.+)$")

	matches := r.FindStringSubmatch(s)
	return Valve{matches[1], convInt(matches[2]), strings.Split(matches[3], ", ")}
}

type Valve struct {
  label string
  rate int
  connections []string
}

type State struct {
  valve int
  open BitSet 
  pressure int
  time int
}

type Stack []State

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

type BitSet uint64

func (b BitSet) WithSet(bit int) BitSet {
	return b | (1 << bit)
}

func (b BitSet) IsEmpty(bit int) bool {
	return b&(1<<bit) == 0
}

func (s *Stack) Pop() State {
  last := len(*s) - 1
  state := (*s)[last]
  *s = (*s)[:last]
  return state
}

func (s *Stack) Push(state State) {
  *s = append(*s, state)
}

func findShortestDistances(valves []Valve) [][]int {
	distances := make([][]int, len(valves))
	for i := range distances {
		distances[i] = make([]int, len(valves))
	}
	infinity := len(valves) * len(valves)
	for i := 0; i < len(valves); i++ {
		for j := 0; j < len(valves); j++ {
			if valves[j].label != valves[i].label {
				distances[i][j] = infinity
				for _, next := range valves[j].connections {
					if next == valves[i].label {
						distances[i][j] = 1
						break
					}
				}
			}
		}
	}
	for k := 0; k < len(distances); k++ {
		for i := 0; i < len(distances); i++ {
			for j := 0; j < len(distances); j++ {
				d := distances[i][k] + distances[k][j]
				if d < distances[i][j] {
					distances[i][j] = d
				}
			}
		}
	}
	return distances
}

func DFS(valves []Valve, time int) map[BitSet]int {
  distances := findShortestDistances(valves)

  rates := make([]int, len(valves))
	start := 0
	for i := range valves {
		rates[i] = valves[i].rate
		if valves[i].label == "AA" {
			start = i
		}
	}
  states := make(map[BitSet]int)
	stack := Stack{
		State{
			start,
			0,
			0,
			time,
		},
	}
  //for _, valve := range valves {
    for len(stack) != 0 {
      current := stack.Pop()
      if states[current.open] < current.pressure {
        states[current.open] = current.pressure
      }
      if current.time > 0 {
        rate := rates[current.valve]
        if rate != 0 && current.open.IsEmpty(current.valve) {
          time := current.time - 1
          stack.Push(State{
            current.valve,
            current.open.WithSet(current.valve),
            current.pressure + rate * time,
            time,
          })
        } else {
          for next, d := range distances[current.valve] {
            if next != current.valve && rates[next] != 0 && current.open.IsEmpty(next) && current.time > d {
              stack.Push(State{
                next,
                current.open,
                current.pressure,
                current.time - d,
              })
            }
          }
        }
      }
    }
  

  return states
}

func part1() int {
	dat, _ := os.ReadFile("./day16-input")
	s := strings.Split(string(dat), "\n")

	var readings []Valve
  for _, ln := range s[:len(s)-1] {
		readings = append(readings, parseLine(ln))
	}

  paths := DFS(readings, 30)

	max := 0
	for _, pressure := range paths {
		if max < pressure {
			max = pressure
		}
	}
	return max
}
