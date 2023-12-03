import { part1 } from "./part1";

describe("Testing part 1 with a string", () => {
  test("Value should be returned for a one-line string", () => {
    expect(part1("34*......23./...*43")).toBe(77);
  });
});

describe("Testing part 1 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part1(`432......+.
./.....87..`),
    ).toBe(519);
  });
});
