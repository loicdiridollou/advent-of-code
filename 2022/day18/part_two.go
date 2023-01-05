package main

import (
	_ "embed"
  "fmt"
	"os"
)

func hashValPoint(p Point) string {
  return fmt.Sprint(p.x) + "_" + fmt.Sprint(p.y) + "_" + fmt.Sprint(p.z)
}
 
func DFS2(points []Point, points_map map[string]Point) int {
  visible_sides := 0
  visited := make(map[string]bool, 0)
  queue_map := make(map[string]bool, 0)
  neighbors := [][3]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

  queue := make([]Point, 0)
  queue = append(queue, Point{0, 0, 0})
  var check Point

  for len(queue) > 0 {
    current := queue[len(queue)-1]
    queue = queue[:len(queue)-1]
    queue_map[hashValPoint(current)] = false
    visited[hashValPoint(current)] = true 
    for _, n := range neighbors {
      check = Point{current.x + n[0], current.y + n[1], current.z + n[2]}
      in_visited := visited[hashValPoint(check)]
      in_queue := queue_map[hashValPoint(check)]
      if -1 <= check.x && check.x <= 22 && -1 <= check.y && check.y <= 22 && -1 <= check.z && check.z <= 22 && !in_visited && !in_queue { 
        if _, ok := points_map[hashValPoint(check)]; ok {
          visible_sides++
        } else {
          queue = append(queue, check)
          queue_map[hashValPoint(check)] = true
        }
      }
    }
  }
  return visible_sides
}

func part2() int {
	dat, _ := os.ReadFile("./day18-input")
  points := parseLine(string(dat))

  points_map := make(map[string]Point, 0)
  for _, point := range points {
    points_map[hashVal(point.x, point.y, point.z)] = point
  }

  faces := DFS2(points, points_map)

  return faces
}
