// Part 2 for day 23 of 2023

let graph: { [index: string]: { [index: string]: number } } = {};
let start: [number, number] = [-1, -1];
let end: [number, number] = [-1, -1];
let seen: Set<string>;

function dfs(pt: [number, number]): number {
  if (pt.toString() == end.toString()) {
    return 0;
  }
  let m = -(10 ** 20);

  seen.add(pt.toString());
  for (let nx in graph[pt.toString()]) {
    if (!seen.has(nx)) {
      let [r, c] = nx.split(",").map((c) => parseInt(c));
      m = Math.max(m, dfs([r, c]) + graph[pt.toString()][nx.toString()]);
    }
  }
  seen.delete(pt.toString());

  return m;
}

export function part2(input: string): number {
  let grid = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) => c.split(""));
  start = [0, grid[0].indexOf(".")];
  end = [grid.length - 1, grid[grid.length - 1].indexOf(".")];
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

  graph = {};
  for (let point of points) {
    graph[point.toString()] = {};
  }

  let dirs: { [index: string]: [number, number][] } = {
    "^": [
      [-1, 0],
      [1, 0],
      [0, -1],
      [0, 1],
    ],
    v: [
      [-1, 0],
      [1, 0],
      [0, -1],
      [0, 1],
    ],
    "<": [
      [-1, 0],
      [1, 0],
      [0, -1],
      [0, 1],
    ],
    ">": [
      [-1, 0],
      [1, 0],
      [0, -1],
      [0, 1],
    ],
    ".": [
      [-1, 0],
      [1, 0],
      [0, -1],
      [0, 1],
    ],
  };

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

  seen = new Set<string>();
  return dfs(start);
}
