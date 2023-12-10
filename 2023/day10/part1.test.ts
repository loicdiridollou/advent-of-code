import { part1 } from "./part1";

describe("Testing part 1 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part1(`.....
.S-7.
.|.|.
.L-J.
.....`),
    ).toBe(4);
  });
});
