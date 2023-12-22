// Part 1 for day 22 of 2023

export function bricksOverlap(a: number[], b: number[]): boolean {
  return (
    Math.max(a[0], b[0]) <= Math.min(a[3], b[3]) &&
    Math.max(a[1], b[1]) <= Math.min(a[4], b[4])
  );
}

export function part1(input: string): number {
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
    let bool = true;
    for (let j of support1[i]) {
      if (support2[j].size < 2) {
        bool = false;
        break;
      }
    }
    total += bool ? 1 : 0;
  }

  return total;
}
