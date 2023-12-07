// Part 2 for day 07 of 2023

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

function createReplacements(hand: string): string[] {
  let hands: string[] = [];
  if (hand == "") {
    return [""];
  }
  if (!hand.includes("J")) {
    return [hand];
  }
  for (let chr of createReplacements(hand.slice(1))) {
    for (let j of "23456789TQKA") {
      let rpl = hand[0] == "J" ? j : hand[0];
      hands.push(rpl + chr);
    }
  }
  return hands;
}

function compareFn(a: any[], b: any[]): number {
  if (a[1] < b[1]) {
    return -1;
  } else if (a[1] == b[1]) {
    let chrMap = { T: "A", J: ".", Q: "C", K: "D", A: "E" };
    let v1 = replaceAll(a[0], chrMap);
    let v2 = replaceAll(b[0], chrMap);
    return v1 > v2 ? 1 : -1;
  }
  // a must be equal to b
  return 1;
}

function valueCards(hand: string): string {
  let vals = createReplacements(hand);

  vals = vals.map((group) => dictToHash(stringToDict(group)));

  return vals.reduce((max, c) => (c > max ? c : max));
}

export function part2(input: string): number {
  let vals: (string | number)[][] = input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => {
      let cards = group.split(" ")[0];
      let bid = parseInt(group.split(" ")[1]);
      return [cards, valueCards(cards), bid];
    });

  vals = vals.sort(compareFn);
  let tot = 0;
  for (let i in vals) {
    tot += (parseInt(i) + 1) * (vals[i][2] as number);
  }
  return tot;
}
