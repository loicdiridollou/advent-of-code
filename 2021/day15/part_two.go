package main

import (
	"container/heap"
	"math"
	"strings"
)

var (
	D = [4]XY{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

type XY struct {
	x, y int
}

type Edge struct {
	u, v int
	w    int
}

type Graph struct {
	g [][]int
	s int
}

func (g *Graph) xyToN(xy XY) int {
	return (xy.y * g.s) + xy.x
}

func (g *Graph) nToXY(n int) *XY {
	return &XY{n % g.s, n / g.s}
}

func (g *Graph) adjacentNodes(from int) (edges []Edge) {
	// if you've reached the lower right corner, there's only the 0-weight path to exit
	if from == g.s*g.s-1 {
		edges = append(edges, Edge{u: from, v: g.s * g.s, w: 0})
	} else {
		xy := g.nToXY(from)

		for _, d := range D {
			nx, ny := xy.x+d.x, xy.y+d.y

			if 0 <= nx && nx < g.s && 0 <= ny && ny < g.s {
				to := g.xyToN(XY{nx, ny})
				edges = append(edges, Edge{u: from, v: to, w: g.g[ny][nx]})
			}
		}
	}

	return edges
}

type Item struct {
	Value    int
	Priority int
	Index    int
}

func parseInputRepeating(lines []string) *Graph {
	g := &Graph{}
	g.s = len(lines) * 5
	g.g = make([][]int, g.s)
	for i := 0; i < g.s; i++ {
		g.g[i] = make([]int, g.s)
	}

	for y, line := range lines {
		for x, cell := range line {
			g.g[y][x] = StrToInt(string(cell))
		}
	}

	blocksize := len(lines)
	// repeat horizontal
	for i := 1; i < 5; i++ {
		for y := 0; y < blocksize; y++ {
			for x := 0; x < blocksize; x++ {
				switch v := ((g.g[y][x] % 9) + i) % 9; v {
				case 0:
					g.g[y][x+i*blocksize] = 9
				default:
					g.g[y][x+i*blocksize] = v
				}
			}
		}
	}
	// and now repeat all downwards
	// repeat horizontal
	for i := 1; i < 5; i++ {
		for y := 0; y < blocksize; y++ {
			for x := 0; x < g.s; x++ {
				switch v := ((g.g[y][x] % 9) + i) % 9; v {
				case 0:
					g.g[y+i*blocksize][x] = 9
				default:
					g.g[y+i*blocksize][x] = v
				}
			}
		}
	}

	return g
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (g *Graph) dijkstra(source int) ([]int, []int) {
	dist := make([]int, g.s*g.s+1)
	prev := make([]int, g.s*g.s+1)
	inPQ := make([]*Item, g.s*g.s+1)

	pq := make(PriorityQueue, 0, g.s*g.s+1)

	// source node
	dist[0] = 0
	item := &Item{
		Value:    0,
		Priority: 0,
		Index:    0,
	}
	pq = append(pq, item)
	inPQ[0] = item

	// nodes in the graph
	for i := 1; i < len(dist); i++ {
		dist[i] = math.MaxInt64
	}

	// initialize priority queue (PQ)
	heap.Init(&pq)
	for pq.Len() != 0 {
		// pop the next best current from the PQ
		current := heap.Pop(&pq).(*Item)
		inPQ[current.Value] = nil

		// early exit
		if current.Value == g.s*g.s {
			return dist, prev
		}

		// and look at all the edges of the current node
		for _, edge := range g.adjacentNodes(current.Value) {
			// check if the edge would be an improvement
			if alt := current.Priority + edge.w; alt < dist[edge.v] {
				prev[edge.v] = current.Value
				dist[edge.v] = alt

				// if the neighbour node is in the PQ already, update it, otherwise add a new one
				if next := inPQ[edge.v]; next != nil {
					next.Value = edge.v
					next.Priority = alt
					heap.Fix(&pq, next.Index)
				} else {
					next := &Item{
						Value:    edge.v,
						Priority: alt,
					}
					heap.Push(&pq, next)
					inPQ[edge.v] = next
				}
			}
		}
	}

	return dist, prev
}

func part2(input string) int {
	lines := make([]string, 0)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}
	g := parseInputRepeating(lines)
	d, _ := g.dijkstra(0)
	return d[g.s*g.s]
}
