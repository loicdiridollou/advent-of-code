// Part 1 for day 02 of 2023

export function solveTime(time: number, distance: number): number {
  let x1 = (time - Math.sqrt(time * time - 4 * distance)) / 2;
  let x2 = (time + Math.sqrt(time * time - 4 * distance)) / 2;
  x1 = Math.ceil(x1);
  x2 = Math.floor(x2);
  if (x1 * (time - x1) == distance) {
    x1++;
  }
  if (x2 * (time - x2) == distance) {
    x2--;
  }

  return Math.max(0, x2 - x1 + 1);
}

export function parseLine(data: string): string[] {
  return data
    .split(":")[1]
    .split(" ")
    .filter((group) => group != "");
}

export function part1(input: string): number {
  let data = input.split("\n").filter((group) => group != "");
  let times = parseLine(data[0]).map((group) => parseInt(group));
  let distances = parseLine(data[1]).map((group) => parseInt(group));

  let prod = 1;
  for (let i in times) {
    prod *= solveTime(times[i], distances[i]);
  }
  return prod;
}
