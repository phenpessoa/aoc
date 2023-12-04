use std::collections::HashMap;

fn main() {
    let input = include_str!("./input.txt");
    let schematic: Vec<Vec<char>> =
        input.lines().map(|line| line.chars().collect()).collect();

    let mut gear_map: GearMap = HashMap::new();

    schematic
        .iter()
        .enumerate()
        .map(|(y, _)| get_gears_numbers(&schematic, y))
        .map(|gear_numbers| {
            gear_numbers.iter().for_each(|gear_number| {
                gear_map
                    .entry(gear_number.0)
                    .or_insert_with(Vec::new)
                    .push(gear_number.1)
            })
        })
        .for_each(|unit| unit);

    let sum: u32 = gear_map
        .iter()
        .filter_map(|n_vec| {
            if n_vec.1.len() != 2 {
                return None;
            };
            Some(n_vec.1.iter().product::<u32>())
        })
        .sum();

    println!("{sum}")
}

type GearMap = HashMap<(usize, usize), Vec<u32>>;

fn get_gears_numbers(
    schematic: &Vec<Vec<char>>,
    y: usize,
) -> Vec<((usize, usize), u32)> {
    let line = &schematic[y];
    let mut out: Vec<((usize, usize), u32)> = Vec::new();
    let mut x = 0;
    'inner: while x < line.len() {
        let c = schematic[y][x];
        if !c.is_ascii_digit() {
            x += 1;
            continue;
        }

        let digits: Vec<_> = line
            .iter()
            .skip(x)
            .take_while(|c| c.is_ascii_digit())
            .collect();

        let n = digits
            .iter()
            .fold(0, |acc, &d| acc * 10 + d.to_digit(10).unwrap());

        let moves = match digits.len() {
            1 => MOVES_1.iter(),
            2 => MOVES_2.iter(),
            3 => MOVES_3.iter(),
            _ => panic!("unknown number length"),
        };

        for m in moves {
            let (px, py) = (m.0 + x as isize, m.1 + y as isize);

            if px < 0
                || py < 0
                || px as usize >= line.len()
                || py as usize >= schematic.len()
            {
                continue;
            }

            let py = py as usize;
            let px = px as usize;

            if schematic[py][px] != '*' {
                continue;
            }

            out.push(((px, py), n));

            x += digits.len();
            continue 'inner;
        }

        x += digits.len();
    }
    out
}

static MOVES_1: [(isize, isize); 8] = [
    (-1, 0),
    (-1, 1),
    (-1, -1),
    (1, 0),
    (1, 1),
    (1, -1),
    (0, 1),
    (0, -1),
];

static MOVES_2: [(isize, isize); 10] = [
    (-1, 0),
    (-1, 1),
    (-1, -1),
    (0, 1),
    (0, -1),
    (1, 1),
    (1, -1),
    (2, 1),
    (2, 0),
    (2, -1),
];

static MOVES_3: [(isize, isize); 12] = [
    (-1, 0),
    (-1, 1),
    (-1, -1),
    (0, 1),
    (0, -1),
    (1, 1),
    (1, -1),
    (2, 1),
    (2, -1),
    (3, 1),
    (3, 0),
    (3, -1),
];
