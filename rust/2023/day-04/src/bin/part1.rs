fn main() {
    let input = include_str!("./input.txt");
    let sum: usize = input
        .lines()
        .map(|line| {
            let (winning_numbers, selected_numbers) = line
                .split_once(":")
                .map(|(_, card_data)| card_data)
                .unwrap()
                .split_once("|")
                .map(|str_cards| {
                    (parse_numbers(str_cards.0), parse_numbers(str_cards.1))
                })
                .unwrap();

            let count = (winning_numbers & selected_numbers).count_ones();

            match count {
                0 => 0,
                _ => 1 << count - 1,
            }
        })
        .sum();
    println!("{sum}");
}

fn parse_numbers(s: &str) -> u128 {
    let mut out: u128 = 0;
    s.split_whitespace()
        .map(|num| out |= 1 << num.trim().parse::<usize>().unwrap())
        .for_each(|_| ());
    out
}
