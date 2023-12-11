// Part 2 for day 11 of 2023

import {
  computeDistances,
  getEmptyRowsCols,
  getGalaxies,
  getPairs,
} from "./part1";

export function part2(input: string): number {
  let maze = input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(""));

  let [empty_rows, empty_cols] = getEmptyRowsCols(maze);
  let pairs = getPairs(getGalaxies(maze));

  return computeDistances(pairs, empty_cols, empty_rows, 1000000);
}
