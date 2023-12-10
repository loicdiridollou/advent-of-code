// Part 1 for day 10 of 2023

function getNeighbors(
  row: number,
  col: number,
  length: number,
  width: number,
): number[][] {
  let neighbors: number[][] = [];

  if (row > 0) {
    neighbors.push([row - 1, col]);
  }
  if (row < length - 1) {
    neighbors.push([row + 1, col]);
  }
  if (col > 0) {
    neighbors.push([row, col - 1]);
  }
  if (col < width - 1) {
    neighbors.push([row, col + 1]);
  }

  return neighbors;
}

export function part1(input: string): number {
  let maze = input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(""));

  let start: number[] = [];
  for (let row = 0; row < maze.length; row++) {
    for (let col = 0; col < maze[row].length; col++) {
      if (maze[row][col] == "S") {
        start = [row, col];
      }
    }
  }

  let visited = new Set();
  visited.add(`${start[0]}_${start[1]}`);
  let queue = [`${start[0]}_${start[1]}`];

  while (queue.length != 0) {
    let new_pos = queue[0];
    queue = queue.slice(1);
    let r = parseInt(new_pos.split("_")[0]);
    let c = parseInt(new_pos.split("_")[1]);
    let chr = maze[r][c];

    if (
      r > 0 &&
      "S|JL".includes(chr) &&
      "|7F".includes(maze[r - 1][c]) &&
      !visited.has(`${r - 1}_${c}`)
    ) {
      visited.add(`${r - 1}_${c}`);
      queue.push(`${r - 1}_${c}`);
    }
    if (
      r < maze.length - 1 &&
      "S|7F".includes(chr) &&
      "|JL".includes(maze[r + 1][c]) &&
      !visited.has(`${r + 1}_${c}`)
    ) {
      visited.add(`${r + 1}_${c}`);
      queue.push(`${r + 1}_${c}`);
    }
    if (
      c > 0 &&
      "S-J7".includes(chr) &&
      "-LF".includes(maze[r][c - 1]) &&
      !visited.has(`${r}_${c - 1}`)
    ) {
      visited.add(`${r}_${c - 1}`);
      queue.push(`${r}_${c - 1}`);
    }
    if (
      c < maze[0].length - 1 &&
      "S-LF".includes(chr) &&
      "-J7".includes(maze[r][c + 1]) &&
      !visited.has(`${r}_${c + 1}`)
    ) {
      visited.add(`${r}_${c + 1}`);
      queue.push(`${r}_${c + 1}`);
    }
  }

  return visited.size / 2;
}
