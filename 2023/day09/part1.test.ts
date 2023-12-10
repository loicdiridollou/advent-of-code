import { part1 } from "./part1";

describe("Testing part 1 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part1(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`),
    ).toBe(114);
  });
});
