// Part 2 for day 21 of 2023

function fill(sr: number, sc: number, grid: string[][], steps: number): number {
  let visited = new Set<string>();
  let ans = new Set<string>();
  visited.add([sr, sc].toString());
  let queue: [number, number, number][] = [];

  queue.push([sr, sc, steps]);

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

export function part2(input: string): number {
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
  let [sr, sc] = start;
  let steps = 26501365;
  let size = grid.length;
  let gridWidth = Math.floor(steps / grid.length) - 1;
  let odd = (Math.floor(gridWidth / 2) * 2 + 1) ** 2;
  let even = (Math.floor((gridWidth + 1) / 2) * 2) ** 2;

  let odd_points = fill(sr, sc, grid, size * 2 + 1);
  let even_points = fill(sr, sc, grid, size * 2);
  let corner_t = fill(size - 1, sc, grid, size - 1);
  let corner_r = fill(sr, 0, grid, size - 1);
  let corner_b = fill(0, sc, grid, size - 1);
  let corner_l = fill(sr, size - 1, grid, size - 1);
  let small_tr = fill(size - 1, 0, grid, Math.floor(size / 2) - 1);
  let small_tl = fill(size - 1, size - 1, grid, Math.floor(size / 2) - 1);
  let small_br = fill(0, 0, grid, Math.floor(size / 2) - 1);
  let small_bl = fill(0, size - 1, grid, Math.floor(size / 2) - 1);
  let large_tr = fill(size - 1, 0, grid, Math.floor((size * 3) / 2) - 1);
  let large_tl = fill(size - 1, size - 1, grid, Math.floor((size * 3) / 2) - 1);
  let large_br = fill(0, 0, grid, Math.floor((size * 3) / 2) - 1);
  let large_bl = fill(0, size - 1, grid, Math.floor((size * 3) / 2) - 1);

  return (
    odd * odd_points +
    even * even_points +
    corner_t +
    corner_r +
    corner_b +
    corner_l +
    (gridWidth + 1) * (small_tr + small_tl + small_br + small_bl) +
    gridWidth * (large_tr + large_tl + large_br + large_bl)
  );
}
