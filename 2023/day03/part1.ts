// Part 1 for day 03 of 2023
function is_numeric(str: string) {
  return /^\d+$/.test(str);
}

function is_symbol(str: string) {
  return !is_numeric(str) && str != ".";
}

function getNeighbors(
  row: number,
  start: number,
  end: number,
  length: number,
  width: number,
): string[] {
  let positions: string[] = [];
  if (row > 0) {
    for (
      let i: number = Math.max(0, start - 1);
      i <= Math.min(end + 1, width - 1);
      i++
    ) {
      positions.push(row - 1 + "_" + i);
    }
  }
  if (start > 0) {
    positions.push(row.toString() + "_" + (start - 1).toString());
  }
  if (end < width - 1) {
    positions.push(row.toString() + "_" + (end + 1).toString());
  }
  if (row < length - 1) {
    for (
      let i: number = Math.max(0, start - 1);
      i <= Math.min(end + 1, width - 1);
      i++
    ) {
      positions.push(row + 1 + "_" + i);
    }
  }

  return positions;
}

function checkNeighbors(
  row: any,
  start: number,
  end: number,
  neighbors: string[],
  grid: string[][],
): number {
  for (var ngb of neighbors) {
    let split = ngb.split("_");
    let ngb_row = parseInt(split[0]);
    let ngb_col = parseInt(split[1]);

    if (is_symbol(grid[ngb_row][ngb_col])) {
      return parseInt(grid[row].slice(start, end + 1).join(""));
    }
  }
  return 0;
}

export function part1(input: string): number {
  let grid = input
    .split("\n")
    .filter((group) => group.length > 0)
    .map((element) => element.split(""));

  let start = 0;
  let end = 0;
  let neighbors: string[];
  let sum: number = 0;

  for (var row in grid) {
    let number = false;
    for (var col in grid[row]) {
      if (is_numeric(grid[row][col]) && !number) {
        number = true;
        start = parseInt(col);
        end = parseInt(col);
      } else if (is_numeric(grid[row][col]) && number) {
        end = parseInt(col);
      } else if (!is_numeric(grid[row][col]) && number) {
        number = false;
        // check around for symbols
        neighbors = getNeighbors(
          parseInt(row),
          start,
          end,
          grid.length,
          grid[0].length,
        );

        sum += checkNeighbors(row, start, end, neighbors, grid);
      }

      // deal with number at the end of the line separately
      if (parseInt(col) == grid[row].length - 1 && number) {
        number = false;
        // check around for symbols
        neighbors = getNeighbors(
          parseInt(row),
          start,
          end,
          grid.length,
          grid[0].length,
        );
        sum += checkNeighbors(row, start, end, neighbors, grid);
      }
    }
  }
  return sum;
}
