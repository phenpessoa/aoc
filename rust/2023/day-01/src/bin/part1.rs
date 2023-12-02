fn main() {
    let sum: u32 = include_str!("./input1.txt")
        .lines()
        .map(|line| {
            let mut chars = line.chars().filter(|c| c.is_digit(10));
            let first_digit = chars.next().unwrap().to_digit(10).unwrap();
            let last_digit = chars
                .last()
                .map_or(first_digit, |c| c.to_digit(10).unwrap());
            first_digit * 10 + last_digit
        })
        .sum();
    println!("{sum}");
}
