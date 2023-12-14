// Part 1 for day 12 of 2023

export function arrayEquals(a: string[], b: string[]) {
  return (
    Array.isArray(a) &&
    Array.isArray(b) &&
    a.length === b.length &&
    a.every((val, index) => val === b[index])
  );
}

export function transpose(matrix: string[][]) {
  return matrix[0].map((_, i) => matrix.map((row) => row[i]));
}

function mirrorRow(array: string[], row: number): boolean {
  let top = array.slice(0, row).reverse();
  let bot = array.slice(row, array.length);

  top = top.slice(0, bot.length);
  bot = bot.slice(0, top.length);
  return arrayEquals(top, bot);
}

export function part1(input: string): number {
  let groups = input
    .split("\n\n")
    .filter((group) => group != "")
    .map((group) => group.split("\n").filter((group) => group != ""));

  let rows = [];
  let cols = [];

  for (let first of groups) {
    for (let i = 1; i < first.length; i++) {
      if (mirrorRow(first, i)) {
        rows.push(100 * i);
        break;
      }
    }

    let transposeMat = transpose(first.map((group) => group.split(""))).map(
      (group) => group.join(""),
    );
    for (let i = 1; i < transposeMat.length; i++) {
      if (mirrorRow(transposeMat, i)) {
        cols.push(i);
        break;
      }
    }
  }

  return rows.reduce((x, y) => x + y, 0) + cols.reduce((x, y) => x + y, 0);
}
