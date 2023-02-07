mod aoc_2022;

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 3 {
        panic!("usage is -- year day")
    }

    let year: usize = args[1]
    .trim()
    .parse()
    .expect("year must be a positive number");

    let day: usize = args[2]
    .trim()
    .parse()
    .expect("day must be a positive number");


    match year {
        2022 => match day {
            1 => {
                println!(
                    "example: {:?}\ninput: {:?}",
                    aoc_2022::day_001::do_magic("example"),
                    aoc_2022::day_001::do_magic("input")
                )
            },
            _ => print!("day {} of year {} has no solutions", day, year),
        },
        _ => print!("year {} has no solutions", year),
    }
}
