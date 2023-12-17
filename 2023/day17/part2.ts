// Part 2 for day 17 of 2023

import { heappop, heappush } from "./part1";

export function part2(input: string): number {
  let grid = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) => c.split("").map((c) => parseInt(c)));

  let hq = [[0, 0, 0, 0, 0, 0]];
  let seen = new Set<string>();
  let total = -1;

  while (hq.length > 0) {
    let [hl, r, c, dr, dc, n] = heappop(hq);

    if (r == grid.length - 1 && c == grid[0].length - 1 && n >= 4) {
      total = hl;
      break;
    }

    if (seen.has(`${r}_${c}_${dr}_${dc}_${n}`)) {
      continue;
    }
    seen.add(`${r}_${c}_${dr}_${dc}_${n}`);

    if (n < 10 && !(dr == 0 && dc == 0)) {
      let [nr, nc] = [r + dr, c + dc];
      if (nr >= 0 && nr < grid.length && nc >= 0 && nc < grid[0].length) {
        heappush(hq, [hl + grid[nr][nc], nr, nc, dr, dc, n + 1]);
      }
    }

    if (n >= 4 || (dr == 0 && dc == 0)) {
      for (let nd of [
        [0, 1],
        [1, 0],
        [0, -1],
        [-1, 0],
      ]) {
        let [ndr, ndc] = nd;
        if ((ndr == dr && ndc == dc) || (ndr == -dr && ndc == -dc)) {
          continue;
        }
        let [nr, nc] = [r + ndr, c + ndc];
        if (nr >= 0 && nr < grid.length && nc >= 0 && nc < grid[0].length) {
          heappush(hq, [hl + grid[nr][nc], nr, nc, ndr, ndc, 1]);
        }
      }
    }
  }

  return total;
}
