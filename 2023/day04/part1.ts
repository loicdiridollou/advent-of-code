// Part 1 for day 02 of 2023

export function part1(input: string): number {
  return input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(":"))
    .map((group): number => {
      let cards = group[1];
      let win_numbers: number[] = cards
        .split("|")[0]
        .split(" ")
        .filter((group) => group != "")
        .map((group): number => parseInt(group));

      let hand_numbers: number[] = cards
        .split("|")[1]
        .split(" ")
        .filter((group) => group != "")
        .map((group): number => parseInt(group));

      let num_win = hand_numbers.filter((num) =>
        win_numbers.includes(num),
      ).length;
      return num_win < 2 ? num_win : Math.pow(2, num_win - 1);
    })
    .reduce((sum, current) => sum + current, 0);
}
