import { part1 } from "./part1";

describe("Testing part 1 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part1(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`),
    ).toBe(6440);
  });
});
