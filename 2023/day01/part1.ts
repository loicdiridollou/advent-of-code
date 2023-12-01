// Part 1 for day 01 of 2023
function is_numeric(str: string) {
  return /^\d+$/.test(str);
}

export function part1(words: string): number {
  return words
    .split("\n")
    .map((group) => group.split("").filter(is_numeric))
    .filter((group) => group.length > 0)
    .map((group) => group[0] + group[group.length - 1])
    .map((val) => parseInt(val))
    .reduce((sum, current) => sum + current, 0);
}
