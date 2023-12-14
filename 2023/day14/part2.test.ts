import { part2 } from "./part2";

describe("Testing part 2 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part2(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`),
    ).toBe(64);
  });
});
