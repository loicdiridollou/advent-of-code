// Part 1 for day 04 of 2023

export function parseCards(cards: string): number[] {
  return cards
    .split(" ")
    .filter((group) => group != "")
    .map((group): number => parseInt(group));
}

export function part1(input: string): number {
  return input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(":"))
    .map((group): number => {
      let win_cards = parseCards(group[1].split("|")[0]);
      let hand_cards = parseCards(group[1].split("|")[1]);
      let num_win = hand_cards.filter((num) => win_cards.includes(num)).length;
      return num_win == 0 ? 0 : Math.pow(2, num_win - 1);
    })
    .reduce((sum, current) => sum + current, 0);
}
