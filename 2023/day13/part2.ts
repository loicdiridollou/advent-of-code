// Part 2 for day 13 of 2023

import { transpose } from "./part1";

function mirrorRow(array: string[], row: number): boolean {
  let top = array.slice(0, row).reverse();
  let bot = array.slice(row, array.length);

  top = top.slice(0, bot.length);
  bot = bot.slice(0, top.length);

  let diffs = 0;
  for (let i = 0; i < top.length; i++) {
    for (let j = 0; j < top[0].length; j++) {
      if (top[i][j] != bot[i][j]) {
        diffs++;
      }
    }
  }

  return diffs == 1;
}

export function part2(input: string): number {
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
