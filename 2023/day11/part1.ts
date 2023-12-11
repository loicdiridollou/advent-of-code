// Part 1 for day 11 of 2023

export function getEmptyRowsCols(maze: string[][]): Set<number>[] {
  let empty_rows = new Set<number>();
  let empty_cols = new Set<number>();

  for (let r = 0; r < maze.length; r++) {
    if (maze[r].every((chr) => chr == ".")) {
      empty_rows.add(r);
    }
  }
  for (let c = 0; c < maze[0].length; c++) {
    let empty = true;
    for (let r = 0; r < maze.length; r++) {
      if (maze[r][c] != ".") {
        empty = false;
        break;
      }
    }
    if (empty) {
      empty_cols.add(c);
    }
  }
  return [empty_rows, empty_cols];
}

export function getGalaxies(maze: string[][]): number[][] {
  let galaxies: number[][] = [];
  for (let r = 0; r < maze.length; r++) {
    for (let c = 0; c < maze[0].length; c++) {
      if (maze[r][c] == "#") {
        galaxies.push([r, c]);
      }
    }
  }
  return galaxies;
}

export function getPairs(galaxies: number[][]): number[][][] {
  let pairs: number[][][] = [];
  for (let i = 0; i < galaxies.length - 1; i++) {
    for (let j = i + 1; j < galaxies.length; j++) {
      pairs.push([galaxies[i], galaxies[j]]);
    }
  }
  return pairs;
}

export function computeDistances(
  pairs: number[][][],
  empty_cols: Set<number>,
  empty_rows: Set<number>,
  expansion: number = 1,
): number {
  let total = 0;
  for (let pair of pairs) {
    let r1 = pair[0][0];
    let c1 = pair[0][1];
    let r2 = pair[1][0];
    let c2 = pair[1][1];

    for (let i = Math.min(r1, r2); i < Math.max(r1, r2); i++) {
      total += empty_rows.has(i) ? expansion : 1;
    }
    for (let i = Math.min(c1, c2); i < Math.max(c1, c2); i++) {
      total += empty_cols.has(i) ? expansion : 1;
    }
  }
  return total;
}

export function part1(input: string): number {
  let maze = input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(""));

  let [empty_rows, empty_cols] = getEmptyRowsCols(maze);
  let pairs = getPairs(getGalaxies(maze));

  return computeDistances(pairs, empty_cols, empty_rows, 2);
}
