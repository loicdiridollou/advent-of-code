// Part 1 for day 21 of 2023

export function part1(input: string): number {
  let grid = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) => c.split("").filter((b) => b != ""));

  let start: [number, number] = [-1, -1];
  for (let r = 0; r < grid.length; r++) {
    for (let c = 0; c < grid[0].length; c++) {
      if (grid[r][c] == "S") {
        start = [r, c];
        break;
      }
    }
  }
  grid[start[0]][start[1]] = ".";
  let visited = new Set<string>();
  let ans = new Set<string>();
  ans.add(start.toString());
  let queue: [number, number, number][] = [];

  queue.push([...start, 64]);

  while (queue.length > 0) {
    let [r, c, n] = queue[0];
    queue = queue.slice(1);
    if (n % 2 == 0) {
      ans.add([r, c].toString());
    }
    if (n == 0) {
      continue;
    }

    for (let [dr, dc] of [
      [0, 1],
      [1, 0],
      [0, -1],
      [-1, 0],
    ]) {
      let [nr, nc] = [r + dr, c + dc];
      if (
        nr < 0 ||
        nr >= grid.length ||
        nc < 0 ||
        nc >= grid[0].length ||
        grid[nr][nc] == "#" ||
        visited.has([nr, nc].toString())
      ) {
        continue;
      }

      visited.add([nr, nc].toString());
      queue.push([nr, nc, n - 1]);
    }
  }
  return ans.size;
}
