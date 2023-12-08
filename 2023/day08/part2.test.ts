import { part2 } from "./part2";

describe("Testing part 1 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part2(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`),
    ).toBe(6);
  });
});
