// Part 1 for day 24 of 2023

export function part1(input: string): number {
  let hailstones: number[][] = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) =>
      c
        .replace("@", ",")
        .split(",")
        .map((c) => parseInt(c)),
    );

  let total = 0;
  let min = 200000000000000;
  let max = 400000000000000;

  for (let i = 0; i < hailstones.length; i++) {
    for (let j = 0; j < i; j++) {
      let hs1 = hailstones[i];
      let hs2 = hailstones[j];

      let [a1, b1, c1] = [hs1[4], -hs1[3], hs1[4] * hs1[0] - hs1[3] * hs1[1]];
      let [a2, b2, c2] = [hs2[4], -hs2[3], hs2[4] * hs2[0] - hs2[3] * hs2[1]];

      if (a1 * b2 == a2 * b1) {
        continue;
      }

      let x = (c2 * b1 - c1 * b2) / (a2 * b1 - a1 * b2);
      let y = (c1 * a2 - c2 * a1) / (a2 * b1 - a1 * b2);

      if (x >= min && x <= max && y >= min && y <= max) {
        if (
          (x - hs1[0]) * hs1[3] >= 0 &&
          (y - hs1[1]) * hs1[4] >= 0 &&
          (x - hs2[0]) * hs2[3] >= 0 &&
          (y - hs2[1]) * hs2[4] >= 0
        ) {
          total++;
        }
      }
    }
  }
  return total;
}
