fn main() {
    let input = include_str!("./input.txt");
    let schematic: Vec<Vec<char>> =
        input.lines().map(|line| line.chars().collect()).collect();
    let sum: u32 = schematic
        .iter()
        .enumerate()
        .map(|(y, _)| sum_line(&schematic, y))
        .sum();
    println!("{sum}")
}

fn sum_line(schematic: &Vec<Vec<char>>, y: usize) -> u32 {
    let line = &schematic[y];
    let mut sum = 0;
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

            if !is_symbol(schematic[py as usize][px as usize]) {
                continue;
            }

            sum += n;
            x += digits.len();
            continue 'inner;
        }

        x += digits.len();
    }
    sum
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

fn is_symbol(c: char) -> bool {
    match c {
        '*' | '$' | '/' | '%' | '@' | '=' | '+' | '-' | '#' | '&' => true,
        _ => false,
    }
}
