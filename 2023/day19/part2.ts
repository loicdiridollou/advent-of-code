// Part 2 for day 19 of 2023

function parseInstruction(input: string): [string, string[][], string] {
  let cnt = input.match(/(.*){(.*)}/);

  let split = cnt![2].split(",");
  let term = split[split.length - 1];
  let instr = [];

  for (let idd of split.slice(0, split.length - 1)) {
    let match = idd.match(/(.*)([<>])(.*):(.*)/);
    if (match) {
      instr.push(match.slice(1));
    }
  }

  return [cnt![1], instr, term];
}

function countPossible(
  ranges: { [index: string]: [number, number] },
  workflows: any,
  pos: string = "in",
) {
  let tot = 1;
  let T: [number, number];
  let F: [number, number];

  if (pos == "A") {
    for (let [lo, hi] of Object.values(ranges)) {
      tot = tot * (hi - lo + 1);
    }
    return tot;
  } else if (pos == "R") {
    return 0;
  } else {
    tot = 0;
    let [rules, fallback] = workflows[pos];
    for (let [key, cmp, n, target] of rules) {
      let [lo, hi] = ranges[key];
      if (cmp == "<") {
        T = [lo, Math.min(hi, parseInt(n) - 1)];
        F = [Math.max(lo, parseInt(n)), hi];
      } else {
        T = [Math.max(lo, parseInt(n) + 1), hi];
        F = [lo, Math.min(hi, parseInt(n))];
      }
      if (T[0] <= T[1]) {
        let newRanges = { ...ranges };
        newRanges[key] = T;
        tot += countPossible(newRanges, workflows, target);
      }
      if (F[0] <= F[1]) {
        ranges = { ...ranges };
        ranges[key] = F;
      } else {
        break;
      }
    }
    tot += countPossible(ranges, workflows, fallback);
  }
  return tot;
}

export function part2(input: string): number {
  let instrRaw = input
    .split("\n\n")
    .filter((c) => c != "")
    .map((c) => c.split("\n").filter((c) => c != ""))[0];

  let instrMap: { [index: string]: [string[][], string] } = {};
  for (let instr of instrRaw) {
    let [name, ins, fallback] = parseInstruction(instr);
    instrMap[name] = [ins, fallback];
  }

  let total = countPossible(
    { x: [1, 4000], m: [1, 4000], a: [1, 4000], s: [1, 4000] },
    instrMap,
  );

  return total;
}
