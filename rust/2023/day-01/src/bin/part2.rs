fn main() {
    let input = include_str!("./input1.txt");
    let lines = input.lines();
    let sum: u32 = lines.map(|line| find_digit(line)).sum();
    println!("{}", sum);
}

static NUMBERS: [&'static str; 18] = [
    "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    "1", "2", "3", "4", "5", "6", "7", "8", "9",
];

fn find_digit(s: &str) -> u32 {
    let digits = NUMBERS.iter().enumerate().fold(
        (usize::MAX, usize::MIN, 0, 0),
        |(mut first_idx, mut last_idx, mut first_digit, mut last_digit),
         (i, num)| {
            if let Some(idx) = s.find(num) {
                if idx < first_idx {
                    first_idx = idx;
                    first_digit = NUMBERS[i % 9 + 9].parse::<u32>().unwrap();
                }
            }

            if let Some(idx) = s.rfind(num) {
                if idx >= last_idx {
                    last_idx = idx;
                    last_digit = NUMBERS[i % 9 + 9].parse::<u32>().unwrap();
                }
            }

            (first_idx, last_idx, first_digit, last_digit)
        },
    );

    digits.2 * 10 + digits.3
}
