import { part1 } from "./part1";

describe("Testing part 1 with a string", () => {
  test("Value should be returned for a one-line string", () => {
    expect(part1("fw2fgjwl4jfowd5")).toBe(25);
  });
});

describe("Testing part 1 with an empty string", () => {
  test("Value 0 should be returned for an empty string", () => {
    expect(part1("")).toBe(0);
  });
});

describe("Testing part 1 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part1(`gfbjrei3mbjgio3bjgr4
    utioerwpq2nbjfudi4fjdks2`),
    ).toBe(56);
  });
});
