// Part 2 for day 01 of 2023
let num_map: { [index: string]: string } = {
  one: "1",
  two: "2",
  three: "3",
  four: "4",
  five: "5",
  six: "6",
  seven: "7",
  eight: "8",
  nine: "9",
  eno: "1",
  owt: "2",
  eerht: "3",
  ruof: "4",
  evif: "5",
  xis: "6",
  neves: "7",
  thgie: "8",
  enin: "9",
  "1": "1",
  "2": "2",
  "3": "3",
  "4": "4",
  "5": "5",
  "6": "6",
  "7": "7",
  "8": "8",
  "9": "9",
};

function processWords(group: string, re: RegExp): string {
  // assemble the first match and last match in the string
  let matches = group.match(re);
  let rev_matches = group.split("").reverse().join("").match(re);
  if (matches == null || rev_matches == null) {
    return "0";
  }
  return num_map[matches[0]] + num_map[rev_matches[0]];
}

export function part2(words: string): number {
  var re = new RegExp(Object.keys(num_map).join("|"), "gi");
  return words
    .split("\n")
    .map((group) => processWords(group, re))
    .map((val) => parseInt(val))
    .reduce((sum, current) => sum + current, 0);
}
