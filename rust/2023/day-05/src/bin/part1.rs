fn main() {
    let input = include_str!("./input.txt").replace("\r\n", "\n");
    let parts: Vec<_> = input.split("\n\n").collect();

    let seeds = parse_seeds(parts[0]);
    let seeds_to_soil = parse_maps(parts[1]);
    let soil_to_fertilizer = parse_maps(parts[2]);
    let fertilizer_to_water = parse_maps(parts[3]);
    let water_to_light = parse_maps(parts[4]);
    let light_to_temperature = parse_maps(parts[5]);
    let temperature_to_humidity = parse_maps(parts[6]);
    let humidity_to_location = parse_maps(parts[7]);

    let lowest_location = seeds
        .iter()
        .map(|&seed| {
            let soil = match_input_to_map(seed, &seeds_to_soil);
            let fertilizer = match_input_to_map(soil, &soil_to_fertilizer);
            let water = match_input_to_map(fertilizer, &fertilizer_to_water);
            let light = match_input_to_map(water, &water_to_light);
            let temperature = match_input_to_map(light, &light_to_temperature);
            let humidity =
                match_input_to_map(temperature, &temperature_to_humidity);
            match_input_to_map(humidity, &humidity_to_location)
        })
        .min()
        .unwrap();
    println!("{lowest_location}");
}

struct Map {
    dst: isize,
    src: isize,
    range: isize,
}

fn parse_seeds(s: &str) -> Vec<isize> {
    let (_, raw_seeds) = s.split_once(":").unwrap();
    raw_seeds
        .split_whitespace()
        .map(|num| num.parse::<isize>().unwrap())
        .collect()
}

fn parse_maps(s: &str) -> Vec<Map> {
    s.lines()
        .skip(1)
        .map(|line| line.split_whitespace())
        .map(|nums| {
            let mut map = Map {
                dst: 0,
                src: 0,
                range: 0,
            };
            nums.map(|num| num.parse::<isize>().unwrap())
                .enumerate()
                .for_each(|(i, parsed)| match i {
                    0 => map.dst = parsed,
                    1 => map.src = parsed,
                    2 => map.range = parsed,
                    _ => panic!("unknown num index: {i}"),
                });
            map
        })
        .collect()
}

fn match_input_to_map(input: isize, maps: &[Map]) -> isize {
    for map in maps.iter() {
        if input >= map.src && input < map.src + map.range {
            return map.dst - map.src + input;
        };
    }

    // Any source numbers that aren't mapped correspond to
    // the same destination number.
    input
}
