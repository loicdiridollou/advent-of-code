// Part 1 for day 14 of 2023

export function transpose(matrix: string[][]) {
  return matrix[0].map((_, i) => matrix.map((row) => row[i]));
}

export function rollRocks(groups: string[][]): string[][] {
  let newGroups = transpose(groups).map((group) => group.join(""));
  newGroups = newGroups.map((group) =>
    group
      .split("#")
      .map((group) =>
        group
          .split("")
          .sort(function (x, _) {
            return x == "O" ? -1 : 1;
          })
          .join(""),
      )
      .join("#"),
  );

  return transpose(newGroups.map((group) => group.split("")));
}

export function part1(input: string): number {
  let groups = rollRocks(
    input
      .split("\n")
      .filter((group) => group != "")
      .map((group) => group.split("")),
  );

  let total = 0;

  for (let i = 0; i < groups.length; i++) {
    for (let j = 0; j < groups[i].length; j++) {
      total += groups[i][j] == "O" ? groups.length - i : 0;
    }
  }

  return total;
}
