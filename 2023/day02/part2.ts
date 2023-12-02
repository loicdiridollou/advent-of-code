// Part 2 for day 02 of 2023

export function part2(input: string): number {
  return input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(":"))
    .map((group): number => {
      let games: string[] = group[1].split(";");
      let new_dic: { [index: string]: number } = {};
      for (let game in games) {
        let elements = games[game].split(",");
        for (let key in elements) {
          let vals = elements[key].trim().split(" ");
          if (new_dic[vals[1]]) {
            new_dic[vals[1]] = Math.max(parseInt(vals[0]), new_dic[vals[1]]);
          } else {
            new_dic[vals[1]] = parseInt(vals[0]);
          }
        }
      }
      return Object.values(new_dic).reduce(
        (mult, current) => mult * current,
        1,
      );
    })
    .reduce((sum, current) => sum + current, 0);
}
