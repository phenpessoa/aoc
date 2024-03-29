fn main() {
    let input = include_str!("./input.txt");
    let mut res: Vec<_> = vec![1; input.lines().count()];
    input
        .lines()
        .enumerate()
        .map(|(i, line)| {
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

            (i, count as usize)
        })
        .for_each(|(i, count)| {
            (0..count).for_each(|j| res[i + j + 1] += res[i])
        });
    let sum: usize = res.iter().sum();
    println!("{sum}");
}

fn parse_numbers(s: &str) -> u128 {
    let mut out: u128 = 0;
    s.split_whitespace()
        .map(|num| num.trim().parse::<usize>().unwrap())
        .for_each(|parsed| out |= 1 << parsed);
    out
}
