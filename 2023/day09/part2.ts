// Part 2 for day 06 of 2023

import { getLast } from "./part1";

function parseInput(input: string): number[][] {
  return input
    .split("\n")
    .filter((group) => group != "")
    .map((group) =>
      group
        .split(" ")
        .reverse()
        .map((x) => parseInt(x)),
    );
}

export function part2(input: string): number {
  return parseInput(input)
    .map(getLast)
    .reduce((x, y) => x + y);
}
