import { part2 } from "./part2";

describe("Testing part 2 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part2(`Time:      7  15   30
Distance:  9  40  200`),
    ).toBe(71503);
  });
});
