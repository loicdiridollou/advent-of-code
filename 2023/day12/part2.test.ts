import { part2 } from "./part2";

describe("Testing part 2 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(
      part2(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`),
    ).toBe(525152);
  });
});
