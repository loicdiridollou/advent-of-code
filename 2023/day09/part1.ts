// Part 1 for day 09 of 2023

export function getLast(sr: number[]): number {
  let sr_c = sr;
  let last = [];
  while (!sr_c.every((x) => x == 0)) {
    let sr_tmp = [];
    for (let i = 0; i < sr_c.length - 1; i++) {
      sr_tmp.push(sr_c[i + 1] - sr_c[i]);
    }
    sr_c = sr_tmp;
    last.push(sr_tmp[sr_tmp.length - 1]);
  }
  return sr[sr.length - 1] + last.reduce((x, y) => x + y);
}

function parseInput(input: string): number[][] {
  return input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(" ").map((x) => parseInt(x)));
}

export function part1(input: string): number {
  return parseInput(input)
    .map(getLast)
    .reduce((x, y) => x + y);
}
