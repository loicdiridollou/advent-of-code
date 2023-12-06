// Part 2 for day 06 of 2023

import { parseLine, solveTime } from "./part1";

export function part2(input: string): number {
  let data = input.split("\n").filter((group) => group != "");
  let time = parseInt(parseLine(data[0]).join(""));
  let distance = parseInt(parseLine(data[1]).join(""));

  return solveTime(time, distance);
}
