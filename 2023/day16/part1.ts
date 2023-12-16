// Part 1 for day 16 of 2023

export function fillGrid(grid: string[][], queue: any[]): number {
  let visited = new Set<string>();
  let visitedAngle = new Set<string>();

  while (queue.length > 0) {
    let [point, dir] = queue[0];
    queue = queue.slice(1);
    if (visitedAngle.has(`${point[0]}_${point[1]}_${dir}`)) {
      continue;
    }
    visited.add(`${point[0]}_${point[1]}`);
    visitedAngle.add(`${point[0]}_${point[1]}_${dir}`);

    if (dir == "R") {
      if (
        ".-".includes(grid[point[0]][point[1]]) &&
        point[1] + 1 < grid[0].length
      ) {
        queue.push([[point[0], point[1] + 1], "R"]);
      } else if (
        grid[point[0]][point[1]] == "\\" &&
        point[0] + 1 < grid.length
      ) {
        queue.push([[point[0] + 1, point[1]], "D"]);
      } else if (grid[point[0]][point[1]] == "/" && point[0] - 1 >= 0) {
        queue.push([[point[0] - 1, point[1]], "U"]);
      } else if (grid[point[0]][point[1]] == "|") {
        if (point[0] - 1 >= 0) {
          queue.push([[point[0] - 1, point[1]], "U"]);
        }
        if (point[0] + 1 < grid.length) {
          queue.push([[point[0] + 1, point[1]], "D"]);
        }
      }
    } else if (dir == "L") {
      if (".-".includes(grid[point[0]][point[1]]) && point[1] - 1 >= 0) {
        queue.push([[point[0], point[1] - 1], "L"]);
      } else if (grid[point[0]][point[1]] == "\\" && point[0] - 1 >= 0) {
        queue.push([[point[0] - 1, point[1]], "U"]);
      } else if (
        grid[point[0]][point[1]] == "/" &&
        point[0] + 1 < grid.length
      ) {
        queue.push([[point[0] + 1, point[1]], "D"]);
      } else if (grid[point[0]][point[1]] == "|") {
        if (point[0] - 1 >= 0) {
          queue.push([[point[0] - 1, point[1]], "U"]);
        }
        if (point[0] + 1 < grid.length) {
          queue.push([[point[0] + 1, point[1]], "D"]);
        }
      }
    } else if (dir == "U") {
      if (".|".includes(grid[point[0]][point[1]]) && point[0] - 1 >= 0) {
        queue.push([[point[0] - 1, point[1]], "U"]);
      } else if (grid[point[0]][point[1]] == "\\" && point[1] - 1 >= 0) {
        queue.push([[point[0], point[1] - 1], "L"]);
      } else if (
        grid[point[0]][point[1]] == "/" &&
        point[1] + 1 < grid[0].length
      ) {
        queue.push([[point[0], point[1] + 1], "R"]);
      } else if (grid[point[0]][point[1]] == "-") {
        if (point[1] - 1 >= 0) {
          queue.push([[point[0], point[1] - 1], "L"]);
        }
        if (point[1] + 1 < grid[0].length) {
          queue.push([[point[0], point[1] + 1], "R"]);
        }
      }
    } else if (dir == "D") {
      if (
        ".|".includes(grid[point[0]][point[1]]) &&
        point[0] + 1 < grid.length
      ) {
        queue.push([[point[0] + 1, point[1]], "D"]);
      } else if (
        grid[point[0]][point[1]] == "\\" &&
        point[1] + 1 < grid[0].length
      ) {
        queue.push([[point[0], point[1] + 1], "R"]);
      } else if (grid[point[0]][point[1]] == "/" && point[1] - 1 >= 0) {
        queue.push([[point[0], point[1] - 1], "L"]);
      } else if (grid[point[0]][point[1]] == "-") {
        if (point[1] - 1 >= 0) {
          queue.push([[point[0], point[1] - 1], "L"]);
        }
        if (point[1] + 1 < grid[0].length) {
          queue.push([[point[0], point[1] + 1], "R"]);
        }
      }
    }
  }
  return visited.size;
}

export function part1(input: string): number {
  let grid = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) => c.split(""));

  return fillGrid(grid, [[[0, 0], "R"]]);
}
