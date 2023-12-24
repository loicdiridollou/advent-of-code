// Part 2 for day 23 of 2023

import { dfs } from "./part1";

export function part2(input: string): number {
  let grid = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) => c.split(""));
  let start: [number, number] = [0, grid[0].indexOf(".")];
  let end: [number, number] = [
    grid.length - 1,
    grid[grid.length - 1].indexOf("."),
  ];
  let points = [start, end];

  for (let r = 0; r < grid.length; r++) {
    for (let c = 0; c < grid[0].length; c++) {
      if (grid[r][c] == "#") {
        continue;
      }
      let neighbors = 0;
      for (let [dr, dc] of [
        [0, -1],
        [0, 1],
        [-1, 0],
        [1, 0],
      ]) {
        let [nr, nc] = [r + dr, c + dc];
        if (
          0 <= nr &&
          nr < grid.length &&
          0 <= nc &&
          nc < grid[0].length &&
          grid[nr][nc] != "#"
        ) {
          neighbors++;
        }
      }
      if (neighbors >= 3) {
        points.push([r, c]);
      }
    }
  }
  let pointsSet = new Set<string>();
  for (let point of points) {
    pointsSet.add(point.toString());
  }

  let graph: { [index: string]: { [index: string]: number } } = {};
  for (let point of points) {
    graph[point.toString()] = {};
  }

  let dirs: { [index: string]: [number, number][] } = {};
  for (let chr of ["^", "v", "<", ">", "."]) {
    dirs[chr] = [
      [-1, 0],
      [1, 0],
      [0, -1],
      [0, 1],
    ];
  }

  for (let [sr, sc] of points) {
    let stack = [[0, sr, sc]];
    let seen = new Set<string>();

    while (stack.length > 0) {
      let [n, r, c] = stack[0];
      stack = stack.slice(1);

      if (n != 0 && pointsSet.has([r, c].toString())) {
        graph[[sr, sc].toString()][[r, c].toString()] = n;
        continue;
      }

      for (let [dr, dc] of dirs[grid[r][c]]) {
        let nr = r + dr;
        let nc = c + dc;
        if (
          0 <= nr &&
          nr < grid.length &&
          0 <= nc &&
          nc < grid[0].length &&
          grid[nr][nc] != "#" &&
          !seen.has([nr, nc].toString())
        ) {
          stack.push([n + 1, nr, nc]);
          seen.add([nr, nc].toString());
        }
      }
    }
  }

  let seen = new Set<string>();
  return dfs(start, graph, seen, end);
}
