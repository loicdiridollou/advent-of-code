// Part 2 for day 22 of 2023

import { bricksOverlap } from "./part1";

function isSubset(set1: Set<number>, set2: Set<number>): boolean {
  for (let key of set1) {
    if (!set2.has(key)) {
      return false;
    }
  }
  return true;
}
export function part2(input: string): number {
  let bricks = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) =>
      c
        .replace("~", ",")
        .split(",")
        .filter((b) => b != "")
        .map((c) => parseInt(c)),
    );

  bricks = bricks.sort((a, b) => (a[2] > b[2] ? 1 : -1));

  for (let i = 0; i < bricks.length; i++) {
    let maxZ = 1;
    for (let j = 0; j < i; j++) {
      if (bricksOverlap(bricks[i], bricks[j])) {
        maxZ = Math.max(bricks[j][5] + 1, maxZ);
      }
    }
    bricks[i][5] -= bricks[i][2] - maxZ;
    bricks[i][2] = maxZ;
  }

  bricks = bricks.sort((a, b) => (a[2] > b[2] ? 1 : -1));
  let support1: { [index: number]: Set<number> } = {};
  let support2: { [index: number]: Set<number> } = {};

  for (let i = 0; i < bricks.length; i++) {
    support1[i] = new Set<number>();
    support2[i] = new Set<number>();
  }

  for (let i = 0; i < bricks.length; i++) {
    for (let j = 0; j < i; j++) {
      if (
        bricksOverlap(bricks[i], bricks[j]) &&
        bricks[i][2] == bricks[j][5] + 1
      ) {
        support1[j].add(i);
        support2[i].add(j);
      }
    }
  }

  let total = 0;

  for (let i = 0; i < bricks.length; i++) {
    let q: number[] = [];
    for (let j of support1[i]) {
      if (support2[j].size === 1) {
        q.push(j);
      }
    }

    let falling = new Set<number>(q);
    falling.add(i);

    while (q.length > 0) {
      const j = q.shift()!;
      for (let k of support1[j]) {
        if (!falling.has(k)) {
          if (isSubset(support2[k], falling)) {
            q.push(k);
            falling.add(k);
          }
        }
      }
    }

    total += falling.size - 1;
  }

  return total;
}
