// Part 2 for day 08 of 2023

import { createNetwork } from "./part1";

function gcd(a: number, b: number): number {
  return b == 0 ? a : gcd(b, a % b);
}

export function part2(input: string): number {
  let instructions = input.split("\n\n")[0];
  let res = createNetwork(input.split("\n\n")[1].split("\n"));
  let starts = Object.keys(res).filter((group) => group.endsWith("A"));

  let cycles = [];

  for (let start of starts) {
    let cycle = [];
    let num_iter = 0;
    let first_hit = null;

    while (true) {
      while (num_iter == 0 || !start.endsWith("Z")) {
        let op = instructions[num_iter % instructions.length];
        start = res[start][op == "L" ? 0 : 1];
        num_iter++;
      }
      cycle.push(num_iter);
      if (first_hit == null) {
        first_hit = start;
        num_iter = 0;
      } else if (start == first_hit) {
        break;
      }
    }
    cycles.push(cycle);
  }

  let numbers = cycles.map((group) => group[0]);
  let _gcd: number = numbers.pop() as number;
  for (let num of numbers) {
    _gcd = (_gcd * num) / gcd(_gcd, num);
  }

  return _gcd;
}
