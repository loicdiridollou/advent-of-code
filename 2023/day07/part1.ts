// Part 1 for day 07 of 2023

export function parseLine(data: string): string[] {
  return data
    .split(":")[1]
    .split(" ")
    .filter((group) => group != "");
}

function stringToDict(hand: string): { [index: string]: number } {
  let dic: { [index: string]: number } = {};
  for (let chr of hand) {
    if (!dic[chr]) {
      dic[chr] = 0;
    }
    dic[chr]++;
  }
  return dic;
}

function dictToHash(dict: { [index: string]: number }): string {
  let five = 0;
  let four = 0;
  let three = 0;
  let two = 0;

  for (let key in dict) {
    if (dict[key] == 5) {
      five++;
    } else if (dict[key] == 4) {
      four++;
    } else if (dict[key] == 3) {
      three++;
    } else if (dict[key] == 2) {
      two++;
    }
  }

  return `${five}_${four}_${three}_${two}`;
}

function replaceAll(str: string, mapObj: { [index: string]: string }): string {
  var re = new RegExp(Object.keys(mapObj).join("|"), "gi");

  return str.replace(re, function (matched) {
    return mapObj[matched];
  });
}

function compareFn(a: any[], b: any[]): number {
  if (a[1] < b[1]) {
    return -1;
  } else if (a[1] == b[1]) {
    let chrMap = { T: "A", J: "B", Q: "C", K: "D", A: "E" };
    let v1 = replaceAll(a[0], chrMap);
    let v2 = replaceAll(b[0], chrMap);
    return v1 > v2 ? 1 : -1;
  }
  return 1;
}

export function part1(input: string): number {
  let vals: (string | number)[][] = input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => {
      let cards = group.split(" ")[0];
      let bid = parseInt(group.split(" ")[1]);
      return [cards, dictToHash(stringToDict(cards)), bid];
    });

  vals = vals.sort(compareFn);
  let tot = 0;
  for (let i in vals) {
    tot += (parseInt(i) + 1) * (vals[i][2] as number);
  }
  return tot;
}
