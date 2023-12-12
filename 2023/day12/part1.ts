// Part 1 for day 12 of 2023

export function arrayEquals(a: number[], b: number[]) {
  return (
    Array.isArray(a) &&
    Array.isArray(b) &&
    a.length === b.length &&
    a.every((val, index) => val === b[index])
  );
}

function countConfig(cfg: string, nums: number[]): number {
  if (cfg == "") {
    return arrayEquals(nums, []) ? 1 : 0;
  }
  if (arrayEquals(nums, [])) {
    return cfg.includes("#") ? 0 : 1;
  }
  let total = 0;

  if (".?".includes(cfg[0])) {
    // we can treat the ? as a . and see
    total += countConfig(cfg.slice(1), nums);
  }
  if ("#?".includes(cfg[0])) {
    // need to do more checks
    if (
      nums[0] <= cfg.length &&
      !cfg.slice(0, nums[0]).includes(".") &&
      (nums[0] == cfg.length || cfg[nums[0]] != "#")
    ) {
      total += countConfig(cfg.slice(nums[0] + 1), nums.slice(1));
    }
  }

  return total;
}

export function part1(input: string): number {
  let dataGroups = input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(" "));

  let total = [];
  for (let data of dataGroups) {
    let code = data[1].split(",").map((group) => parseInt(group));
    total.push(countConfig(data[0], code));
  }

  return total.reduce((x, y) => x + y, 0);
}
