// Part 2 for day 04 of 2023

function parseCards(cards: string): number[] {
  return cards
    .split(" ")
    .filter((group) => group != "")
    .map((group): number => parseInt(group));
}

export function part2(input: string): number {
  let card_map: { [index: number]: number } = {};
  input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(":"))
    .forEach((group) => {
      let cards = group[1];
      let tmp_num = group[0].split(" ");
      let card_number = parseInt(tmp_num[tmp_num.length - 1]);

      let win_numbers = parseCards(cards.split("|")[0]);
      let hand_numbers = parseCards(cards.split("|")[1]);

      let num_win = hand_numbers.filter((num) =>
        win_numbers.includes(num),
      ).length;
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
