import { part2 } from "./part2";

describe("Testing part 2 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(part2(`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`)).toBe(
      145,
    );
  });
});
