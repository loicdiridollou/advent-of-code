import { part1 } from "./part1";

describe("Testing part 1 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part1(`1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`),
    ).toBe(5);
  });
});
