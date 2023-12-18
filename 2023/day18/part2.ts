// Part 2 for day 18 of 2023

import { dirs } from "./part1";

const dirCodes: { [index: number]: string } = {
  0: "R",
  1: "D",
  2: "L",
  3: "U",
};

export function part2(input: string): number {
  let instrs = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) => c.split(" "));

  let loop: [number, number][] = [[0, 0]];
  let loopLen = 0;
  for (let instr of instrs) {
    let val = instr[2];
    let len = parseInt(val.slice(2, 7), 16);
    let dir = dirCodes[parseInt(val.slice(7, 8))];
    loopLen += len;

    let nr = loop[loop.length - 1][0] + dirs[dir][0] * len;
    let nc = loop[loop.length - 1][1] + dirs[dir][1] * len;
    loop.push([nr, nc]);
  }

  let total = 0;

  for (let i = 0; i < loop.length; i++) {
    total +=
      loop[i][0] *
      (loop[(i + loop.length - 1) % loop.length][1] -
        loop[(i + 1) % loop.length][1]);
  }

  return Math.abs(total) / 2 - loopLen / 2 + 1 + loopLen;
}
