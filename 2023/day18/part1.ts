// Part 1 for day 18 of 2023

export const dirs: { [index: string]: [number, number] } = {
  U: [-1, 0],
  D: [1, 0],
  L: [0, -1],
  R: [0, 1],
};

export function part1(input: string): number {
  let instrs = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) => c.split(" "));

  let loop: [number, number][] = [[0, 0]];
  let loopLen = 0;
  for (let instr of instrs) {
    let dir = instr[0];
    let len = parseInt(instr[1]);
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
