// Part 1 for day 02 of 2023

export function part1(input: string): number {
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
  let locations: number[] = [];

  for (let seed of seeds) {
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
    locations.push(seed);
  }
  return Math.min.apply(Math, locations);
}
