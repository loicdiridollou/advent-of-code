import { part2 } from "./part2";

describe("Testing part 2 with a string", () => {
  test("Value should be returned for a one-line string", () => {
    expect(part2("onefw2fgjwl4jfowd5")).toBe(15);
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
      part2(`gfonejrei3mbjgio3bjgr4
    utioerwpq2nbjfudi4fjdthreeks`),
    ).toBe(37);
  });
});
