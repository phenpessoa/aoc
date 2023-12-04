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

            let count = selected_numbers
                .iter()
                .filter(|&x| winning_numbers.contains(x))
                .count();

            match count {
                0 => 0,
                _ => 1 << count - 1,
            }
        })
        .sum();
    println!("{sum}");
}

fn parse_numbers(s: &str) -> Vec<usize> {
    s.split(" ")
        .filter(|&num| !num.is_empty())
        .map(|num| num.trim().parse::<usize>().unwrap())
        .collect()
}
