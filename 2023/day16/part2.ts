// Part 2 for day 16 of 2023

import { fillGrid } from "./part1";

export function part2(input: string): number {
  let grid = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) => c.split(""));

  let total = 0;

  for (let i = 0; i < grid.length; i++) {
    total = Math.max(total, fillGrid(grid, [[[i, 0], "R"]]));
    total = Math.max(total, fillGrid(grid, [[[i, grid[0].length - 1], "L"]]));
  }

  for (let i = 0; i < grid[0].length; i++) {
    total = Math.max(total, fillGrid(grid, [[[0, i], "D"]]));
    total = Math.max(total, fillGrid(grid, [[[grid.length - 1, i], "U"]]));
  }

  return total;
}
