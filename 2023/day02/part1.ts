// Part 1 for day 02 of 2023

export function part1(input: string): number {
  let max_balls: { [index: string]: number } = { red: 12, green: 13, blue: 14 };
  return input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(":"))
    .map((group): number => {
      let game_id: number = parseInt(group[0].split(" ")[1]);
      let games: string[] = group[1].split(";");
      let playable = true;
      for (let game in games) {
        let new_dic: { [index: string]: number } = {};
        let elements = games[game].split(",");
        for (let key in elements) {
          let vals = elements[key].trim().split(" ");
          if (new_dic[vals[1]] && new_dic[vals[1]] < parseInt(vals[0])) {
            new_dic[vals[1]] = parseInt(vals[0]);
          } else {
            new_dic[vals[1]] = parseInt(vals[0]);
          }
        }
        for (let color in max_balls) {
          if (new_dic[color] && max_balls[color] < new_dic[color]) {
            playable = false;
          }
        }
      }
      return playable ? game_id : 0;
    })
    .reduce((sum, current) => sum + current, 0);
}
