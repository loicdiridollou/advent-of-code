// Part 1 for day 15 of 2023

export function convertStep(step: string): number {
  let total = 0;

  for (let chr of step) {
    total += chr.charCodeAt(0);
    total *= 17;
    total %= 256;
  }

  return total;
}

export function part1(input: string): number {
  return input
    .split(",")
    .filter((c) => c != "")
    .map((val) => val.replace(/(\r\n|\n|\r)/gm, ""))
    .reduce(
      (curr, step) =>
        curr +
        step.split("").reduce((x, y) => ((x + y.charCodeAt(0)) * 17) % 256, 0),
      0,
    );
}
