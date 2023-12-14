// Part 2 for day 14 of 2023

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

function rotate(matrix: string[][]): string[][] {
  return matrix[0].map((_, index) => matrix.map((row) => row[index]).reverse());
}

function cycle(groups: string[][]): string[][] {
  for (let i = 0; i < 4; i++) {
    groups = rollRocks(groups);
    groups = rotate(groups);
  }
  return groups;
}

export function part2(input: string): number {
  let groups = rollRocks(
    input
      .split("\n")
      .filter((group) => group != "")
      .map((group) => group.split("")),
  );

  let seen: Set<string> = new Set();
  let values = [groups.map((group) => group.join("")).join("")];
  seen.add(groups.map((group) => group.join("")).join(""));

  let firstHit = -1;
  let i = 0;
  while (true) {
    i++;
    groups = cycle(groups);
    if (seen.has(groups.map((group) => group.join("")).join(""))) {
      firstHit = i;
      break;
    }
    seen.add(groups.map((group) => group.join("")).join(""));
    values.push(groups.map((group) => group.join("")).join(""));
  }

  let first = values.indexOf(groups.map((group) => group.join("")).join(""));
  let rem = (1000000000 - first) % (firstHit - first);

  for (let i = 0; i < rem; i++) {
    groups = cycle(groups);
  }

  let total = 0;

  for (let i = 0; i < groups.length; i++) {
    for (let j = 0; j < groups[i].length; j++) {
      total += groups[i][j] == "O" ? groups.length - i : 0;
    }
  }

  return total;
}
