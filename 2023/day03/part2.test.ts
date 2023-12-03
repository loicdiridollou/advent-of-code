import { part2 } from "./part2";

describe("Testing part 2 with a string", () => {
  test("Value should be returned for a one-line string", () => {
    expect(part2("3*2.../54..")).toBe(6);
  });
});

describe("Testing part 2 with an empty string", () => {
  test("Value 0 should be returned for an empty string", () => {
    expect(part2("")).toBe(0);
  });
});

describe("Testing part 2 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part2(`...39*.../
......3..4`),
    ).toBe(117);
  });
});
