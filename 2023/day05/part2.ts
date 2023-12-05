// Part 2 for day 05 of 2023

export function part2(input: string): number {
  let data: string[] = input.split("\n\n").filter((group) => group != "");

  let seeds = data[0]
    .split(" ")
    .filter((group) => group != "")
    .slice(1)
    .map((group) => parseInt(group));
  let seeds_to_soil = data[1]
    .split("\n")
    .filter((group) => group != "")
    .slice(1)
    .map((group) => group.split(" ").map((group) => parseInt(group)));
  let soil_to_fertilizer = data[2]
    .split("\n")
    .filter((group) => group != "")
    .slice(1)
    .map((group) => group.split(" ").map((group) => parseInt(group)));
  let fertilizer_to_water = data[3]
    .split("\n")
    .filter((group) => group != "")
    .slice(1)
    .map((group) => group.split(" ").map((group) => parseInt(group)));
  let water_to_light = data[4]
    .split("\n")
    .filter((group) => group != "")
    .slice(1)
    .map((group) => group.split(" ").map((group) => parseInt(group)));
  let light_to_temperature = data[5]
    .split("\n")
    .filter((group) => group != "")
    .slice(1)
    .map((group) => group.split(" ").map((group) => parseInt(group)));
  let temperature_to_humidity = data[6]
    .split("\n")
    .filter((group) => group != "")
    .slice(1)
    .map((group) => group.split(" ").map((group) => parseInt(group)));
  let humidity_to_location = data[7]
    .split("\n")
    .filter((group) => group != "")
    .slice(1)
    .map((group) => group.split(" ").map((group) => parseInt(group)));
  let new_seeds: number[][] = [];

  for (let i in seeds) {
    if (parseInt(i) % 2 == 0) {
      new_seeds.push([seeds[i], seeds[i] + seeds[parseInt(i) + 1] - 1]);
    }
  }
  let location = 1.797693134862315e308;

  for (let pair of new_seeds) {
    for (let i = pair[0]; i <= pair[1]; i++) {
      let seed = i;
      for (let map of seeds_to_soil) {
        if (map[1] <= seed && seed < map[1] + map[2]) {
          seed = map[0] + (seed - map[1]);
          break;
        }
      }
      for (let map of soil_to_fertilizer) {
        if (map[1] <= seed && seed < map[1] + map[2]) {
          seed = map[0] + (seed - map[1]);
          break;
        }
      }
      for (let map of fertilizer_to_water) {
        if (map[1] <= seed && seed < map[1] + map[2]) {
          seed = map[0] + (seed - map[1]);
          break;
        }
      }
      for (let map of water_to_light) {
        if (map[1] <= seed && seed < map[1] + map[2]) {
          seed = map[0] + (seed - map[1]);
          break;
        }
      }
      for (let map of light_to_temperature) {
        if (map[1] <= seed && seed < map[1] + map[2]) {
          seed = map[0] + (seed - map[1]);
          break;
        }
      }
      for (let map of temperature_to_humidity) {
        if (map[1] <= seed && seed < map[1] + map[2]) {
          seed = map[0] + (seed - map[1]);
          break;
        }
      }
      for (let map of humidity_to_location) {
        if (map[1] <= seed && seed < map[1] + map[2]) {
          seed = map[0] + (seed - map[1]);
          break;
        }
      }

      location = Math.min(location, seed);
    }
  }
  return location;
}
