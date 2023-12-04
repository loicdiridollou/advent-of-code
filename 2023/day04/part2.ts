// Part 2 for day 04 of 2023

import { parseCards } from "./part1";

export function part2(input: string): number {
  let card_map: { [index: number]: number } = {};
  input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(":"))
    .forEach((group) => {
      let tmp_num = group[0].split(" ");
      let card_number = parseInt(tmp_num[tmp_num.length - 1]);
      let win_cards = parseCards(group[1].split("|")[0]);
      let hand_cards = parseCards(group[1].split("|")[1]);
      let num_win = hand_cards.filter((num) => win_cards.includes(num)).length;

      if (!card_map[card_number]) {
        card_map[card_number] = 0;
      }
      card_map[card_number] += 1;
      for (let i: number = 1; i <= num_win; i++) {
        if (!card_map[card_number + i]) {
          card_map[card_number + i] = 0;
        }
        card_map[card_number + i] += card_map[card_number];
      }
    });

  return Object.values(card_map).reduce((sum, current) => sum + current, 0);
}
